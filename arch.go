package main

import (
	"runtime"

	"github.com/oalders/is/compare"
	"github.com/oalders/is/types"
)

// Run "is arch ...".
func (r *ArchCmd) Run(ctx *types.Context) error {
	return compare.Strings(ctx, r.Op, runtime.GOARCH, r.Val)
}
