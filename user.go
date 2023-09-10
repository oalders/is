// Package main contains the logic for the "user" command
package main

import (
	"errors"
	"io"
	"log"
	"os/exec"
	"strings"

	"github.com/oalders/is/types"
)

// Run "is user ...".
func (r *UserCmd) Run(ctx *types.Context) error {
	if ctx.Debug {
		log.Printf("Running \"sudo -n true\"\n")
	}
	cmd := exec.Command("sudo", "-n", "true")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return errors.Join(errors.New("open STDERR pipe"), err)
	}

	if err := cmd.Start(); err != nil {
		return errors.Join(errors.New("starting command"), err)
	}

	slurp, _ := io.ReadAll(stderr)
	if ctx.Debug {
		log.Printf("STDERR: %s", string(slurp))
	}
	if strings.Contains(string(slurp), "password") {
		return nil
	}

	if err := cmd.Wait(); err != nil {
		return errors.Join(errors.New("waiting for command"), err)
	}
	ctx.Success = true
	return nil
}
