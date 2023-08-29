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

		// Returns -1 if cli age is older than target time
		// Returns 0 if they are the same
		// Returns 1 if cli age is younger than target time
		compare := info.ModTime().Compare(targetTime)
		if (r.Age.Op == "gt" || r.Age.Op == "gte") && compare < 1 {
			ctx.Success = true
		} else if (r.Age.Op == "lt" || r.Age.Op == "lte") && compare >= 0 {
			ctx.Success = true
		}

		if ctx.Debug {
			translate := map[string]string{"gt": "before", "lt": "after"}
			log.Printf(
				"Comparison:\n%s (%s last modification)\n%s\n%s\n",
				info.ModTime().Format("2006-01-02 15:04:05"),
				path,
				translate[r.Age.Op],
				targetTime.Format("2006-01-02 15:04:05"),
			)
		}
	}

	return nil
}
