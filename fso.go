// Package main contains the logic for the "fso" command
package main

import (
	"errors"

	"github.com/oalders/is/types"
)

func (r *FSOCmd) Run(ctx *types.Context) error {
	if r.LastModifiedTime.Name != "" {
		return runAge(ctx, r.LastModifiedTime.Name, r.LastModifiedTime.Op, r.LastModifiedTime.Val, r.LastModifiedTime.Unit)
	}
	return errors.New("unimplemented command")
}
