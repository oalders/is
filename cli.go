// Package main contains the logic for the "cli" command
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/oalders/is/age"
	"github.com/oalders/is/command"
	"github.com/oalders/is/compare"
	"github.com/oalders/is/ops"
	"github.com/oalders/is/parser"
	"github.com/oalders/is/types"
)

func execCommand(ctx *types.Context, stream, cmd string, args []string) (string, error) {
	if cmd == "bash -c" {
		cmd = "bash"
		args = append([]string{"-c"}, args...)
	}
	if ctx.Debug {
		log.Printf("Running command %s with args: %v", cmd, args)
	}
	return command.Output(
		exec.Command(cmd, args...), stream,
	)
}

// Run "is cli ...".
func (r *CLICmd) Run(ctx *types.Context) error {
	if r.Age.Name != "" {
		return runCliAge(ctx, r.Age.Name, r.Age.Op, r.Age.Val, r.Age.Unit)
	}
	if r.Version.Name != "" {
		output, err := parser.CLIOutput(ctx, r.Version.Name)
		if err != nil {
			return err
		}
		switch {
		case r.Version.Major:
			return compare.VersionSegment(ctx, r.Version.Op, output, r.Version.Val, 0)
		case r.Version.Minor:
			return compare.VersionSegment(ctx, r.Version.Op, output, r.Version.Val, 1)
		case r.Version.Patch:
			return compare.VersionSegment(ctx, r.Version.Op, output, r.Version.Val, 2)
		}

		return compare.Versions(ctx, r.Version.Op, output, r.Version.Val)
	}

	output, err := execCommand(ctx, r.Output.Stream, r.Output.Command, r.Output.Arg)
	if err != nil {
		return err
	}

	return compareOutput(ctx, r.Output.Compare, r.Output.Op, output, r.Output.Val)
}

func compareOutput(ctx *types.Context, comparisonType, operator, output, want string) error {
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
		return compare.Optimistic(ctx, operator, output, want)
	}
}

func compareAge(ctx *types.Context, modTime, targetTime time.Time, operator, path string) {
	// Returns -1 if cli age is older than target time
	// Returns 0 if they are the same
	// Returns 1 if cli age is younger than target time
	compare := modTime.Compare(targetTime)
	if (operator == ops.Gt || operator == ops.Gte) && compare < 1 {
		ctx.Success = true
	} else if (operator == ops.Lt || operator == ops.Lte) && compare >= 0 {
		ctx.Success = true
	}
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
}

func runCliAge(ctx *types.Context, name, ageOperator, ageValue, ageUnit string) error {
	path, err := exec.LookPath(name)
	if err != nil {
		return fmt.Errorf("could not find command: %w", err)
	}
	return runAge(ctx, path, ageOperator, ageValue, ageUnit)
}

func runAge(ctx *types.Context, path, ageOperator, ageValue, ageUnit string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("could not stat command: %w", err)
	}

	dur, err := age.StringToDuration(ageValue, ageUnit)
	if err != nil {
		return err
	}
	targetTime := time.Now().Add(*dur)

	compareAge(ctx, info.ModTime(), targetTime, ageOperator, path)
	return err
}
