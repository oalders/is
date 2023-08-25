// This file handles OS info parsing
package main

import (
	"errors"
	"fmt"
	"log"

	goversion "github.com/hashicorp/go-version"
	"github.com/oalders/is/compare"
	"github.com/oalders/is/os"
	"github.com/oalders/is/types"
)

const (
	darwin        = "darwin"
	name          = "name"
	linux         = "linux"
	osReleaseFile = "/etc/os-release"
)

// Run "is os ..."
func (r *OSCmd) Run(ctx *types.Context) error {
	want := r.Val

	attr, err := os.Info(ctx, r.Attr)
	if err != nil {
		return err
	}

	switch r.Attr {
	case "version":
		if r.Op == "like" || r.Op == "unlike" {
			err := compare.Strings(ctx, r.Op, attr, r.Val)
			if err != nil {
				return errors.Join(fmt.Errorf(
					"could not compare the version (%s) using (%s)",
					attr,
					r.Val,
				), err)
			}
			return nil
		}

		got, err := goversion.NewVersion(attr)
		if err != nil {
			return errors.Join(fmt.Errorf(
				"could not parse the version (%s) found for (%s)",
				attr,
				got,
			), err)
		}

		want, err := goversion.NewVersion(r.Val)
		if err != nil {
			return errors.Join(fmt.Errorf(
				"could not parse the version (%s) which you provided",
				r.Val,
			), err)
		}

		ctx.Success = compare.CLIVersions(r.Op, got, want)
		if !ctx.Success && ctx.Debug {
			log.Printf("Comparison failed: %s %s %s\n", r.Attr, r.Op, want)
		}
	default:
		if r.Op == "like" || r.Op == "unlike" {
			err := compare.Strings(ctx, r.Op, attr, r.Val)
			if err != nil {
				return errors.Join(fmt.Errorf(
					"could not compare the version (%s) using (%s)",
					attr,
					r.Val,
				), err)
			}
			return nil
		}
		switch r.Op {
		case "eq":
			ctx.Success = attr == want
			if ctx.Debug {
				log.Printf("Comparison %s == %s %t\n", attr, want, ctx.Success)
			}
		case "ne":
			ctx.Success = attr != want
			if ctx.Debug {
				log.Printf("Comparison %s != %s %t\n", attr, want, ctx.Success)
			}
		case "like":
		case "unlike":
		default:
			ctx.Success = false
			return fmt.Errorf(
				"the \"os\" command cannot perform the \"%s\" comparison on the \"%s\" attribute",
				r.Op,
				r.Attr,
			)
		}
	}

	if ctx.Debug {
		os, err := os.Aggregated()
		if err != nil {
			return err
		}
		log.Printf("%s\n", os)
	}

	return nil
}
