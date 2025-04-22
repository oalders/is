// This file handles battery info checks
package main

import (
	"errors"
	"strconv"

	"github.com/oalders/is/battery"
	"github.com/oalders/is/compare"
	"github.com/oalders/is/types"
)

// Run "is battery ...".
//
//nolint:cyclop
func (r *BatteryCmd) Run(ctx *types.Context) error {
	attr, err := battery.GetAttr(ctx, r.Attr, r.Nth)
	ctx.Success = false

	if err != nil {
		return err
	}

	switch got := attr.(type) {
	case float64:
		want, err := strconv.ParseFloat(r.Val, 64)
		if err != nil {
			return errors.Join(
				errors.New("wanted result could not be converted to a float"),
				err,
			)
		}
		ok, err := compare.IntegersOrFloats(ctx, r.Op, got, want)
		if err != nil {
			return err
		}
		ctx.Success = ok
	case int:
		want, err := strconv.ParseInt(r.Val, 0, 32)
		if err != nil {
			return errors.Join(
				errors.New("wanted result could not be converted to an integer"),
				err,
			)
		}
		ok, err := compare.IntegersOrFloats(ctx, r.Op, got, int(want))
		if err != nil {
			return err
		}
		ctx.Success = ok
	case string:
		ok, err := compare.Strings(ctx, r.Op, got, r.Val)
		if err != nil {
			return err
		}
		ctx.Success = ok
	default:
		return errors.New("unexpected type for " + r.Val)
	}
	return nil
}
