// Package main contains the logic for the "user" command
package main

import (
	"errors"
	"log"
	"os/exec"
	"strings"

	"github.com/oalders/is/command"
	"github.com/oalders/is/types"
)

// Run "is user ...".
func (r *UserCmd) Run(ctx *types.Context) error {
	if ctx.Debug {
		log.Printf("Running \"sudo -n true\"\n")
	}
	cmd := exec.CommandContext(ctx.Context, "sudo", "-n", "true")
	output, err := command.Output(cmd, "stderr")
	if err != nil {
		if !errors.Is(err, exec.ErrNotFound) {
			return err
		}
		return nil
	}
	if strings.Contains(output, "password") {
		return nil
	}

	ctx.Success = true
	return nil
}
