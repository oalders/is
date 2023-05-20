// package main includes Run funcs
package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/hashicorp/go-version"
)

// Run logic for CLI checks
func (r *CommandCmd) Run(ctx *Context, info *meta) error {
	output, err := cliOutput(r.Name.Name)
	if err != nil {
		return err
	}

	got, err := version.NewVersion(output)
	if err != nil {
		// fmt.Printf("Could not parse %s %v", output, err)
		return err
	}

	want, err := version.NewVersion(r.Name.Val)
	if err != nil {
		// fmt.Printf("Could not parse %s %v", wantArg, err)
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
		case "name":
			info.Success = true
			fmt.Printf("%s\n", runtime.GOOS)
		case "version":
			if runtime.GOOS == "darwin" {
				o, err := exec.Command("sw_vers", "-productVersion").Output()
				if err != nil {
					return err
				}
				fmt.Printf("%s\n", strings.TrimRight(string(o), "\n"))
				info.Success = true
			}
		}
	case "command-version":
		output, err := cliOutput(r.Name.Val)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", string(output))
		info.Success = true
	}

	return nil
}

// Run logic for There checks
func (r *ThereCmd) Run(ctx *Context, info *meta) error {
	cmd := exec.Command("command", "-v", r.Name)
	err := cmd.Run()
	if err != nil {
		cmd := exec.Command("which", r.Name)
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	info.Success = true
	return nil
}
