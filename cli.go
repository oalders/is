// Package main contains the logic for the "cli" command
package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/oalders/is/age"
	"github.com/oalders/is/compare"
	"github.com/oalders/is/ops"
	"github.com/oalders/is/parser"
	"github.com/oalders/is/types"
)

// Run "is cli ...".
func (r *CLICmd) Run(ctx *types.Context) error {
	if r.Version.Name != "" {
		output, err := parser.CLIOutput(ctx, r.Version.Name)
		if err != nil {
			return err
		}
		return compare.CLIVersions(ctx, r.Version.Op, output, r.Version.Val)
	} else if r.Age.Name != "" {
		path, err := exec.LookPath(r.Age.Name)
		if err != nil {
			return errors.Join(errors.New("could not find command"), err)
		}

		info, err := os.Stat(path)
		if err != nil {
			return errors.Join(errors.New("could not stat command"), err)
		}

		dur, err := age.StringToDuration(r.Age.Val, r.Age.Unit)
		if err != nil {
			return err
		}
		targetTime := time.Now().Add(*dur)

		compareAge(ctx, info.ModTime(), targetTime, r.Age.Op, path)
		return err
	}

	return errors.New("unimplemented comparison")
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
