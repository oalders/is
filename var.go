// This file handles environment variable parsing
package main

import (
	"fmt"
	"os"

	"github.com/oalders/is/types"
)

func (r *VarCmd) Run(ctx *types.Context) error {
	ctx.Success = false

	val, set := os.LookupEnv(r.Name)
	switch r.Op {
	case "set":
		ctx.Success = set
		return nil
	case "unset":
		ctx.Success = !set
		return nil
	default:
		if !set {
			return fmt.Errorf("environment variable %s is not set", r.Name)
		}
		success, err := compareOutput(ctx, r.Compare, r.Op, val, r.Val)
		ctx.Success = success
		return err
	}
}
