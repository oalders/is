// Package main contains the logic for the "user" command
package main

import (
	"io"
	"log"
	"os/exec"
	"strings"

	"github.com/oalders/is/types"
)

// Run "is user ...".
func (r *UserCmd) Run(ctx *types.Context) error {
	if r.Sudoer != "" {
		if ctx.Debug {
			log.Printf("Running \"sudo -n true\"\n")
		}
		cmd := exec.Command("sudo", "-n", "true")
		stderr, err := cmd.StderrPipe()
		if err != nil {
			return err
		}

		if err := cmd.Start(); err != nil {
			return err
		}

		slurp, _ := io.ReadAll(stderr)
		if ctx.Debug {
			log.Printf("STDERR: %s", string(slurp))
		}
		if strings.Contains(string(slurp), "password") {
			return nil
		}

		if err := cmd.Wait(); err != nil {
			return err
		}
		ctx.Success = true
	}
	return nil
}
