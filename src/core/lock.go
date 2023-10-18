// The logic below relies heavily on flock (advisory locks).

package core

import (
	"fmt"
	"github.com/thought-machine/please/src/fs"
	"os"
)

const repoLockFilePath = "plz-out/.lock"

const cachedirTagFile = "plz-out/CACHEDIR.TAG"

const cachedirTagFileContents = `Signature: 8a477f597d28d172789f06886806bc55
# This file is a cache directory tag created by Please.
# For information about cache directory tags see https://bford.info/cachedir/
`

var repoLockFile *os.File

// ReleaseRepoLock releases any lock mode on the repo lock file.
func ReleaseRepoLock() {
	ReleaseFileLock(repoLockFile)
	repoLockFile = nil
}

// Base function that allows to set up different repo lock modes and facilitate testing.
func acquireRepoLock(how int) error {
	if err := openRepoLockFile(); err != nil {
		return err
	} else if err := acquireFileLock(repoLockFile, how, log.Warning); err != nil {
		return err
	}
	// Write a cachedir file that indicates to some tools that this is non-essential and not to back up.
	if err := os.WriteFile(cachedirTagFile, []byte(cachedirTagFileContents), 0644); err != nil {
		log.Warningf("Failed to write cachedir tag file: %s", err)
	}
	return nil
}

// This acts like a singleton allowing the same file descriptor to used to override a previously set lock
// (from shared to exclusive and vice-versa) within the same process.
func openRepoLockFile() error {
	// Already opened
	if repoLockFile != nil {
		return nil
	}

	var err error
	if repoLockFile, err = openLockFile(repoLockFilePath); err != nil {
		return err
	}
	return nil
}

// Base function that allows to set up different lock modes and facilitate testing.
func acquireOpenFileLock(filePath string, how int) (*os.File, error) {
	lockFile, err := openLockFile(filePath)
	if err != nil {
		return nil, err
	}

	if err = acquireFileLock(lockFile, how, log.Debug); err != nil {
		return nil, err
	}

	return lockFile, nil
}

// Try to open a file for locking
func openLockFile(filePath string) (*os.File, error) {
	lockFile, err := fs.OpenDirFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("Failed to open %s to acquire lock: %w", filePath, err)
	}
	return lockFile, nil
}
