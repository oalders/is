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
	if !r.All && !r.Verbose && !r.JSON {
		err := runCommand(ctx, r.Name)
		if err == nil {
			ctx.Success = true
			return nil
		}
		if ctx.Debug {
			log.Printf("🚀 which %s\n", r.Name)
			log.Printf("💥 %v\n", err)
		}
	}

	err := runWhich(ctx, r.Name, r.All, r.JSON)
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

//nolint:cyclop
func runWhich(ctx *types.Context, name string, all, asJSON bool) error {
	args := []string{name}
	if all {
		args = append([]string{"-a"}, args...)
	}
	cmd := exec.Command("which", args...)
	output, err := cmd.Output()
	if ctx.Debug {
		log.Printf("Running: which %s", strings.Join(args, " "))
		if len(output) != 0 {
			log.Printf("😅 %s", output)
		}
		if err != nil {
			log.Printf("💥 %v\n", err)
		}
	}
	if err != nil {
		return fmt.Errorf("command run error: %w", err)
	}
	found := strings.Split(strings.TrimSpace(
		string(output),
	), "\n")

	if asJSON {
		results := make([]map[string]string, 0, len(found))
		for _, v := range found {
			version, err := runCLI(ctx, v)
			if err != nil {
				return err
			}
			results = append(results, map[string]string{"path": v, "version": version})
		}
		encoded, err := toJSON(results)
		if err != nil {
			return err
		}
		success(ctx, encoded)
		return nil
	}

	headers := []string{
		"Path",
		"Version",
	}

	rows := make([][]string, 0, len(found))

	for _, path := range found {
		version, err := runCLI(ctx, path)
		if err != nil {
			return err
		}
		rows = append(rows, []string{path, version})
	}
	success(ctx, tabular(headers, rows))
	return nil
}
