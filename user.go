// Package main contains the logic for the "user" command
package main

import (
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
	cmd := exec.Command("sudo", "-n", "true")
	output, err := command.Output(cmd, "stderr")
	if err != nil {
		return err
	}
	if strings.Contains(output, "password") {
		return nil
	}

	ctx.Success = true
	return nil
}
