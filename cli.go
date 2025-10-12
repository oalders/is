// Package main contains the logic for the "cli" command
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/oalders/is/age"
	"github.com/oalders/is/command"
	"github.com/oalders/is/compare"
	"github.com/oalders/is/ops"
	"github.com/oalders/is/parser"
	"github.com/oalders/is/types"
)

func execCommand(ctx *types.Context, stream, cmdLine string, args []string) (string, error) {
	cmd, cmdArgs := parseCommand(cmdLine, args)

	if ctx.Debug {
		log.Printf("Running command %s with args: %v", cmd, cmdArgs)
	}

	execCmd := exec.CommandContext(ctx.Context, cmd, cmdArgs...)
	return command.Output(execCmd, stream)
}

func parseCommand(cmdLine string, args []string) (string, []string) {
	if cmdLine == "bash -c" {
		return "bash", append([]string{"-c"}, args...)
	}

	// If no explicit args provided via --arg flags, parse command line for embedded arguments
	if len(args) == 0 {
		// Split command on spaces to extract command and its arguments
		parts := strings.Fields(cmdLine)
		if len(parts) > 1 {
			return parts[0], parts[1:]
		}
	}

	return cmdLine, args
}

// Run "is cli ...".
func (r *CLICmd) Run(ctx *types.Context) error {
	if r.Age.Name != "" {
		success, err := runCliAge(ctx, r.Age.Name, r.Age.Op, r.Age.Val, r.Age.Unit)
		ctx.Success = success
		return err
	}
	if r.Version.Name != "" {
		output, parserErr := parser.CLIOutput(ctx, r.Version.Name)
		if parserErr != nil {
			return parserErr
		}
		var success bool
		var err error
		switch {
		case r.Version.Major:
			success, err = compare.VersionSegment(ctx, r.Version.Op, output, r.Version.Val, 0)
			ctx.Success = success
			return err
		case r.Version.Minor:
			success, err = compare.VersionSegment(ctx, r.Version.Op, output, r.Version.Val, 1)
			ctx.Success = success
			return err
		case r.Version.Patch:
			success, err = compare.VersionSegment(ctx, r.Version.Op, output, r.Version.Val, 2)
			ctx.Success = success
			return err
		}

		success, err = compare.Versions(ctx, r.Version.Op, output, r.Version.Val)
		ctx.Success = success
		return err
	}

	output, err := execCommand(ctx, r.Output.Stream, r.Output.Command, r.Output.Arg)
	if err != nil {
		return err
	}

	success, err := compareOutput(ctx, r.Output.Compare, r.Output.Op, output, r.Output.Val)
	ctx.Success = success
	return err
}

func compareOutput(
	ctx *types.Context,
	comparisonType, operator, output, want string,
) (bool, error) {
	switch comparisonType {
	case "string":
		return compare.Strings(ctx, operator, output, want)
	case "version":
		return compare.Versions(ctx, operator, output, want)
	case "integer":
		return compare.Integers(ctx, operator, output, want)
	case "float":
		return compare.Floats(ctx, operator, output, want)
	default:
		return compare.Optimistic(ctx, operator, output, want), nil
	}
}

func compareAge(ctx *types.Context, modTime, targetTime time.Time, operator, path string) bool {
	// Returns -1 if cli age is older than target time
	// Returns 0 if they are the same
	// Returns 1 if cli age is younger than target time

	if ctx.Debug {
		translate := map[string]string{"gt": "before", "lt": "after"}
		log.Printf(
			"Comparison:\n%s (%s last modification)\n%s\n%s\n",
			modTime.Format("2006-01-02 15:04:05"),
			path,
			translate[operator],
			targetTime.Format("2006-01-02 15:04:05"),
		)
	}

	compare := modTime.Compare(targetTime)
	if (operator == ops.Gt || operator == ops.Gte) && compare < 1 {
		return true
	} else if (operator == ops.Lt || operator == ops.Lte) && compare >= 0 {
		return true
	}
	return false
}

func runCliAge(ctx *types.Context, name, ageOperator, ageValue, ageUnit string) (bool, error) {
	path, err := exec.LookPath(name)
	if err != nil {
		return false, fmt.Errorf("could not find command: %w", err)
	}
	return runAge(ctx, path, ageOperator, ageValue, ageUnit)
}

func runAge(ctx *types.Context, path, ageOperator, ageValue, ageUnit string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, fmt.Errorf("could not stat command: %w", err)
	}

	dur, err := age.StringToDuration(ageValue, ageUnit)
	if err != nil {
		return false, err
	}

	targetTime := time.Now().Add(*dur)
	return compareAge(ctx, info.ModTime(), targetTime, ageOperator, path), nil
}
