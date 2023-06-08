// package main includes Run funcs
package main

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/hashicorp/go-version"
)

const osReleaseFile = "/etc/os-release"

// Run logic for CLI checks
func (r *CommandCmd) Run(ctx *Context, info *meta) error {
	output, err := cliOutput(ctx, r.Name.Name)
	if err != nil {
		return err
	}

	got, err := version.NewVersion(output)
	if err != nil {
		if ctx.Debug {
			fmt.Printf("Could not parse %s %v", output, err)
		}
		return err
	}

	want, err := version.NewVersion(r.Name.Val)
	if err != nil {
		if ctx.Debug {
			fmt.Printf("Could not parse %s %v", r.Name.Val, err)
		}
		return err
	}

	info.Success = compareCLIVersions(r.Name.Op, got, want)
	if !info.Success {
		if ctx.Debug {
			fmt.Printf("Comparison failed: %s %s %s\n", output, r.Name.Op, want)
		}
	}

	if ctx.Verbose {
		fmt.Println(output)
	}
	return nil
}

// Run logic for OS checks
func (r *OSCmd) Run(ctx *Context, info *meta) error {
	got := runtime.GOOS
	want := r.Name.Val

	switch r.Name.Op {
	case "eq":
		info.Success = got == want
		if ctx.Debug {
			fmt.Printf("Comparison %s == %s %t\n", got, want, info.Success)
		}
	case "ne":
		info.Success = got != want
		if ctx.Debug {
			fmt.Printf("Comparison %s != %s %t\n", got, want, info.Success)
		}
	}

	return nil
}

// Run logic for printing
func (r *KnownCmd) Run(ctx *Context, info *meta) error {
	switch r.Name.Name {
	case "os":
		switch r.Name.Val {
		case "arch":
			fmt.Println(runtime.GOARCH)
			info.Success = true
		case "id":
			if runtime.GOOS == "linux" {
				if ctx.Debug {
					fmt.Println("Trying to parse " + osReleaseFile)
				}
				release, err := maybeReadINI(osReleaseFile)
				if err == nil && release != nil && release.ID != "" {
					fmt.Printf("%s\n", release.ID)
					info.Success = true
				}
			}
		case "id-like":
			if runtime.GOOS == "linux" {
				if ctx.Debug {
					fmt.Println("Trying to parse " + osReleaseFile)
				}
				release, err := maybeReadINI(osReleaseFile)
				if err == nil && release != nil && release.IDLike != "" {
					fmt.Printf("%s\n", release.IDLike)
					info.Success = true
				}
			}
		case "pretty-name":
			if runtime.GOOS == "linux" {
				if ctx.Debug {
					fmt.Println("Trying to parse " + osReleaseFile)
				}
				release, err := maybeReadINI(osReleaseFile)
				if err == nil && release != nil && release.PrettyName != "" {
					fmt.Printf("%s\n", release.PrettyName)
					info.Success = true
				}
			}
		case "name":
			info.Success = true
			fmt.Printf("%s\n", runtime.GOOS)
		case "version":
			if runtime.GOOS == "darwin" {
				o, err := macVersion()
				if err != nil {
					return err
				}
				fmt.Printf("%s\n", o)
				info.Success = true
			} else if runtime.GOOS == "linux" {
				if ctx.Debug {
					fmt.Println("Trying to parse " + osReleaseFile)
				}
				release, err := maybeReadINI(osReleaseFile)
				if err == nil && release != nil && release.VersionID != "" {
					fmt.Printf("%s\n", release.VersionID)
					info.Success = true
				}
			}
		case "version-codename":
			if runtime.GOOS == "darwin" {
				if ctx.Debug {
					fmt.Println("Trying to parse " + osReleaseFile)
				}
				release, err := maybeReadINI(osReleaseFile)
				if err == nil && release != nil && release.VersionCodeName != "" {
					fmt.Printf("%s\n", release.VersionCodeName)
					info.Success = true
				}
			} else if runtime.GOOS == "darwin" {
				o, err := macVersion()
				if err != nil {
					return err
				}
				name := macCodeName(o)
				if name != "" {
					fmt.Println(name)
					info.Success = true
				}
			}
		}
	case "command-version":
		output, err := cliOutput(ctx, r.Name.Val)
		if err != nil {
			re := regexp.MustCompile(`executable file not found`)
			if re.MatchString(err.Error()) {
				if ctx.Debug {
					fmt.Printf("executable file \"%s\" not found", r.Name.Val)
				}
				return nil
			}
			return err
		}
		if len(output) > 0 {
			info.Success = true
			fmt.Printf("%s\n", strings.TrimRight(string(output), "\n"))
		}
	}

	return nil
}

// Run logic for There checks
func (r *ThereCmd) Run(ctx *Context, info *meta) error {
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
	info.Success = true
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
