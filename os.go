// This file handles OS info parsing
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"

	goversion "github.com/hashicorp/go-version"
	"github.com/oalders/is/compare"
	"github.com/oalders/is/types"
)

const darwin = "darwin"
const name = "name"
const linux = "linux"
const osReleaseFile = "/etc/os-release"

// Run "is os ..."
func (r *OSCmd) Run(ctx *types.Context) error {
	want := r.Val

	attr, err := osInfo(ctx, r.Attr)
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
			fmt.Printf("Comparison failed: %s %s %s\n", r.Attr, r.Op, want)
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
				fmt.Printf("Comparison %s == %s %t\n", attr, want, ctx.Success)
			}
		case "ne":
			ctx.Success = attr != want
			if ctx.Debug {
				fmt.Printf("Comparison %s != %s %t\n", attr, want, ctx.Success)
			}
		case "like":
		case "unlike":
		default:
			ctx.Success = false
			return fmt.Errorf(
				"The \"os\" command cannot perform the \"%s\" comparison on the \"%s\" attribute",
				r.Op,
				r.Attr,
			)
		}
	}

	if ctx.Debug {
		os, err := aggregatedOS()
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", os)
	}

	return nil
}

func osInfo(ctx *types.Context, argName string) (string, error) {
	result := ""
	switch argName {
	case "id":
		if runtime.GOOS == linux {
			if ctx.Debug {
				fmt.Println("Trying to parse " + osReleaseFile)
			}
			release, err := maybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.ID != "" {
				result = release.ID
			}
		}
	case "id-like":
		if runtime.GOOS == linux {
			if ctx.Debug {
				fmt.Println("Trying to parse " + osReleaseFile)
			}
			release, err := maybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.IDLike != "" {
				result = release.IDLike
			}
		}
	case "pretty-name":
		if runtime.GOOS == linux {
			if ctx.Debug {
				fmt.Println("Trying to parse " + osReleaseFile)
			}
			release, err := maybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.PrettyName != "" {
				result = release.PrettyName
			}
		}
	case name:
		result = runtime.GOOS
	case "version":
		if runtime.GOOS == darwin {
			o, err := macVersion()
			if err != nil {
				return result, err
			}
			result = o
		} else if runtime.GOOS == linux {
			if ctx.Debug {
				fmt.Println("Trying to parse " + osReleaseFile)
			}
			release, err := maybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.Version != "" {
				result = release.Version
			}
		}
	case "version-codename":
		if runtime.GOOS == linux {
			if ctx.Debug {
				fmt.Println("Trying to parse " + osReleaseFile)
			}
			release, err := maybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.VersionCodeName != "" {
				result = release.VersionCodeName
			}
		} else if runtime.GOOS == darwin {
			o, err := macVersion()
			if err != nil {
				return result, err
			}
			name := macCodeName(o)
			if name != "" {
				result = name
			}
		}
	}
	if result != "" {
		ctx.Success = true
	}

	return result, nil
}

func aggregatedOS() (string, error) {
	release, err := maybeReadINI(osReleaseFile)
	if err != nil {
		return "", err
	}
	if release == nil {
		release = &osRelease{}
	}
	release.Name = runtime.GOOS

	if runtime.GOOS == darwin {
		v, err := macVersion()
		if err != nil {
			return "", err
		}
		release.Version = v
		release.VersionCodeName = macCodeName(release.Version)
	}
	data, err := json.MarshalIndent(release, "", "    ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}
