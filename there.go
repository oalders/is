// package main contains the logic for the "there" command
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"

	"github.com/oalders/is/attr"
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
	if ctx.Debug {
		log.Printf("🚀 sh -c %q -- %q\n", "command -v \"$1\"", name)
	}
	cmd := exec.CommandContext(ctx.Context, "sh", "-c", "command -v \"$1\"", "--", name)
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
	cmd := exec.CommandContext(ctx.Context, "which", args...)
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

	versions, err := cliVersions(ctx, found)
	if err != nil {
		return err
	}

	if asJSON {
		results := make([]map[string]string, 0, len(found))
		for i, path := range found {
			results = append(results, map[string]string{"path": path, attr.Version: versions[i]})
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

	for i, path := range found {
		rows = append(rows, []string{path, versions[i]})
	}
	success(ctx, tabular(headers, rows, false))
	return nil
}

// cliVersions looks up the version of each path concurrently. A path whose
// lookup times out gets an empty version rather than failing every path.
func cliVersions(ctx *types.Context, paths []string) ([]string, error) {
	versions := make([]string, len(paths))
	errs := make([]error, len(paths))
	var waitGroup sync.WaitGroup
	for idx, path := range paths {
		waitGroup.Go(func() {
			pathCtx := *ctx // runCLI may set Success, so don't share it
			version, err := runCLI(&pathCtx, path)
			if errors.Is(err, context.DeadlineExceeded) {
				if ctx.Debug {
					log.Printf("⏰ %s: %v", path, err)
				}
				err = nil
			}
			versions[idx], errs[idx] = version, err
		})
	}
	waitGroup.Wait()
	return versions, errors.Join(errs...)
}
