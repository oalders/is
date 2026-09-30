//go:build unix

package parser

import (
	"os/exec"
	"syscall"
)

// killProcessGroupOnCancel runs cmd in its own process group and, when its
// context is done, kills the whole group rather than just cmd's process, so
// that children (e.g. a PyInstaller bootloader's worker) don't outlive it.
func killProcessGroupOnCancel(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Cancel = func() error {
		return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}
}
