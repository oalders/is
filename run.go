// package main includes Run funcs
package main

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/hashicorp/go-version"
)

// Run logic for CLI checks
func (r *CLICmd) Run(ctx *Context) error {
	output, err := cliOutput(ctx, r.Version.Name)
	if err != nil {
		return err
	}

	got, err := version.NewVersion(output)
	if err != nil {
		return errors.Join(fmt.Errorf(
			"Could not parse the version (%s) found for (%s)",
			output,
			got,
		), err)
	}

	want, err := version.NewVersion(r.Version.Val)
	if err != nil {
		return errors.Join(fmt.Errorf(
			"Could not parse the version (%s) which you provided",
			r.Version.Val,
		), err)
	}

	ctx.Success = compareCLIVersions(r.Version.Op, got, want)
	if !ctx.Success && ctx.Debug {
		fmt.Printf("Comparison failed: %s %s %s\n", output, r.Version.Op, want)
	}

	return nil
}

// Run logic for OS checks
func (r *OSCmd) Run(ctx *Context) error {
	want := r.Val

	attr, err := osInfo(ctx, r.Attr)
	if err != nil {
		return err
	}

	switch r.Val {
	case "version":
		got, err := version.NewVersion(attr)
		if err != nil {
			return errors.Join(fmt.Errorf(
				"Could not parse the version (%s) found for (%s)",
				attr,
				got,
			), err)
		}

		want, err := version.NewVersion(r.Val)
		if err != nil {
			return errors.Join(fmt.Errorf(
				"Could not parse the version (%s) which you provided",
				r.Val,
			), err)
		}

		ctx.Success = compareCLIVersions(r.Op, got, want)
		if !ctx.Success && ctx.Debug {
			fmt.Printf("Comparison failed: %s %s %s\n", r.Attr, r.Op, want)
		}
	default:
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
		}
	}

	return nil
}

// Run logic for printing
func (r *KnownCmd) Run(ctx *Context) error {
	result := ""
	var err error
	if r.OS.Attr != "" {
		result, err = osInfo(ctx, r.OS.Attr)
	} else if r.CLI.Attr != "" {
		result, err = cliOutput(ctx, r.CLI.Name)
		if err != nil {
			re := regexp.MustCompile(`executable file not found`)
			if re.MatchString(err.Error()) {
				if ctx.Debug {
					fmt.Printf("executable file \"%s\" not found", r.CLI.Name)
				}
				return nil
			}
			return err
		}
		if len(result) > 0 {
			if err != nil {
				result = strings.TrimRight(string(result), "\n")
			}
		}
	}
	if err != nil {
		return err
	}
	fmt.Println(result)
	ctx.Success = true

	return nil
}

// Run logic for There checks
func (r *ThereCmd) Run(ctx *Context) error {
	cmd := exec.Command("command", "-v", r.Name)
	if ctx.Debug {
		fmt.Printf("Running \"command -v %s\"\n", r.Name)
	}
	err := cmd.Run()
	if err != nil {
		if ctx.Debug {
			fmt.Printf("Running \"which %s\"\n", r.Name)
		}
		cmd := exec.Command("which", r.Name)
		err := cmd.Run()
		if err != nil {
			if e := (&exec.ExitError{}); errors.As(err, &e) {
				return nil
			}
			return err
		}
	}
	ctx.Success = true
	return nil
}

func macCodeName(osVersion string) string {
	got, err := version.NewVersion(osVersion)
	if err != nil {
		return ""
	}
	// https://en.wikipedia.org/wiki/List_of_Apple_codenames
	segments := got.Segments()
	name := ""
	switch segments[0] {
	case 13:
		name = "ventura"
	case 12:
		name = "monterey"
	case 11:
		name = "big sur"
	case 10:
		switch segments[1] {
		case 15:
			name = "catalina"
		case 14:
			name = "mojave"
		case 13:
			name = "high sierra"
		case 12:
			name = "sierra"
		case 11:
			name = "el capitan"
		case 10:
			name = "yosemite"
		case 9:
			name = "mavericks"
		case 8:
			name = "mountain lion" // released 2012
		}
	}
	return name
}

func macVersion() (string, error) {
	o, err := exec.Command("sw_vers", "-productVersion").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(o), "\n"), nil
}
