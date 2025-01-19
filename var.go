// This file handles environment variable parsing
package main

import (
	"errors"
	"os"

	"github.com/oalders/is/types"
)

func (r *VarCmd) Run(ctx *types.Context) error {
	ctx.Success = false

	switch r.Op {
	case "set":
		_, exists := os.LookupEnv(r.Name)
		if exists {
			ctx.Success = true
		}
		return nil
	case "unset":
		_, exists := os.LookupEnv(r.Name)
		if !exists {
			ctx.Success = true
		}
		return nil
	default:
		val, exists := os.LookupEnv(r.Name)
		if !exists {
			return errors.New("environment variable not set")
		}
		return compareOutput(ctx, r.Compare, r.Op, val, r.Val)
	}
}
