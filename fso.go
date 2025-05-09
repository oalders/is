// Package main contains the logic for the "fso" command
package main

import (
	"errors"

	"github.com/oalders/is/types"
)

func (r *FSOCmd) Run(ctx *types.Context) error {
	if r.Age.Name != "" {
		success, err := runAge(ctx, r.Age.Name, r.Age.Op, r.Age.Val, r.Age.Unit)
		ctx.Success = success
		return err
	}
	return errors.New("unimplemented command")
}
