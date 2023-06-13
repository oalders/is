// package main contains the logic for the "there" command
package main

import (
	"errors"
	"fmt"
	"os/exec"
)

// Run logic for There checks
func (r *ThereCmd) Run(ctx *Context) error {
	cmd := exec.Command("command", "-v", r.Name)
	if ctx.Debug {
		fmt.Printf("Running \"command -v %s\"\n", r.Name)
	}
	err := cmd.Run()
	if err != nil {
		if ctx.Debug {
			fmt.Printf("Running \"which %s\"\n", r.Name)
		}
		cmd := exec.Command("which", r.Name)
		err := cmd.Run()
		if err != nil {
			if e := (&exec.ExitError{}); errors.As(err, &e) {
				return nil
			}
			return err
		}
	}
	ctx.Success = true
	return nil
}
