# is (is) an inspector for your environment

## Introduction

`is` is an inspector for your environment. `is` tries to make it just a little bit
easier to run commands which rely on a specific OS or a specific CLI version.

### Is the minimum version of this tool available?

```text
is cli version tmux gt 3.2 && echo 🥳 || echo 😢
```

### Is this the target Operating System?

```text
(is os name eq darwin && echo 🍏 ) || (is os name eq linux && echo 🐧)
```

### Do we have go? Then update go binaries used by vim-go

```bash
#!/bin/bash

if is there go; then
    nvim +':GoUpdateBinaries' +qa || true
fi
```

Or, as a one-liner:

```text
is there go && nvim +':GoUpdateBinaries' +qa
```

### What's the version of bash?

```text
$ is known cli version bash
5.2.15
```

### Has gofumpt been modified in the last week?

```text
$ is cli age gofumpt lt 7 d
5.2.15
```

### echo the OS name

```text
$ echo "i'm on $(is known os name) for sure"
i'm on darwin for sure
```

### Get some debugging information about the OS

```text
$ is known os name --debug
{
    "name": "darwin",
    "version": "13.4",
    "version-codename": "ventura"
}
darwin
```

```text
is known os name --debug
{
    "id": "ubuntu",
    "id-like": "debian",
    "name": "linux",
    "pretty-name": "Ubuntu 22.04.2 LTS",
    "version": "22.04",
    "version-codename": "jammy"
}
linux
```

## Exit Codes are Everything

`is` returns an exit code of `0` on success and non-zero (usually `1`) on
failure. You can leverage this in shell scripting:

In a script:

```bash
#!/bin/bash

if is os name eq darwin; then
  # This is a Mac. Creating karabiner config dir."
  mkdir -p "$HOME/.config/karabiner"
fi
```

At the command line:

```bash
is os name ne darwin && echo "this is not a mac"
```

### Debugging error Codes

In `bash` and `zsh` (and possibly other shells), `$?` contains the value of the
last command's exit code.


```text
$ is os name eq x
$ echo $?
1
$ is os name eq darwin
0
```

## Top Level Commands

### cli

#### age

Compare against the last modified date of a command's file.

```text
is cli age tmux lt 18 hours
```

Don't let `goimports` get more than a week out of date.

```text
is cli age goimports gt 7 days && go install golang.org/x/tools/cmd/goimports@latest
```

Update `shfmt` basically daily, but only if `go` is installed.

```bash
if is there go; then
    if is cli age shfmt gt 18 hours; then
        go install mvdan.cc/sh/v3/cmd/shfmt@latest
    fi
else
    echo "Go not found. Not installing shfmt or gofumpt"
fi
```

Supported comparisons are:

* `lt`
* `gt`

Supported units are:

* `s`
* `second`
* `seconds`
* `m`
* `minute`
* `minutes`
* `h`
* `hour`
* `hours`
* `d`
* `day`
* `days`

Note that `d|day|days` is shorthand for 24 hours. DST offsets are not taken
into account here.

The `--debug` flag can give you some helpful information when troubleshooting
date math.

#### version

Compare versions of available commands. Returns exit code of 0 if condition is
true and exit code of 1 if condition is false.

```text
is cli version go gte 1.20.4 || bash upgrade-go.sh
```

Supported comparisons are:

* `lt`
* `lte`
* `eq`
* `gte`
* `gt`
* `ne`

### os

Information specific to the current operating system

#### version

```text
is os version gt 22
```

Supported comparisons are:

* `lt`
* `lte`
* `eq`
* `gte`
* `gt`
* `ne`

#### name

Under the hood, this returns the value of `runtime.GOOS`, a constant which is
set at compile time. So, `is` reports on on the OS name which is the target of
build rather than running `uname` or something like that. This is "good enough"
for my purposes. If it's not good enough for yours, we probably need to add
more build targets.

Available comparisons are:

* `eq`
* `ne`

##### Equality

```text
is os name eq darwin
```

##### Inequality

```text
is os name ne linux
```

Possible values for `name`:

```text
darwin
linux
```

##### pretty-name

Linux only.

```text
is os pretty-name eq "Ubuntu 22.04.2 LTS"
```

Available comparisons are:

* `eq`
* `ne`

##### id

Linux only.

```text
is os id eq ubuntu
```

Available comparisons are:

* `eq`
* `ne`

##### id-like

```text
$ is os id-like eq debian
```

Available comparisons are:

* `eq`
* `ne`

##### version-codename

```text
is os version-codename eq jammy
```

