# is (is) an inspector for your environment

## Why?

I want to know things about my environment when I'm installing my dot files and I'm tired of having to remember the different
incantations of software versioning.

`go version` vs `perl --version` vs `tmux -V`

## How Does it Work?

`is` returns an exit code of `0` on success and non-zero (usually `1`) on failure. You can leverage this in shell scripting:

In a script:

```bash
#!/bin/bash

if eval is os name eq darwin; then
  echo "this is a mac"
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

[Download releases](https://github.com/oalders/is/releases) or use [ubi](https://github.com/houseabsolute/ubi).

```
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
