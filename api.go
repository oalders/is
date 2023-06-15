// package main contains the api for the CLI
package main

import "github.com/alecthomas/kong"

var api struct {
	Debug   bool             `help:"turn on debugging statements"`
	OS      OSCmd            `cmd:"" help:"Check OS attributes. e.g. \"is os name eq darwin\""`
	CLI     CLICmd           `cmd:"" help:"Check cli version. e.g. \"is cli version tmux gte 3\""`
	Known   KnownCmd         `cmd:""`
	There   ThereCmd         `cmd:"" help:"Check if command exists. e.g. \"is there git\""`
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
		Name string `arg:"" required:"" help:"[name of command or path to command]"`
		Op   string `arg:"" required:"" enum:"eq,ne,gt,gte,lt,lte" help:"[eq|ne|gt|gte|lt|lte]"`
		Val  string `arg:"" required:""`
	} `cmd:"" help:"Check version of command. e.g. \"is cli version tmux gte 3\""`
}

// OSCmd type is configuration for OS level checks
type OSCmd struct {
	Attr string `arg:"" required:"" name:"attribute" help:"[arch|id|id-like|pretty-name|name|version|version-codename]"`
	Op   string `arg:"" required:"" enum:"eq,ne,gt,gte,lt,lte" help:"[eq|ne|gt|gte|lt|lte]"`
	Val  string `arg:"" required:""`
}

// KnownCmd type is configuration for printing environment info
type KnownCmd struct {
	OS struct {
		Attr string `arg:"" required:"" name:"attribute" help:"[arch|id|id-like|pretty-name|name|version|version-codename]"`
	} `cmd:"" help:"Print without check. e.g. \"is known os name\""`
	CLI struct {
		Attr string `arg:"" name:"attribute" required:"" enum:"version"`
		Name string `arg:"" required:""`
	} `cmd:"" help:"Print without check. e.g. \"is known cli version git\""`
}

// ThereCmd is configuration for finding executables
type ThereCmd struct {
	Name string `arg:"" required:""`
}
