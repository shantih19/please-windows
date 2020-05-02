package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

type matcher interface {
	Match(name string) (bool, error)
}

type builtInGlob string

func (p builtInGlob) Match(name string) (bool, error) {
	matched, err := filepath.Match(string(p), name)
	if err != nil {
		return false, fmt.Errorf("failed to glob, invalid patern: %v, %w", string(p), err)
	}
	return matched, nil
}

type regexGlob struct {
	regex *regexp.Regexp
}

func (r regexGlob) Match(name string) (bool, error) {
	return r.regex.Match([]byte(name)), nil
}

// This converts the string pattern into a matcher. A matcher can either be one of our homebrew compiled regexs that
//support ** or a matcher that uses the built in filesystem.Match functionality.
func patternToMatcher(root, pattern string) (matcher, error) {
	fullPattern := filepath.Join(root, pattern)

	// Use the built in filesystem.Match globs when not using double star as it's far more efficient
	if !strings.Contains(pattern, "**") {
		return builtInGlob(fullPattern), nil
	}
	regex, err := regexp.Compile(toRegexString(fullPattern))
	if err != nil {
		return nil, fmt.Errorf("failed to compile glob pattern %s, %w", pattern, err)
	}
	return regexGlob{regex: regex}, nil
}

func toRegexString(pattern string) string {
	pattern = "^" + pattern + "$"
	pattern = strings.Replace(pattern, "+", "\\+", -1)         // escape +
	pattern = strings.Replace(pattern, ".", "\\.", -1)         // escape .
	pattern = strings.Replace(pattern, "?", ".", -1)           // match ? as any single char
	pattern = strings.Replace(pattern, "*", "[^/]*", -1)       // handle single (all) * components
	pattern = strings.Replace(pattern, "[^/]*[^/]*", ".*", -1) // handle ** components
	pattern = strings.Replace(pattern, "/.*/", "/(.*/)?", -1)  // Allow ** to match zero directories
	return pattern
}

// IsGlob returns true if the given pattern requires globbing (i.e. contains characters that would be expanded by it)
func IsGlob(pattern string) bool {
	return strings.ContainsAny(pattern, "*?[")
}

// Glob implements matching using Go's built-in filepath.Glob, but extends it to support
// Ant-style patterns using **.
func Glob(buildFileNames []string, rootPath string, includes, excludes []string, includeHidden bool) []string {
	if rootPath == "" {
		rootPath = "."
	}

	var filenames []string
	for _, include := range includes {
		matches, err := glob(rootPath, include, excludes, buildFileNames, includeHidden)
		if err != nil {
			panic(fmt.Errorf("error globbing files with %v: %v", include, err))
		}
		// Remove the root path from the returned files and add them to the output
		for _, filename := range matches {
			filenames = append(filenames, strings.TrimPrefix(filename, rootPath+"/"))
		}
	}
	return filenames
}

func glob(rootPath string, glob string, excludes []string, buildFileNames []string, includeHidden bool) ([]string, error) {
	globMatches, subPackages := internalGlob(rootPath, excludes, buildFileNames, filepath.Join(rootPath, glob))

	var matches []string
	for _, m := range globMatches {
		if isInDirectories(m, subPackages) || IsPackage(buildFileNames, m) {
			continue
		}
		if !includeHidden && isHidden(m) {
			continue
		}

		matches = append(matches, m)
	}
	return matches, nil
}

// shouldExcludeMatch checks if the match also matches any of the exclude patterns. If the exclude pattern is a relative
// pattern i.e. doesn't contain any /'s, then the pattern is checked against the file name part only. Otherwise the
// pattern is checked against the whole path. This is so `glob(["**/*.go"], exclude = ["*_test.go"])` will match as
// you'd expect.
func shouldExcludeMatch(root, match string, excludes []string) (bool, error) {
	for _, excl := range excludes {
		rootPath := root
		m := match

		// If the exclude pattern doesn't contain any slashes and the match does, we only match against the base of the
		// match path.
		if strings.ContainsRune(match, '/') && !strings.ContainsRune(excl, '/') {
			m = path.Base(match)
			rootPath = ""
		}

		matcher, err := patternToMatcher(rootPath, excl)
		if err != nil {
			return false, err
		}

		match, err := matcher.Match(m)
		if err != nil {
			return false, err
		}
		if match {
			return true, nil
		}
	}
	return false, nil
}

