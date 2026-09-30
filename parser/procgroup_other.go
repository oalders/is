//go:build !unix

package parser

import "os/exec"

// killProcessGroupOnCancel is a no-op where process groups are unavailable.
func killProcessGroupOnCancel(_ *exec.Cmd) {}
