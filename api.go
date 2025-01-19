// package main contains the api for the CLI
package main

import (
	"fmt"
)

type AgeCmp struct {
	Name string `arg:"" required:"" help:"[name of command or path to command]"`
	Op   string `arg:"" required:"" enum:"gt,lt" help:"[gt|lt]"`
	Val  string `arg:"" required:""`
	Unit string `arg:"" required:"" enum:"s,second,seconds,m,minute,minutes,h,hour,hours,d,day,days"`
}

//nolint:lll,govet
type OutputCmp struct {
	Stream  string   `arg:"" required:"" enum:"stdout,stderr,combined" help:"[output stream to capture: (stdout|stderr|combined)]"`
	Command string   `arg:"" required:"" help:"[name of command or path to command plus any arguments e.g. \"uname -a\"]"`
	Op      string   `arg:"" required:"" enum:"eq,ne,gt,gte,in,lt,lte,like,unlike" help:"[eq|ne|gt|gte|in|like|lt|lte|unlike]"`
	Val     string   `arg:"" required:""`
	Arg     []string `short:"a" optional:"" help:"--arg=\"-V\" --arg foo"`
	Compare string   `default:"optimistic" enum:"float,integer,string,version,optimistic" help:"[float|integer|string|version|optimistic]"`
}

//nolint:lll
type VersionCmp struct {
	Name  string `arg:"" required:"" help:"[name of command or path to command]"`
	Op    string `arg:"" required:"" enum:"eq,ne,gt,gte,in,lt,lte,like,unlike" help:"[eq|ne|gt|gte|in|like|lt|lte|unlike]"`
	Val   string `arg:"" required:""`
	Major bool   `xor:"Major,Minor,Patch" help:"Only match on the major version (e.g. major.minor.patch)"`
	Minor bool   `xor:"Major,Minor,Patch" help:"Only match on the minor version (e.g. major.minor.patch)"`
	Patch bool   `xor:"Major,Minor,Patch" help:"Only match on the patch version (e.g. major.minor.patch)"`
}

type ArchCmd struct {
	Op  string `arg:"" required:"" enum:"eq,ne,in,like,unlike" help:"[eq|ne|in|like|unlike]"`
	Val string `arg:"" required:""`
}

// CLICmd type is configuration for CLI checks.
//
//nolint:lll,govet
type CLICmd struct {
	Version VersionCmp `cmd:"" help:"Check version of command. e.g. \"is cli version tmux gte 3\""`
	Age     AgeCmp     `cmd:"" help:"Check last modified time of cli (2h, 4d). e.g. \"is cli age tmux gt 1 d\""`
	Output  OutputCmp  `cmd:"" help:"Check output of a command. e.g. \"is cli output stdout \"uname -a\" like \"Kernel Version 22.5\""`
}

// FSOCmd type is configuration for FSO checks.
//
//nolint:lll
type FSOCmd struct {
	Age AgeCmp `cmd:"" help:"Check age (last modified time) of an fso (2h, 4d). e.g. \"is fso age /tmp/log.txt gt 1 d\""`
}

// OSCmd type is configuration for OS level checks.
//
//nolint:lll
type OSCmd struct {
	Attr  string `arg:"" required:"" name:"attribute" help:"[id|id-like|pretty-name|name|version|version-codename]"`
	Op    string `arg:"" required:"" enum:"eq,ne,gt,gte,in,lt,lte,like,unlike" help:"[eq|ne|gt|gte|in|like|lt|lte|unlike]"`
	Val   string `arg:"" required:""`
	Major bool   `xor:"Major,Minor,Patch" help:"Only match on the major OS version (e.g. major.minor.patch)"`
	Minor bool   `xor:"Major,Minor,Patch" help:"Only match on the minor OS version (e.g. major.minor.patch)"`
	Patch bool   `xor:"Major,Minor,Patch" help:"Only match on the patch OS version (e.g. major.minor.patch)"`
}

// UserCmd type is configuration for user level checks.
//
//nolint:lll
type UserCmd struct {
	Sudoer string `arg:"" required:"" default:"sudoer" enum:"sudoer" help:"is current user a passwordless sudoer. e.g. \"is user sudoer\""`
}

// VarCmd type is configuration for environment variable checks.
//
//nolint:lll
type VarCmd struct {
	Name    string `arg:"" required:""`
	Op      string `arg:"" required:"" enum:"set,unset,eq,ne,gt,gte,in,lt,lte,like,unlike" help:"[set|unset|eq|ne|gt|gte|in|like|lt|lte|unlike]"`
	Val     string `arg:"" optional:""`
	Compare string `default:"optimistic" enum:"float,integer,string,version,optimistic" help:"[float|integer|string|version|optimistic]"`
}

func (r *VarCmd) Validate() error {
	if r.Op != "set" && r.Op != "unset" && r.Val == "" {
		return fmt.Errorf("missing required argument: val")
	}
	return nil
}

type KnownCLI struct {
	Attr string `arg:"" name:"attribute" required:"" enum:"version"`
	Name string `arg:"" required:""`
}

// KnownCmd type is configuration for printing environment info.
//
//nolint:lll
type KnownCmd struct {
	Arch struct {
		Attr string `arg:"" required:"" default:"arch" enum:"arch"`
	} `cmd:"" help:"Print arch without check. e.g. \"is known arch\""`
	OS struct {
		Attr string `arg:"" required:"" name:"attribute" help:"[id|id-like|pretty-name|name|version|version-codename]"`
	} `cmd:"" help:"Print without check. e.g. \"is known os name\""`
	CLI   KnownCLI `cmd:"" help:"Print without check. e.g. \"is known cli version git\""`
	Major bool     `xor:"Major,Minor,Patch" help:"Only print the major OS or CLI version (e.g. major.minor.patch)"`
	Minor bool     `xor:"Major,Minor,Patch" help:"Only print the minor OS or CLI version (e.g. major.minor.patch)"`
	Patch bool     `xor:"Major,Minor,Patch" help:"Only print the patch OS or CLI version (e.g. major.minor.patch)"`
}

// ThereCmd is configuration for finding executables.
type ThereCmd struct {
	Name string `arg:"" required:""`
}
