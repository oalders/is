# is (is) an inspector for your environment

## What is `is`?

`is` is an inspector for your environment. It tries to make it a little bit
easier to run commands which rely on a specific OS or a specific CLI version.

### Quick Examples

#### Is the minimum version of this tool available?

```bash
$ is command git gt 2.40.0 && echo 🥳 || echo 😢
```

#### Is this the target Operating System?
```bash
$ (is os name eq darwin && echo 🍏 ) || (is os name eq linux && echo 🐧)
```

#### Do we have go? Then update `go` binaries used by `vim-go`.

```bash
#!/bin/bash

if eval is there go; then
    nvim +':GoUpdateBinaries' +qa || true
fi
```

#### Extracting Versions of Available Tools

Forget about the remembering the different `version` incantations of command
line tools. Don't make up regexes to extract the actual version number.

### Go (version)

```text
$ go version
go version go1.20.4 darwin/arm64
```

```text
$ is known command-version go
1.20.4
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
$ is known command-version perl
v5.36.1
```

### tmux (-V)

```text
$ tmux -V
tmux 3.3a
```

```text
$ is known command-version tmux
3.3a
```



## How Does it Work?

`is` returns an exit code of `0` on success and non-zero (usually `1`) on failure. You can leverage this in shell scripting:

In a script:

```bash
#!/bin/bash

if eval is os name eq darwin; then
  # This is a Mac. Creating karabiner config dir."
  mkdir -p "$HOME/.config/karabiner"
fi
```

At the command line:

```bash
is os name ne darwin && echo "this is not a mac"
```

## `is os name`: Check OS Name

Available comparisons are:

* `eq`
* `ne`

### Equality

```text
is os name eq darwin
```

### Inequality

```text
is os name ne linux
```

## `is command`: Check Command version

Available comparisons are:

* `lt`
* `lte`
* `eq`
* `gte`
* `gt`
* `ne`

```text
is command go lt 1.20.5

is command go eq 1.20.4

is command go gt 1.20.3

is command go ne 1.20.2
```

## `is there`: Check if Command is Available

```text
is there tmux && echo "we have tmux"
```

## `is known`: Print Information Without Testing It

```text
$ is known command-version tmux
3.3a
```

```text
$ is known os name
darwin
```

## `--debug`: Get Hints in Debug Mode

```text
$ is os name eq darwins --debug
Comparison failed: darwin eq darwins
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
