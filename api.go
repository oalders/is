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

//nolint:lll,govet,nolintlint
type OutputCmp struct {
	Stream  string   `arg:"" required:"" enum:"stdout,stderr,combined" help:"[output stream to capture: (stdout|stderr|combined)]"`
	Command string   `arg:"" required:"" help:"[name of command or path to command plus any arguments e.g. \"uname -a\"]"`
	Op      string   `arg:"" required:"" enum:"eq,ne,gt,gte,in,lt,lte,like,unlike" help:"[eq|ne|gt|gte|in|like|lt|lte|unlike]"`
	Val     string   `arg:"" required:""`
	Arg     []string `short:"a" optional:"" help:"--arg=\"-V\" --arg foo"`
	Compare string   `default:"optimistic" enum:"float,integer,string,version,optimistic" help:"[float|integer|string|version|optimistic]"`
}

//nolint:lll,govet,nolintlint
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
//nolint:lll,govet,nolintlint
type CLICmd struct {
	Version VersionCmp `cmd:"" help:"Check version of command. e.g. \"is cli version tmux gte 3\""`
	Age     AgeCmp     `cmd:"" help:"Check last modified time of cli (2h, 4d). e.g. \"is cli age tmux gt 1 d\""`
	Output  OutputCmp  `cmd:"" help:"Check output of a command. e.g. \"is cli output stdout \"uname -a\" like \"Kernel Version 22.5\""`
}

// FSOCmd type is configuration for FSO checks.
//
//nolint:lll,govet,nolintlint
type FSOCmd struct {
	Age AgeCmp `cmd:"" help:"Check age (last modified time) of an fso (2h, 4d). e.g. \"is fso age /tmp/log.txt gt 1 d\""`
}

// OSCmd type is configuration for OS level checks.
//
//nolint:lll,govet,nolintlint
type OSCmd struct {
	Attr  string `arg:"" required:"" name:"attribute" help:"[id|id-like|pretty-name|name|version|version-codename]"`
	Op    string `arg:"" required:"" enum:"eq,ne,gt,gte,in,lt,lte,like,unlike" help:"[eq|ne|gt|gte|in|like|lt|lte|unlike]"`
	Val   string `arg:"" required:""`
	Major bool   `xor:"Major,Minor,Patch" help:"Only match on the major OS version (e.g. major.minor.patch)"`
	Minor bool   `xor:"Major,Minor,Patch" help:"Only match on the minor OS version (e.g. major.minor.patch)"`
	Patch bool   `xor:"Major,Minor,Patch" help:"Only match on the patch OS version (e.g. major.minor.patch)"`
}

// Battery type is configuration for battery information.
//
//nolint:lll,govet,nolintlint
type Battery struct {
	Attr  string `arg:"" required:"" name:"attribute" enum:"charge-rate,count,current-capacity,current-charge,design-capacity,design-voltage,last-full-capacity,state,voltage" help:"[charge-rate|count|current-capacity|current-charge|design-capacity|design-voltage|last-full-capacity|state|voltage]"`
	Nth   int    `optional:"" default:"1" help:"Specify which battery to use (1 for the first battery)"`
	Round bool   `help:"Round float values to the nearest integer"`
}

type Summary struct {
	Attr string `arg:"" required:"" name:"attribute" enum:"battery,os" help:"[battery|os]"`
	Nth  int    `optional:"" default:"1" help:"Specify which battery to use (1 for the first battery)"`
	JSON bool   `help:"print summary as JSON"`
}

//nolint:lll,govet,nolintlint
type BatteryCmd struct {
	Battery
	Op  string `arg:"" required:"" enum:"eq,ne,gt,gte,in,lt,lte,like,unlike" help:"[eq|ne|gt|gte|in|like|lt|lte|unlike]"`
	Val string `arg:"" required:""`
}

// UserCmd type is configuration for user level checks.
//
//nolint:lll,govet,nolintlint
type UserCmd struct {
	Sudoer string `arg:"" required:"" default:"sudoer" enum:"sudoer" help:"is current user a passwordless sudoer. e.g. \"is user sudoer\""`
}

// VarCmd type is configuration for environment variable checks.
//
//nolint:lll,govet,nolintlint
type VarCmd struct {
	Name    string `arg:"" required:""`
	Op      string `arg:"" required:"" enum:"set,unset,eq,ne,gt,gte,in,lt,lte,like,unlike" help:"[set|unset|eq|ne|gt|gte|in|like|lt|lte|unlike]"`
	Val     string `arg:"" optional:""`
	Compare string `default:"optimistic" enum:"float,integer,string,version,optimistic" help:"[float|integer|string|version|optimistic]"`
}

// Allow for the following:
//
// is var EDITOR eq ""
// is var EDITOR ne ""
//
// .
func (r *VarCmd) Validate() error {
	if r.Op != "set" && r.Op != "unset" && r.Op != "eq" && r.Op != "ne" && r.Val == "" {
		return fmt.Errorf("missing required argument: val")
	}
	return nil
}

type Version struct {
	Major bool `xor:"Major,Minor,Patch" help:"Only print the major version (e.g. major.minor.patch)"`
	Minor bool `xor:"Major,Minor,Patch" help:"Only print the minor version (e.g. major.minor.patch)"`
	Patch bool `xor:"Major,Minor,Patch" help:"Only print the patch version (e.g. major.minor.patch)"`
}

type KnownCLI struct {
	Attr string `arg:"" name:"attribute" required:"" enum:"version"`
	Name string `arg:"" required:""`
	Version
}

type KnownVar struct {
	Name string `arg:"" required:""`
	JSON bool   `help:"Print output in JSON format"`
}

type KnownOS struct {
	//nolint:lll
	Attr string `arg:"" required:"" name:"attribute" help:"[id|id-like|pretty-name|name|version|version-codename]"`
	Version
}

// KnownCmd type is configuration for printing environment info.
type KnownCmd struct {
	Arch struct {
		Attr string `arg:"" required:"" default:"arch" enum:"arch"`
	} `cmd:"" help:"Print arch without check. e.g. \"is known arch\""`
	OS      KnownOS  `cmd:"" help:"Print without check. e.g. \"is known os name\""`
	CLI     KnownCLI `cmd:"" help:"Print without check. e.g. \"is known cli version git\""`
	Var     KnownVar `cmd:"" help:"Print env var without a check. e.g. \"is known var PATH\""`
	Battery Battery  `cmd:"" help:"Print battery information. e.g. \"is known battery state\""`
	Summary Summary  `cmd:"" help:"summary of available data."`
}

// ThereCmd is configuration for finding executables.
type ThereCmd struct {
	Name    string `arg:"" required:""`
	All     bool   `help:"Print all found binaries"`
	JSON    bool   `help:"Print output in JSON format"`
	Verbose bool   `help:"Show binary versions"`
}
