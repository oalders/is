//go:build !unix

package parser

import "os/exec"

// useProcessGroup is a no-op where process groups are unavailable.
func useProcessGroup(_ *exec.Cmd) {}

// killProcessGroup is a no-op where process groups are unavailable.
func killProcessGroup(_ *exec.Cmd) {}
