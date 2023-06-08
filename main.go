// The main package is the command line runner
package main

import (
	"os"

	"github.com/alecthomas/kong"
)

var cli struct {
	Debug   bool             `help:"turn on debugging statements"`
	OS      OSCmd            `cmd:""`
	Command CommandCmd       `cmd:""`
	Known   KnownCmd         `cmd:""`
	There   ThereCmd         `cmd:"" help:"Check if command exists"`
	Verbose bool             `help:"Print output to screen"`
	Version kong.VersionFlag `help:"Print version to screen"`
}

// Context type tracks top level debugging flag
type Context struct {
	Debug   bool
	Success bool
	Verbose bool
}

// CommandCmd type is configuration for CLI level checks
type CommandCmd struct {
	Name struct {
		Name string `arg:"" required:""`
		Op   string `arg:"" required:"" enum:"eq,ne,gt,gte,lt,lte"`
		Val  string `arg:"" required:""`
	} `arg:"" help:"Check version of command"`
}

// OSCmd type is configuration for OS level checks
type OSCmd struct {
	Name struct {
		Op  string `arg:"" required:"" enum:"eq,ne"`
		Val string `arg:"" required:""`
	} `cmd:"" help:"Check OS name"`
	Version struct {
		Op  string `arg:"" required:"" enum:"eq,ne"`
		Val string `arg:"" required:""`
	} `cmd:"" help:"Check OS version"`
}

// KnownCmd type is configuration for printing environment info
type KnownCmd struct {
	Name struct {
		Name string `arg:"" required:"" enum:"os,command-version"`
		Val  string `arg:"" required:""`
	} `arg:"" help:"Print without testing condition. e.g. \"is known os name\""`
}

// ThereCmd is configuration for finding executables
type ThereCmd struct {
	Name string `arg:"" required:""`
}

func main() {
	ctx := kong.Parse(&cli,
		kong.Vars{
			"version": "0.0.5",
		})
	runContext := Context{Debug: cli.Debug, Verbose: cli.Verbose}
	err := ctx.Run(&runContext)
	ctx.FatalIfErrorf(err)

	if runContext.Success {
		os.Exit(0)
	}
	os.Exit(1)
}
