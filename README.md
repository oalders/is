# is: an inspector for your environment

<p align="center">
  <img src="logo.png" />
</p>

## Introduction

`is` is an inspector for your environment. I made it for my
[dot-files](https://github.com/oalders/dot-files). `is` tries to make it just a
little bit easier to run commands which rely on a specific OS or a specific CLI
version. Aside from the docs below, this [quick
introduction](https://www.olafalders.com/2023/09/28/is-an-inspector-for-your-environment/)
can get you up and running in a hurry.

### Is the minimum version of this tool available?

```bash
is cli version tmux gt 3.2 && echo 🥳 || echo 😢
```

### Is this the target Operating System?

```bash
(is os name eq darwin && echo 🍏 ) ||
  (is os name eq linux && echo 🐧) ||
  echo 💣
```

### Check the OS with a regex

```bash
is os name like "da\w{4}" && echo 🍏
```

### Is this a recent macOS?

```bash
is os version-codename in ventura,monterey
```

### Who am I?

```bash
is cli output stdout whoami eq olaf
```

### Do we have go? Then install `goimports`

```bash
is there go && go install golang.org/x/tools/cmd/goimports@latest
```

### What's the version of bash?

```text
$ is known cli version bash
5.2.15
```

### What's the major version of zsh?

```text
$ is known cli version --major zsh
5
```

### Has gofumpt been modified in the last week?

```text
$ is cli age gofumpt lt 7 d
```

### echo the OS name

```text
$ echo "i'm on $(is known os name) for sure"
i'm on darwin for sure
```

### Get some debugging information about the OS

macOS:

```text
$ is known os name --debug
{
    "name": "darwin",
    "version": "13.4",
    "version-codename": "ventura"
}
darwin
```

Linux:

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

### Can user sudo without a password?

```bash
is user sudoer || echo 😭
```

## Exit Codes are Everything

`is` returns an exit code of `0` on success and non-zero (usually `1`) on
failure. We can leverage this in shell scripting:

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

### Using a Regex

The `like` and `unlike` operators accept a regular expression. We may need to
quote our regex. For instance:

```bash
is os name like darw\w
```

should be

```bash
is os name like "darw\w"
```

We can the debug flag to see how our regex may have been changed by our shell:

```text
$ is os name like darw\w --debug
comparing regex "darww" with darwin
```

In this case we can see that the unquoted `\w` is turned into `w` by the shell
because it was not quoted.

We can also use regexes with no special characters at all:

```bash
is os version-codename unlike ventura
```

#### Under the Hood

🚨 Leaky abstraction alert!

Regex patterns are passed directly to Golang's `regexp.MatchString`. We can take advantage of this when crafting regexes. For instance, for a case insensitive search:

```text
is cli output stdout date like "(?i)wed"
```

## Top Level Commands

### arch

Checks against the arch which this binary has been compiled for.

```text
is arch eq amd64
```

```text
is arch like 64
```

Supported comparisons are:

* `eq`
* `ne`
* `in`
* `like`
* `unlike`

We can try `is known arch` to get the value for our installed binary or run this
command with the `--debug` flag.

Theoretical possibilities are:

```text
386
amd64
arm
arm64
loong64
mips
mips64
mips64le
mipsle
ppc64
ppc64le
riscv64
s390x
wasm
```

however, there are available binaries for only some of these values.

### cli

#### age

Compare against the last modified date of a command's file.

```bash
is cli age tmux lt 18 hours
```

Don't let `goimports` get more than a week out of date.

```bash
is cli age goimports gt 7 days && go install golang.org/x/tools/cmd/goimports@latest
```

Update `shfmt` basically daily, but only if `go` is installed.

```bash
if is there go; then
    if is cli age shfmt gt 18 hours; then
        go install mvdan.cc/sh/v3/cmd/shfmt@latest
    fi
else
    echo "Go not found. Not installing shfmt"
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

The `--debug` flag can give us some helpful information when troubleshooting
date math.

#### version

Compare versions of available commands. Returns exit code of 0 if condition is
true and exit code of 1 if condition is false.

```bash
is cli version go gte 1.20.4 || bash upgrade-go.sh
```

##### version segments --major | --minor | --patch

If we want to match only on part of a version, we can do that by specifying
which segment of the version we want to match on, where the pattern is
`major.minor.patch`. Given a version number of `1.2.3` that would make 1 the
`major` version, 2 the `minor` version and `3` the patch version. If the
version number does not include a `patch`, then it is assumed to be zero.
Likewise for a missing `minor` segment.

The segment flags are `--major`, `--minor` and `--patch`. They are mutually
exclusive, only one of these flags can be passed per command.

Let's rework the example above to say that any version of `go` with a minor
version >= 20 means we don't need to upgrade.

```bash
is cli version --minor go gte 20 || bash upgrade-go.sh
```

We could also express this in the following way, if we want to match on the
major version as well.

```bash
is cli version go gte 1.20
```

Supported comparisons are:

* `lt`
* `lte`
* `eq`
* `gte`
* `gt`
* `in`
* `ne`
* `like`
* `unlike`

#### output

Run an arbitrary command and compare the output of `stdout`, `stderr` or
`combined`. Whitespace is automatically trimmed from the left and right of
output before any comparisons are attempted. So, we don't need to worry about
trimming output with leading spaces like this:

```text
cat README.md | wc -l
     847
```

The format for this command is:

```text
is cli output              \
  [stdout|stderr|combined] \
  some-command             \
  --arg foo --arg bar      \
  [lt|lte|eq|gte|ne|like|unlike] "string/regex to match"
```

##### stdout

```bash
is cli output stdout date like Wed
```

Case insensitive:

```bash
is cli output stdout date like "(?i)wed"
```

##### stderr

`ssh` prints its version to `stderr`.

```text
is cli output stderr ssh --arg="-V" like 9.0p1
```

##### combined

`combined` allows us to match on `stdout` and `stderr` at the same time.

```text
is cli output combined ssh --arg="-V" like 9.0p1
```

##### ---arg (-a)

Optional argument to command. Can be used more than once.

Let's match on the results of `uname -m -n`.

```
is cli output stdout uname --arg="-m" --arg="-n" eq "olafs-mbp-2.lan x86_64"
```

If our args don't contain special characters or spaces, we may not need to
quote them. Let's match on the results of `cat README.md`.

```
is cli output stdout cat --arg README.md like "an inspector for your environment"
```

##### --compare

Optional argument to command. Defaults to `optimistic`. Because comparisons
like `eq` mean different things when comparing strings, integers and floats, we
can tell `is` what sort of a comparison to perform. Our options are:

* float
* integer
* string
* version
* optimistic

`optimistic` will first try a `string` comparison. If this fails, it will try a
`version` comparison. This will "Do What I Mean" in a lot of cases, but if we
want to constrain the check to a specific type, we can certainly do that.

```text
is cli output stdout                            \
  bash --arg="-c" --arg="cat README.md | wc -l" \
  gt 10                                         \
  --debug --compare integer
```

#### Tip: Using pipes

To pipe output from one command to another, we'll need to do something that is
equivalent to: `bash -c "some-command | other-command"`

To count the number of lines returned by `date`, we might normally write:

```text
$ date | wc -l
       1
```

Via `bash -c`:

```text
$ bash -c "date | wc -l"
       1
```

Now, run via `is` and assert that there really is just one line:

```text
is cli output stdout bash --arg='-c' --arg="date|wc -l" eq 1
```

Let's make this more succinct. We can make this a little shorter, because `is`
handles `bash -c` as a special case:

```text
is cli output stdout "bash -c" -a "date|wc -l" eq 1
```

#### Tip: Using Negative Numbers

Passing negative integers as expected values is a bit tricky, since we don't
want them to be interpreted as flags.

```
$ is cli output stdout 'bash -c' -a 'date|wc -l' gt -1
```

> 💥 is: error: unknown flag -1, did you mean one of "-h", "-a"?

We can use `--` before the expected value to get around this. 😅

```
$ is cli output stdout 'bash -c' -a 'date|wc -l' gt -- -1
```

##### --debug

Use the `--debug` flag to see where comparisons are failing:

```text
is cli output stdout uname --arg="-m" --arg="-n" eq "olafs-mbp-2.lan x86_65" --debug
2023/09/13 23:05:26 comparison "olafs-mbp-2.lan x86_64" eq "olafs-mbp-2.lan x86_65"
2023/09/13 23:05:26 comparison failed: olafs-mbp-2.lan x86_64 eq olafs-mbp-2.lan x86_65
```

Supported comparisons are:

* `lt`
* `lte`
* `eq`
* `gte`
* `gt`
* `in`
* `ne`
* `like`
* `unlike`

👉 Nota bene: because `is` doesn't know what you're trying to match, it will,
in some cases try to do an optimistic comparison. That is, it will try a string
comparison first and then a numeric comparison. Hopefully this will "do the
right thing" for you. If not, please open an issue.

### os

Information specific to the current operating system

#### version

```bash
is os version gt 22
```

```bash
is os version like "13.4.\d"
```

##### version segments --major | --minor | --patch

If we want to match only on part of a version, we can do that by specifying
which segment of the version we want to match on, where the pattern is
`major.minor.patch`. Given a version number of `1.2.3` that would make 1 the
`major` version, 2 the `minor` version and `3` the patch version. If the
version number does not include a `patch`, then it is assumed to be zero.
Likewise for a missing `minor` segment.

The segment flags are `--major`, `--minor` and `--patch`. They are mutually
exclusive, only one of these flags can be passed per command.

```text
is os version --major eq 13
```

Supported comparisons are:

* `lt`
* `lte`
* `eq`
* `gte`
* `gt`
* `in`
* `ne`
* `like`
* `unlike`

#### name

Under the hood, this returns the value of `runtime.GOOS`, a constant which is
set at compile time. So, `is` reports on on the OS name which is the target of
build rather than running `uname` or something like that. This is "good enough"
for my purposes. If it's not good enough for yours, we probably need to add
more build targets.

Available comparisons are:

* `eq`
* `ne`
* `in`
* `like`
* `unlike`

##### Equality

```bash
is os name eq darwin
```

##### Inequality

```bash
is os name ne linux
```

##### In a comma-delimited list

```bash
is os name in darwin,linux
```

##### Regex

```bash
is os name like darw
```

```bash
is os name like "dar\w{3}"
```

```bash
is os name unlike "foo\d"
```

Possible values for `name`:

```text
darwin
linux
```

##### pretty-name

Linux only.

```bash
is os pretty-name eq "Ubuntu 22.04.2 LTS"
```

Available comparisons are:

* `eq`
* `ne`
* `in`
* `like`
* `unlike`

##### id

Linux only.

```bash
is os id eq ubuntu
```

Available comparisons are:

* `eq`
* `ne`
* `in`
* `like`
* `unlike`

##### id-like

```bash
is os id-like eq debian
```

Available comparisons are:

* `eq`
* `ne`
* `in`
* `like`
* `unlike``

##### version-codename

```bash
is os version-codename eq jammy
```

Available comparisons are:

* `eq`
* `ne`
* `in`
* `like`
* `unlike`

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

```bash
is there tmux && echo "we have tmux"
```

### user

#### sudoer

```bash
is user sudoer && sudo apt-get install ripgrep
```

Returns 1 if the current appears to be able to `sudo` without being prompted
for a password.

This is useful for scripts where we want to install via `sudo`, but we don't
want the script to be interactive. That means we can skip installing things
that require `sudo` and handle them in some other place.

### known

Prints known information about a resource to `STDOUT`. Returns `0` on success
and `1` if info cannot be found.

#### arch

Prints the value of golang's `runtime.GOARCH`. Note that this is the arch that
the binary was compiled for. It's not running `uname` under the hood.

Theoretical possibilities are:

```text
386
amd64
arm
arm64
loong64
mips
mips64
mips64le
mipsle
ppc64
ppc64le
riscv64
s390x
wasm
```

however, there are available binaries for only some of these values.

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

```text
$ is known os version --major
22
```

```text
$ is known os version --minor
04
```

```text
$ is known os version --patch
0
```

Please see the docs on `os version` for more information on `--major`,
`--minor` and `--patch`.

##### version-codename

```text
$ is known os version-codename
jammy
```

#### cli version

```text
$ is known cli version tmux
2.7
```

```text
$ is known cli version --major tmux
2
```

```text
$ is known cli version --minor tmux
7
```

```text
$ is known cli version --patch tmux
0
```
Please see the docs on `os version` for more information on `--major`,
`--minor` and `--patch`.

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
  arch <op> <val>
    Check arch e.g. "is arch like x64"

  cli version <name> <op> <val>
    Check version of command. e.g. "is cli version tmux gte 3"

  cli age <name> <op> <val> <unit>
    Check last modified time of cli (2h, 4d). e.g. "is cli age tmux gt 1 d"

  cli output <stream> <command> <op> <val>
    Check output of a command. e.g. "is cli output stdout "uname -a" like
    "Kernel Version 22.5"

  known arch [<attr>]
    Print arch without check. e.g. "is known arch"

  known os <attribute>
    Print without check. e.g. "is known os name"

  known cli <attribute> <name>
    Print without check. e.g. "is known cli version git"

  os <attribute> <op> <val>
    Check OS attributes. e.g. "is os name eq darwin"

  there <name>
    Check if command exists. e.g. "is there git"

  user [<sudoer>]
    Info about current user. e.g. "is user sudoer"

Run "is <command> --help" for more information on a command.
```

#### subcommand --help

```text
Usage: is os <attribute> <op> <val>

Check OS attributes. e.g. "is os name eq darwin"

Arguments:
  <attribute>    [id|id-like|pretty-name|name|version|version-codename]
  <op>           [eq|ne|gt|gte|in|like|lt|lte|unlike]
  <val>

Flags:
  -h, --help       Show context-sensitive help.
      --debug      turn on debugging statements
      --version    Print version to screen

      --major      Only match on the major OS version (e.g. major.minor.patch)
      --minor      Only match on the minor OS version (e.g. major.minor.patch)
      --patch      Only match on the patch OS version (e.g. major.minor.patch)
```

### --version

Print current version of `is`

```bash
is --version
```

## Installation

Choose from the following options to install `is`.

1. [Download a release](https://github.com/oalders/is/releases)
1. Use `go install`
  * `go install github.com/oalders/is@latest`
  * `go install github.com/oalders/is@v0.4.3`
1. Use [ubi](https://github.com/houseabsolute/ubi)

```bash
#!/usr/bin/env bash

set -e -u -x -o pipefail

# Or choose a different dir in your $PATH
dir="$HOME/local/bin"

if [ ! "$(command -v ubi)" ]; then
    curl --silent --location \
        https://raw.githubusercontent.com/houseabsolute/ubi/master/bootstrap/bootstrap-ubi.sh |
        TARGET=$dir sh
fi

ubi --project oalders/is --in "$dir"
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