Available comparisons are:

* `eq`
* `ne`

On Linux, the value for `version-codename` is taken from `/etc/os-release`. For
Macs, the values are mapped inside this application.

Possible values for Mac:

* ventura
* monterey
* big sur
* catalina
* mojave
* high sierra
* sierra
* el capitan
* yosemite
* mavericks
* mountain lion

### there

Returns exit code of 0 if command exists and exit code of 1 if command cannot
be found.

```text
is there tmux && echo "we have tmux"
```

### known

Prints known information about a resource to `STDOUT`. Returns `0` on success
and `1` if info cannot be found.

#### os

Details specific to the current operating system.

##### name

Under the hood, this returns the value of `runtime.GOOS`, a constant which is
set at compile time. So, `is` reports on on the OS name which is the target of
build rather than running `uname` or something like that. This is "good enough"
for my purposes. If it's not good enough for yours, we probably need to add
more build targets.

```text
$ is known os name
linux
```

Possible values for `name`:

```text
darwin
linux
```

##### pretty-name

Linux only.

```text
$ is known os pretty-name
Ubuntu 22.04.2 LTS
```

##### id

Linux only.

```text
$ is known os id
ubuntu
```

##### id-like

Linux only.

```text
$ is known os id-like
debian
```

##### version

```text
$ is known os version
22.04
```

##### version-codename

```text
$ is known os version-codename
jammy
```

#### cli version

```text
$ is known cli version tmux
3.3a
```

### --debug

Print some debugging information to `STDOUT`.

```text
$ is os name eq darwins --debug
Comparison failed: darwin eq darwins
```

### --help

Top level command help.

```text
Usage: is <command>

Flags:
  -h, --help       Show context-sensitive help.
      --debug      turn on debugging statements
      --version    Print version to screen

Commands:
  os <attribute> <op> <val>
    Check OS attributes. e.g. "is os name eq darwin"

  cli version <name> <op> <val>
    Check version of command. e.g. "is cli version tmux gte 3"

  known os <attribute>
    Print without check. e.g. "is known os name"

  known cli <attribute> <name>
    Print without check. e.g. "is known cli version git"

  there <name>
    Check if command exists. e.g. "is there git"

Run "is <command> --help" for more information on a command.
```

#### subcommand --help

```text
Usage: is os <attribute> <op> <val>

Check OS attributes. e.g. "is os name eq darwin"

Arguments:
  <attribute>    [arch|id|id-like|pretty-name|name|version|version-codename]
  <op>           [eq|ne|gt|gte|lt|lte]
  <val>

Flags:
  -h, --help       Show context-sensitive help.
      --debug      turn on debugging statements
      --version    Print version to screen
```

### --version

Print current version of `is`

```text
is --version
```

## Installation

* [Download a release](https://github.com/oalders/is/releases)
* `go install`
  * `go install github.com/oalders/is@latestgithub.com/oalders/is@latest`
  * `go install github.com/oalders/is@latestgithub.com/oalders/is@v0.0.5`
* Use [ubi](https://github.com/houseabsolute/ubi)

```bash
#!/usr/bin/env bash

set -eux

INSTALL_DIR="$HOME/local/bin"

if [ ! "$(command -v ubi)" ]; then
    curl --silent --location \
        https://raw.githubusercontent.com/houseabsolute/ubi/master/bootstrap/bootstrap-ubi.sh |
        TARGET=$INSTALL_DIR sh
fi

ubi --project oalders/is --in "$INSTALL_DIR"
```

## Bonus: Easier Version Parsing of Available Tools

Forget about the remembering the different `version` incantations of command
line tools. Don't make up regexes to extract the actual version number.

### Go (version)

```text
$ go version
go version go1.20.4 darwin/amd64
```

```text
$ is known cli version go
1.20.4
$ is known os name
darwin
```

### Perl (--version)

```text
$ perl --version

This is perl 5, version 36, subversion 1 (v5.36.1) built for darwin-2level

Copyright 1987-2023, Larry Wall

Perl may be copied only under the terms of either the Artistic License or the
GNU General Public License, which may be found in the Perl 5 source kit.

Complete documentation for Perl, including FAQ lists, should be found on
this system using "man perl" or "perldoc perl".  If you have access to the
Internet, point your browser at https://www.perl.org/, the Perl Home Page.
```

```text
$ is known cli version perl
v5.36.1
```

### tmux (-V)

```text
$ tmux -V
tmux 3.3a
```

```text
$ is known cli version tmux
3.3a
```
