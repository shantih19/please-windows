// Representation of the build graph.
// The graph of build targets forms a DAG which we discover from the top
// down and then build bottom-up.

package core

import (
	"reflect"
	"sort"
	"sync"
)

// A BuildGraph contains all the loaded targets and packages and maintains their
// relationships, especially reverse dependencies which are calculated here.
type BuildGraph struct {
	// Map of all currently known targets by their label.
	targets sync.Map
	// Targets that have been depended on by something that we're waiting to appear.
	pendingTargets sync.Map
	// Map of all currently known packages.
	packages sync.Map
	// Registered subrepos, as a map of their name to their root.
	subrepos sync.Map
	// Reverse dependencies
	revdeps     map[*BuildTarget][]*BuildTarget
	revdepsOnce sync.Once
}

// AddTarget adds a new target to the graph.
func (graph *BuildGraph) AddTarget(target *BuildTarget) *BuildTarget {
	if _, loaded := graph.targets.LoadOrStore(target.Label, target); loaded {
		panic("Attempted to re-add existing target to build graph: " + target.Label.String())
	}
	go target.registerDependencies(graph)
	// Notify anything that called WaitForTarget
	if pkg, present := graph.pendingTargets.Load(target.Label.packageKey()); present {
		m := pkg.(*sync.Map)
		if ch, present := m.Load(target.Label.Name); present {
			close(ch.(chan struct{}))
			m.Delete(target.Label.Name)
		}
	}
	return target
}

// AddPackage adds a new package to the graph with given name.
func (graph *BuildGraph) AddPackage(pkg *Package) {
	key := packageKey{Name: pkg.Name, Subrepo: pkg.SubrepoName}
	if _, loaded := graph.packages.LoadOrStore(key, pkg); loaded {
		panic("Attempt to readd existing package: " + key.String())
	}
	// Notify anything left waiting for any targets in this package (which likely don't exist)
	if pkg, present := graph.pendingTargets.Load(key); present {
		pkg.(*sync.Map).Range(func(k, v interface{}) bool {
			close(v.(chan struct{}))
			// TODO(peterebden): Temp sanity check...
			l := BuildLabel{Subrepo: key.Subrepo, PackageName: key.Name, Name: k.(string)}
			if graph.Target(l) == nil {
				log.Warning("target %s still in map", l)
			}
			return true
		})
		graph.pendingTargets.Delete(key)
	}
}

// Target retrieves a target from the graph by label
func (graph *BuildGraph) Target(label BuildLabel) *BuildTarget {
	t, ok := graph.targets.Load(label)
	if !ok {
		return nil
	}
	return t.(*BuildTarget)
}

// TargetOrDie retrieves a target from the graph by label. Dies if the target doesn't exist.
func (graph *BuildGraph) TargetOrDie(label BuildLabel) *BuildTarget {
	target := graph.Target(label)
	if target == nil {
		log.Fatalf("Target %s not found in build graph\n", label)
	}
	return target
}

// WaitForTarget returns the given target, waiting for it to be added if it doesn't now.
// It returns nil if the target finally turns out not to exist.
func (graph *BuildGraph) WaitForTarget(label BuildLabel) *BuildTarget {
	if t := graph.Target(label); t != nil {
		return t
	} else if graph.PackageByLabel(label) != nil {
		// Check target again to avoid race conditions
		return graph.Target(label)
	}
	pkg, _ := graph.pendingTargets.LoadOrStore(label.packageKey(), &sync.Map{})
	ch, _ := pkg.(*sync.Map).LoadOrStore(label.Name, make(chan struct{}))
	<-ch.(chan struct{})
	return graph.Target(label)
}

// PackageByLabel retrieves a package from the graph using the appropriate parts of the given label.
// The Name entry is ignored.
func (graph *BuildGraph) PackageByLabel(label BuildLabel) *Package {
	return graph.Package(label.PackageName, label.Subrepo)
}