// isBuildFile checks if the filename is considered a build filename
func isBuildFile(buildFileNames []string, name string) bool {
	fileName := filepath.Base(name)
	for _, buildFileName := range buildFileNames {
		if fileName == buildFileName {
			return true
		}
	}
	return false
}

// isInDirectories checks to see if the file is in any of the provided directories
func isInDirectories(name string, directories []string) bool {
	for _, dir := range directories {
		if strings.HasPrefix(name, dir+"/") || name == dir {
			return true
		}
	}
	return false
}

// isHidden checks if the file is a hidden file i.e. starts with . or, starts and ends with #.
func isHidden(name string) bool {
	file := filepath.Base(name)
	return strings.HasPrefix(file, ".") || (strings.HasPrefix(file, "#") && strings.HasSuffix(file, "#"))
}


func internalGlob(rootPath string, excludes []string, buildFileNames []string, pattern string) ([]string, []string) {
	dir, file := filepath.Split(pattern)

	if !IsGlob(dir) {
		return globDir(rootPath, dir, file, excludes, buildFileNames, nil, nil)
	}

	matchedDirs, subPackages := internalGlob(rootPath, excludes, buildFileNames, cleanGlobPath(dir))

	var matches []string
	for _, m := range matchedDirs {
		matches, subPackages = globDir(rootPath, m, file, excludes, buildFileNames, matches, subPackages)
	}
	return matches, subPackages
}

// cleanGlobPath prepares path for glob matching.
func cleanGlobPath(path string) string {
	switch path {
	case "":
		return "."
	case string(filepath.Separator):
		// do nothing to the path
		return path
	default:
		return path[0 : len(path)-1] // chop off trailing separator
	}
}

func globDir(rootPath string, dir string, pattern string, excludes []string, buildFileNames []string, matches []string, subPackages []string) ([]string, []string) {
	f, err := os.Stat(dir)
	if err != nil {
		panic(err)
	}
	if !f.IsDir() {
		return matches, subPackages
	}

	var files []string
	if strings.HasPrefix(pattern, "**") {
		files, subPackages = ReadDirTree(dir, buildFileNames, subPackages)
		pattern = strings.TrimPrefix(pattern, "*")
	} else {
		files = ReadDir(dir)
	}

	for _, f := range files {
		if isBuildFile(buildFileNames, f) {
			packageName := filepath.Dir(f)
			if packageName == dir {
				return matches, subPackages
			}
			subPackages = append(subPackages, packageName)
		}

		matched, err := filepath.Match(pattern, filepath.Base(f))
		if err != nil {
			panic(err)
		}
		if matched {
			shouldExclude, err := shouldExcludeMatch(rootPath, f, excludes)
			if err != nil {
				panic(err)
			}
			if !shouldExclude {
				matches = append(matches, f)
			}
		}
	}

	return matches, subPackages
}

func ReadDirTree(dir string, buildFileNames []string, subPackages []string) ([]string, []string) {
	var files []string
	err := Walk(dir, func(name string, isDir bool) error {
		if dir == name {
			return nil
		}
		if isBuildFile(buildFileNames, name) {
			subPackages = append(subPackages, filepath.Dir(name))
			return filepath.SkipDir
		}
		files = append(files, name)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files, subPackages
}

func ReadDir(dir string) []string {
	var files []string

	fInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, info := range fInfos {
		files = append(files, filepath.Join(dir, info.Name()))
	}

	return files
}
