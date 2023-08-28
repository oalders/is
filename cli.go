// Package main contains the logic for the "cli" command
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

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

		units := map[string]string{
			"s":       "s",
			"second":  "s",
			"seconds": "s",
			"m":       "m",
			"minute":  "m",
			"minutes": "m",
			"h":       "h",
			"hour":    "h",
			"hours":   "h",
			"d":       "d",
			"day":     "d",
			"days":    "d",
		}

		unit := units[r.Age.Unit]
		unitMultiplier := -1
		if unit == "d" {
			unitMultiplier = -24
			unit = "h"
		}

		value, err := strconv.Atoi(r.Age.Val)
		if err != nil {
			return errors.Join(fmt.Errorf(
				"the value (%s) does not appear to be an integer",
				r.Age.Val,
			), err)
		}
		durationString := fmt.Sprintf("%d%s", value*unitMultiplier, unit)
		dur, err := time.ParseDuration(durationString)
		if err != nil {
			return errors.Join(errors.New("Cannot parse duration"), err)
		}
		targetTime := time.Now().Add(dur)

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