// Package retrieves a package from the graph by name & subrepo, or nil if it can't be found.
func (graph *BuildGraph) Package(name, subrepo string) *Package {
	p, present := graph.packages.Load(packageKey{Name: name, Subrepo: subrepo})
	if !present {
		return nil
	}
	return p.(*Package)
}

// PackageOrDie retrieves a package by label, and dies if it can't be found.
func (graph *BuildGraph) PackageOrDie(label BuildLabel) *Package {
	pkg := graph.PackageByLabel(label)
	if pkg == nil {
		log.Fatalf("Package %s doesn't exist in graph", label.packageKey())
	}
	return pkg
}

// AddSubrepo adds a new subrepo to the graph. It dies if one is already registered by this name.
func (graph *BuildGraph) AddSubrepo(subrepo *Subrepo) {
	if _, loaded := graph.subrepos.LoadOrStore(subrepo.Name, subrepo); loaded {
		log.Fatalf("Subrepo %s is already registered", subrepo.Name)
	}
}

// MaybeAddSubrepo adds the given subrepo to the graph, or returns the existing one if one is already registered.
func (graph *BuildGraph) MaybeAddSubrepo(subrepo *Subrepo) *Subrepo {
	if sr, present := graph.subrepos.LoadOrStore(subrepo.Name, subrepo); present {
		s := sr.(*Subrepo)
		if !reflect.DeepEqual(s, subrepo) {
			log.Fatalf("Found multiple definitions for subrepo '%s' (%+v s %+v)", s.Name, s, subrepo)
		}
		return s
	}
	return subrepo
}

// Subrepo returns the subrepo with this name. It returns nil if one isn't registered.
func (graph *BuildGraph) Subrepo(name string) *Subrepo {
	subrepo, present := graph.subrepos.Load(name)
	if !present {
		return nil
	}
	return subrepo.(*Subrepo)
}

// SubrepoOrDie returns the subrepo with this name, dying if it doesn't exist.
func (graph *BuildGraph) SubrepoOrDie(name string) *Subrepo {
	subrepo := graph.Subrepo(name)
	if subrepo == nil {
		log.Fatalf("No registered subrepo by the name %s", name)
	}
	return subrepo
}

// AllTargets returns a consistently ordered slice of all the targets in the graph.
func (graph *BuildGraph) AllTargets() BuildTargets {
	targets := BuildTargets{}
	graph.targets.Range(func(k, v interface{}) bool {
		targets = append(targets, v.(*BuildTarget))
		return true
	})
	sort.Sort(targets)
	return targets
}

// PackageMap returns a copy of the graph's internal map of name to package.
func (graph *BuildGraph) PackageMap() map[string]*Package {
	packages := map[string]*Package{}
	graph.packages.Range(func(k, v interface{}) bool {
		packages[k.(packageKey).String()] = v.(*Package)
		return true
	})
	return packages
}

// NewGraph constructs and returns a new BuildGraph.
func NewGraph() *BuildGraph {
	return &BuildGraph{}
}

// ReverseDependencies returns the set of revdeps on the given target.
// N.B. This runs in amortised constant time and initialises itself once, so should
//      only be used for queries.
func (graph *BuildGraph) ReverseDependencies(target *BuildTarget) []*BuildTarget {
	graph.revdepsOnce.Do(func() {
		graph.revdeps = map[*BuildTarget][]*BuildTarget{}
		for _, t := range graph.AllTargets() {
			for _, d := range t.Dependencies() {
				graph.revdeps[d] = append(graph.revdeps[d], t)
			}
		}
	})
	return graph.revdeps[target]
}

// DependentTargets returns the labels that 'from' should actually depend on when it declared a dependency on 'to'.
// This is normally just 'to' but could be otherwise given require/provide shenanigans.
func (graph *BuildGraph) DependentTargets(from, to BuildLabel) []BuildLabel {
	fromTarget := graph.Target(from)
	if toTarget := graph.Target(to); fromTarget != nil && toTarget != nil {
		return toTarget.ProvideFor(fromTarget)
	}
	return []BuildLabel{to}
}
