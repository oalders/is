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
	err := runCommand(ctx, r.Name)
	if err == nil {
		ctx.Success = true
		return nil
	}
	if ctx.Debug {
		log.Printf("🚀 which %s\n", r.Name)
		log.Printf("💥 %v\n", err)
	}

	err = runWhich(ctx, r.Name)
	if err != nil {
		if e := (&exec.ExitError{}); errors.As(err, &e) {
			return nil
		}
		return err
	}
	ctx.Success = true
	return nil
}

func runCommand(ctx *types.Context, name string) error {
	args := []string{"-c", "command -v " + name}
	if ctx.Debug {
		log.Printf("🚀 sh -c %q\n", strings.Join(args[1:], " "))
	}
	cmd := exec.Command("sh", args...)
	output, err := cmd.Output()
	if ctx.Debug && len(output) != 0 {
		log.Printf("😅 %s", output)
	}
	return err //nolint:wrapcheck
}

func runWhich(ctx *types.Context, name string) error {
	cmd := exec.Command("which", name)
	output, err := cmd.Output()
	if ctx.Debug {
		if len(output) != 0 {
			log.Printf("😅 %s", output)
		}
		if err != nil {
			log.Printf("💥 %v\n", err)
		}
	}
	return fmt.Errorf("command run error: %w", err)
}
