// Package main contains the logic for the "cli" command
package main

import (
	"errors"
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

// Run "is cli ...".
func (r *CLICmd) Run(ctx *types.Context) error {
	if r.Age.Name != "" {
		return runAge(ctx, r.Age.Name, r.Age.Op, r.Age.Val, r.Age.Unit)
	}
	if r.Version.Name != "" {
		output, err := parser.CLIOutput(ctx, r.Version.Name)
		if err != nil {
			return err
		}
		return compare.Versions(ctx, r.Version.Op, output, r.Version.Val)
	}

	// This is output
	//nolint:gosec
	output, err := command.Output(
		exec.Command(r.Output.Command, r.Output.Arg...), r.Output.Stream,
	)
	if err != nil {
		return err
	}

	switch r.Output.Compare {
	case "string":
		return compare.Strings(ctx, r.Output.Op, output, r.Output.Val)
	case "version":
		return compare.Versions(ctx, r.Output.Op, output, r.Output.Val)
	case "integer":
		return compare.Integers(ctx, r.Output.Op, output, r.Output.Val)
	case "float":
		return compare.Floats(ctx, r.Output.Op, output, r.Output.Val)
	default:
		return compare.Optimistic(ctx, r.Output.Op, output, r.Output.Val)
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

func runAge(ctx *types.Context, name, ageOperator, ageValue, ageUnit string) error {
	path, err := exec.LookPath(name)
	if err != nil {
		return errors.Join(errors.New("could not find command"), err)
	}

	info, err := os.Stat(path)
	if err != nil {
		return errors.Join(errors.New("could not stat command"), err)
	}

	dur, err := age.StringToDuration(ageValue, ageUnit)
	if err != nil {
		return err
	}
	targetTime := time.Now().Add(*dur)

	compareAge(ctx, info.ModTime(), targetTime, ageOperator, path)
	return err
}
