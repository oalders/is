// The main package is the command line runner
package main

import (
	"os"

	"github.com/alecthomas/kong"
)

var api struct {
	Debug   bool             `help:"turn on debugging statements"`
	OS      OSCmd            `cmd:""`
	CLI     CLICmd           `cmd:""`
	Known   KnownCmd         `cmd:""`
	There   ThereCmd         `cmd:"" help:"Check if command exists"`
	Version kong.VersionFlag `help:"Print version to screen"`
}

// Context type tracks top level debugging flag
type Context struct {
	Debug   bool
	Success bool
}

// CLICmd type is configuration for CLI checks
type CLICmd struct {
	Version struct {
		Name string `arg:"" required:""`
		Op   string `arg:"" required:"" enum:"eq,ne,gt,gte,lt,lte"`
		Val  string `arg:"" required:""`
	} `cmd:"" help:"Check OS version"`
}

// OSCmd type is configuration for OS level checks
type OSCmd struct {
	Attr string `arg:"" required:"" name:"attribute"`
	Op   string `arg:"" required:"" enum:"eq,ne"`
	Val  string `arg:"" required:""`
}

// KnownCmd type is configuration for printing environment info
type KnownCmd struct {
	OS struct {
		Attr string `arg:"" required:"" enum:"arch,id,id-like,pretty-name,name,version,version-codename"`
	} `cmd:"" help:"Print without testing condition. e.g. \"is known os name\""`
	CLI struct {
		Attr string `arg:"" name:"attribute" required:"" enum:"version"`
		Name string `arg:"" required:""`
	} `cmd:"" help:"Print without testing condition. e.g. \"is known cli version git\""`
}

// ThereCmd is configuration for finding executables
type ThereCmd struct {
	Name string `arg:"" required:""`
}

func main() {
	ctx := kong.Parse(&api,
		kong.Vars{
			"version": "0.0.5",
		})
	runContext := Context{Debug: api.Debug}
	err := ctx.Run(&runContext)
	ctx.FatalIfErrorf(err)

	if runContext.Success {
		os.Exit(0)
	}
	os.Exit(1)
}
