// This file handles environment variable parsing
package main

import (
	"fmt"
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
			return fmt.Errorf("environment variable %s is not set", r.Name)
		}
		success, err := compareOutput(ctx, r.Compare, r.Op, val, r.Val)
		ctx.Success = success
		return err
	}
}
