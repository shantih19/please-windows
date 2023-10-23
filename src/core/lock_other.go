//go:build !windows
// +build !windows

package core

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

type logFunc func(format string, args ...interface{})

func acquireFileLock(file *os.File, how int, levelLog logFunc) error {
	// Try a non-blocking acquire first so we can warn the user if we're waiting.
	log.Debug("Attempting to acquire lock for %s...", file.Name())
	err := syscall.Flock(int(file.Fd()), how|syscall.LOCK_NB)
	if err != nil {
		pid, err := os.ReadFile(file.Name())
		if err == nil && len(pid) > 0 {
			levelLog("Looks like process with PID %s has already acquired the lock for %s. Waiting for it to finish...", string(pid), file.Name())
		} else {
			levelLog("Looks like another process has already acquired the lock for %s. Waiting for it to finish...", file.Name())
		}

		if err := syscall.Flock(int(file.Fd()), how); err != nil {
			return fmt.Errorf("Failed to acquire lock for %s: %w", file.Name(), err)
		}
	}
	log.Debug("Acquired lock for %s", file.Name())

	// Record content
	if err := file.Truncate(0); err == nil {
		file.WriteAt([]byte(strconv.Itoa(os.Getpid())), 0)
	}

	return nil
}

// ReleaseFileLock releases the lock and closes the file handle.
// Does not die on errors, at this point it wouldn't really do any good.
func ReleaseFileLock(file *os.File) {
	if file == nil {
		return
	}

	if err := syscall.Flock(int(file.Fd()), syscall.LOCK_UN); err != nil {
		log.Errorf("Failed to release lock for %s: %s", file.Name(), err) // No point making this fatal really
	}
	if err := file.Close(); err != nil {
		log.Errorf("Failed to close lock file %s: %s", file.Name(), err)
	}
}

// AcquireSharedRepoLock acquires a shared lock on the repo lock file. The file descriptor is reused if already opened
// allowing its lock mode to be replaced. Dies if the lock cannot be successfully acquired.
func AcquireSharedRepoLock() {
	if err := acquireRepoLock(syscall.LOCK_SH); err != nil {
		log.Fatal(err)
	}
}

// AcquireExclusiveRepoLock acquires an exclusive lock on the repo lock file. The file descriptor is reused if already opened
// allowing its lock mode to be replaced. Dies if the lock cannot be successfully acquired.
func AcquireExclusiveRepoLock() {
	if err := acquireRepoLock(syscall.LOCK_EX); err != nil {
		log.Fatal(err)
	}
}

// AcquireExclusiveFileLock opens a file to acquire an exclusive lock.
// Dies if the lock cannot be successfully acquired.
func AcquireExclusiveFileLock(filePath string) *os.File {
	lockFile, err := acquireOpenFileLock(filePath, syscall.LOCK_EX)
	if err != nil {
		log.Fatal(err)
	}
	return lockFile
}
