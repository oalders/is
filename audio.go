// This file handles battery info checks
package main

import (
	"fmt"

	"github.com/oalders/is/audio"
	"github.com/oalders/is/compare"
	"github.com/oalders/is/types"
)

// Run "is audio...".
func (r *AudioCmd) Run(ctx *types.Context) error {
	summary, err := audio.Summary(ctx)
	ctx.Success = false

	if err != nil {
		return err
	}

	if r.Attr == "muted" {
		ctx.Success = summary.Muted
		return nil
	}

	ok, err := compare.Integers(ctx, r.Op, fmt.Sprintf("%d", summary.Level), r.Val)
	if err != nil {
		return err
	}
	ctx.Success = ok
	return nil
}
