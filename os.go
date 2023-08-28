// This file handles OS info parsing
package main

import (
	"log"

	"github.com/oalders/is/compare"
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
		if r.Op == like || r.Op == unlike {
			return compare.Strings(ctx, r.Op, attr, r.Val)
		}

		err = compare.CLIVersions(ctx, r.Op, attr, r.Val)
		if err != nil {
			return err
		}
	default:
		switch r.Op {
		case "eq", "ne", like, unlike:
			err = compare.Strings(ctx, r.Op, attr, r.Val)
		}
	}

	if ctx.Debug {
		if !ctx.Success {
			log.Printf("Comparison failed: %s %s %s\n", r.Attr, r.Op, r.Val)
		}
		os, err := os.Aggregated(ctx)
		if err != nil {
			return err
		}
		log.Printf("%s\n", os)
	}

	return err
}
