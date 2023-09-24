// package main contains the api for the CLI
package main

type AgeCmp struct {
	Name string `arg:"" required:"" help:"[name of command or path to command]"`
	Op   string `arg:"" required:"" enum:"gt,lt" help:"[gt|lt]"`
	Val  string `arg:"" required:""`
	Unit string `arg:"" required:"" enum:"s,second,seconds,m,minute,minutes,h,hour,hours,d,day,days"`
}

//nolint:lll
type OutputCmp struct {
	Stream  string   `arg:"" required:"" enum:"stdout,stderr,combined" help:"[output stream to capture: (stdout|stderr|combined)]"`
	Command string   `arg:"" required:"" help:"[name of command or path to command plus any arguments e.g. \"uname -a\"]"`
	Op      string   `arg:"" required:"" enum:"eq,ne,gt,gte,lt,lte,like,unlike" help:"[eq|ne|gt|gte|like|lt|lte|unlike]"`
	Val     string   `arg:"" required:""`
	Arg     []string `short:"a" optional:"" help:"--arg=\"-V\" --arg foo"`
	Compare string   `default:"optimistic" enum:"float,integer,string,version,optimistic" help:"[float|integer|string|version|optimistic]"`
}

//nolint:lll
type VersionCmp struct {
	Name  string `arg:"" required:"" help:"[name of command or path to command]"`
	Op    string `arg:"" required:"" enum:"eq,ne,gt,gte,lt,lte,like,unlike" help:"[eq|ne|gt|gte|like|lt|lte|unlike]"`
	Val   string `arg:"" required:""`
	Major bool   `xor:"Major,Minor,Patch" help:"Only match on the major version (e.g. major.minor.patch)"`
	Minor bool   `xor:"Major,Minor,Patch" help:"Only match on the minor version (e.g. major.minor.patch)"`
	Patch bool   `xor:"Major,Minor,Patch" help:"Only match on the patch version (e.g. major.minor.patch)"`
}

type ArchCmd struct {
	Op  string `arg:"" required:"" enum:"eq,ne,like,unlike" help:"[eq|ne|like|unlike]"`
	Val string `arg:"" required:""`
}

// CLICmd type is configuration for CLI checks.
//
//nolint:lll
type CLICmd struct {
	Version VersionCmp `cmd:"" help:"Check version of command. e.g. \"is cli version tmux gte 3\""`
	Age     AgeCmp     `cmd:"" help:"Check last modified time of cli (2h, 4d). e.g. \"is cli age tmux gt 1 d\""`
	Output  OutputCmp  `cmd:"" help:"Check output of a command. e.g. \"is cli output stdout \"uname -a\" like \"Kernel Version 22.5\""`
}

// OSCmd type is configuration for OS level checks.
//
//nolint:lll
type OSCmd struct {
	Attr  string `arg:"" required:"" name:"attribute" help:"[id|id-like|pretty-name|name|version|version-codename]"`
	Op    string `arg:"" required:"" enum:"eq,ne,gt,gte,lt,lte,like,unlike" help:"[eq|ne|gt|gte|like|lt|lte|unlike]"`
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
	CLI struct {
		Attr string `arg:"" name:"attribute" required:"" enum:"version"`
		Name string `arg:"" required:""`
	} `cmd:"" help:"Print without check. e.g. \"is known cli version git\""`
	Major bool `xor:"Major,Minor,Patch" help:"Only print the major OS or CLI version (e.g. major.minor.patch)"`
	Minor bool `xor:"Major,Minor,Patch" help:"Only print the minor OS or CLI version (e.g. major.minor.patch)"`
	Patch bool `xor:"Major,Minor,Patch" help:"Only print the patch OS or CLI version (e.g. major.minor.patch)"`
}

// ThereCmd is configuration for finding executables.
type ThereCmd struct {
	Name string `arg:"" required:""`
}
