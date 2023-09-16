// This file handles OS info parsing
package main

import (
	"github.com/oalders/is/compare"
	"github.com/oalders/is/ops"
	"github.com/oalders/is/os"
	"github.com/oalders/is/types"
)

// Run "is os ...".
func (r *OSCmd) Run(ctx *types.Context) error {
	attr, err := os.Info(ctx, r.Attr)
	if err != nil {
		return err
	}

	switch r.Attr {
	case "version":
		if r.Op == ops.Like || r.Op == ops.Unlike {
			return compare.Strings(ctx, r.Op, attr, r.Val)
		}

		err = compare.Versions(ctx, r.Op, attr, r.Val)
		if err != nil {
			ctx.Success = false
			return err
		}
	default:
		switch r.Op {
		case ops.Eq, ops.Ne, ops.Like, ops.Unlike:
			err = compare.Strings(ctx, r.Op, attr, r.Val)
		}
	}

	return err
}
