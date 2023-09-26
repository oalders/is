// This file handles OS info parsing
package main

import (
	"errors"

	"github.com/oalders/is/compare"
	"github.com/oalders/is/ops"
	"github.com/oalders/is/os"
	"github.com/oalders/is/types"
)

// Run "is os ...".
func (r *OSCmd) Run(ctx *types.Context) error { //nolint:cyclop
	attr, err := os.Info(ctx, r.Attr)
	ctx.Success = false // os.Info set success to true

	if err != nil {
		return err
	}

	switch r.Attr {
	case "version":
		switch {
		case r.Major:
			return compare.VersionSegment(ctx, r.Op, attr, r.Val, 0)
		case r.Minor:
			return compare.VersionSegment(ctx, r.Op, attr, r.Val, 1)
		case r.Patch:
			return compare.VersionSegment(ctx, r.Op, attr, r.Val, 2)
		}

		if r.Op == ops.Like || r.Op == ops.Unlike {
			return compare.Strings(ctx, r.Op, attr, r.Val)
		}

		err = compare.Versions(ctx, r.Op, attr, r.Val)
		if err != nil {
			ctx.Success = false
			return err
		}
	default:
		switch {
		case r.Major:
			return errors.New("--major can only be used with version")
		case r.Minor:
			return errors.New("--minor can only be used with version")
		case r.Patch:
			return errors.New("--patch can only be used with version")
		}

		switch r.Op {
		case ops.Eq, ops.In, ops.Ne, ops.Like, ops.Unlike:
			err = compare.Strings(ctx, r.Op, attr, r.Val)
		}
	}

	return err
}
