// package main contains the logic for the "there" command
package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/oalders/is/types"
)

// Run "is there ...".
func (r *ThereCmd) Run(ctx *types.Context) error {
	args := []string{"-c", "command -v " + r.Name}
	cmd := exec.Command("sh", args...)
	if ctx.Debug {
		log.Printf("Running \"sh %s\"\n", strings.Join(args, " "))
	}
	err := cmd.Run()
	if err != nil {
		if ctx.Debug {
			log.Printf("Error was: %v\n", err)
			log.Printf("Running \"which %s\"\n", r.Name)
		}
		cmd := exec.Command("which", r.Name) //nolint:gosec
		err := cmd.Run()
		if err != nil {
			if e := (&exec.ExitError{}); errors.As(err, &e) {
				return nil
			}
			return fmt.Errorf("command run error: %w", err)
		}
	}
	ctx.Success = true
	return nil
}
