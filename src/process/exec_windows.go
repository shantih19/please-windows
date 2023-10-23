//go:build windows
// +build windows

package process

import (
	"errors"
	"os/exec"
	"syscall"
)

// ExecCommand executes an external command.
// N.B. This does not start the command - the caller must handle that (or use one
//
//	of the other functions which are higher-level interfaces).
func (e *Executor) ExecCommand(_ SandboxConfig, _ bool, command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	e.mutex.Lock()
	defer e.mutex.Unlock()

	return cmd
}

// MustSandboxCommand modifies the given command to run in the sandbox.
// On non-Linux platforms this is a no-op since namespaces aren't available.
func (e *Executor) MustSandboxCommand(cmd []string) []string {
	return cmd
}

// TODO(jpoole): figure out how to do these on windows
func Kill(pid int, sig syscall.Signal) error {
	return nil
}

// ForkExec runs the cmd asynchronously
func ForkExec(cmd string, args []string) error {
	return errors.New("ForExec isn't supported on windows")
}