// package main contains the logic for the "cli" command
package main

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-version"
)

// Run "is cli ..."
func (r *CLICmd) Run(ctx *Context) error {
	output, err := cliOutput(ctx, r.Version.Name)
	if err != nil {
		return err
	}

	got, err := version.NewVersion(output)
	if err != nil {
		return errors.Join(fmt.Errorf(
			"Could not parse the version (%s) found for (%s)",
			output,
			got,
		), err)
	}

	want, err := version.NewVersion(r.Version.Val)
	if err != nil {
		return errors.Join(fmt.Errorf(
			"Could not parse the version (%s) which you provided",
			r.Version.Val,
		), err)
	}

	ctx.Success = compareCLIVersions(r.Version.Op, got, want)
	if !ctx.Success && ctx.Debug {
		fmt.Printf("Comparison failed: %s %s %s\n", output, r.Version.Op, want)
	}

	return nil
}
