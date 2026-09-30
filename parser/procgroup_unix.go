//go:build unix

package parser

import (
	"errors"
	"os"
	"os/exec"
	"syscall"
)

// useProcessGroup runs cmd in its own process group so that its children
// (e.g. a PyInstaller bootloader's worker) can be signalled along with it.
// On cancel the group gets SIGTERM, giving it until cmd.WaitDelay to clean up
// (PyInstaller deletes its unpacked _MEI directory) before exec kills cmd.
func useProcessGroup(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Cancel = func() error {
		return signalGroup(cmd, syscall.SIGTERM)
	}
}

// killProcessGroup SIGKILLs whatever is left of cmd's process group.
func killProcessGroup(cmd *exec.Cmd) {
	if cmd.Process != nil { // nil if the context was done before Start
		_ = signalGroup(cmd, syscall.SIGKILL)
	}
}

func signalGroup(cmd *exec.Cmd, sig syscall.Signal) error {
	err := syscall.Kill(-cmd.Process.Pid, sig)
	if errors.Is(err, syscall.ESRCH) {
		return os.ErrProcessDone
	}
	return err //nolint:wrapcheck
}
