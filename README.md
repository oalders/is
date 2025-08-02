# is: an inspector for your environment

<!-- vim-markdown-toc GFM -->

* [Introduction](#introduction)
  * [Is the minimum version of this tool available?](#is-the-minimum-version-of-this-tool-available)
  * [Is Neovim the Default Editor?](#is-neovim-the-default-editor)
  * [Is this the target Operating System?](#is-this-the-target-operating-system)
  * [Check the OS with a regex](#check-the-os-with-a-regex)
  * [Is this a recent macOS?](#is-this-a-recent-macos)
  * [Who am I?](#who-am-i)
  * [Do we have go? Then install `goimports`](#do-we-have-go-then-install-goimports)
  * [What's the version of bash?](#whats-the-version-of-bash)
  * [What's the major version of zsh?](#whats-the-major-version-of-zsh)
  * [Has gofumpt been modified in the last week?](#has-gofumpt-been-modified-in-the-last-week)
  * [Has a file been modified in the last hour?](#has-a-file-been-modified-in-the-last-hour)
  * [echo the OS name](#echo-the-os-name)
  * [Get some debugging information about the OS](#get-some-debugging-information-about-the-os)
  * [Pretty-Print Summaries in GitHub Actions](#pretty-print-summaries-in-github-actions)
  * [Can user sudo without a password?](#can-user-sudo-without-a-password)
* [Exit Codes are Everything](#exit-codes-are-everything)
  * [Debugging error Codes](#debugging-error-codes)
  * [Using a Regex](#using-a-regex)
    * [Under the Hood](#under-the-hood)
* [Top Level Commands](#top-level-commands)
  * [audio](#audio)
  * [arch](#arch)
  * [battery](#battery)
    * [state](#state)
    * [current-charge](#current-charge)
    * [count](#count)
    * [charge-rate](#charge-rate)
    * [current-capacity](#current-capacity)
    * [design-capacity](#design-capacity)
    * [design-voltage](#design-voltage)
    * [last-full-capacity](#last-full-capacity)
    * [voltage](#voltage)
    * [--nth](#--nth)
    * [--round](#--round)
  * [cli](#cli)
    * [age](#age)
    * [version](#version)
      * [version segments --major | --minor | --patch](#version-segments---major----minor----patch)
    * [output](#output)
      * [stdout](#stdout)
      * [stderr](#stderr)
      * [combined](#combined)
      * [---arg (-a)](#---arg--a)
      * [--compare](#--compare)
    * [Tip: Using pipes](#tip-using-pipes)
    * [Tip: Using Negative Numbers](#tip-using-negative-numbers)
      * [--debug](#--debug)
  * [fso](#fso)
    * [age](#age-1)
  * [os](#os)
    * [version](#version-1)
      * [version segments --major | --minor | --patch](#version-segments---major----minor----patch-1)
    * [name](#name)
      * [Equality](#equality)
      * [Inequality](#inequality)
      * [In a comma-delimited list](#in-a-comma-delimited-list)
      * [Regex](#regex)
      * [pretty-name](#pretty-name)
      * [id](#id)
      * [id-like](#id-like)
      * [version-codename](#version-codename)
  * [there](#there)
    * [--verbose](#--verbose)
    * [--all](#--all)
    * [--json](#--json)
  * [user](#user)
    * [sudoer](#sudoer)
  * [var](#var)
    * [set](#set)
    * [unset](#unset)
      * [--compare](#--compare-1)
  * [known](#known)
  * [audio](#audio-1)
    * [arch](#arch-1)
    * [battery](#battery-1)
    * [os](#os-1)
      * [name](#name-1)
      * [pretty-name](#pretty-name-1)
      * [id](#id-1)
      * [id-like](#id-like-1)
      * [version](#version-2)
      * [version-codename](#version-codename-1)
    * [cli version](#cli-version)
    * [var](#var-1)
    * [summary](#summary)
      * [battery](#battery-2)
      * [os](#os-2)
      * [var](#var-2)
  * [install-completions](#install-completions)
  * [--debug](#--debug-1)
  * [--help](#--help)
    * [subcommand --help](#subcommand---help)
  * [--version](#--version)
* [Installation](#installation)
* [Bonus: Easier Version Parsing of Available Tools](#bonus-easier-version-parsing-of-available-tools)
  * [Go (version)](#go-version)
  * [Perl (--version)](#perl---version)
  * [tmux (-V)](#tmux--v)

<!-- vim-markdown-toc -->

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

### Is Neovim the Default Editor?

```bash
is var EDITOR set && is var EDITOR eq nvim
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
is cli age gofumpt lt 7 d
```

### Has a file been modified in the last hour?

```text
is fso age ./stats.txt lt 1 h
```

### echo the OS name

```text
$ echo "i'm on $(is known os name) for sure"
i'm on darwin for sure
```

### Get some debugging information about the OS

macOS:

```shell
$ is known summary os
┏━━━━━━━━━━━━━━━━━━┳━━━━━━━━━┓
┃ Attribute        ┃ Value   ┃
┣━━━━━━━━━━━━━━━━━━╋━━━━━━━━━┫
┃ name             ┃ darwin  ┃
┣━━━━━━━━━━━━━━━━━━╋━━━━━━━━━┫
┃ version          ┃ 13.7.6  ┃
┣━━━━━━━━━━━━━━━━━━╋━━━━━━━━━┫
┃ version-codename ┃ ventura ┃
┗━━━━━━━━━━━━━━━━━━┻━━━━━━━━━┛
```

```shell
is known summary os --json
```

```json
{
    "name": "darwin",
    "version": "15.4.1",
    "version-codename": "sequoia"
}
```

```shell
is known summary os --md
```


| Attribute | Value |
|---|---|
| name | linux |
| version | 3.22.0 |
| id | alpine |
| pretty-name | Alpine Linux v3.22 |


### Pretty-Print Summaries in GitHub Actions

After you have `is` installed, you can echo markdown tables to the
`$GITHUB_STEP_SUMMARY` environment variable. This will give you a readable
tables that live outside of the log files.

```yaml
jobs:
  linux:
    runs-on: ubuntu-latest
    steps:
      - name: Install ubi
        uses: oalders/install-ubi-action@v0.0.2

      - name: Install "is"
        run: sudo ubi --project oalders/is --in /usr/local/bin

      - name: Display summaries
        run: |
          is known summary os  --md >> $GITHUB_STEP_SUMMARY
          is known summary var --md >> $GITHUB_STEP_SUMMARY

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

Regex patterns are passed directly to Golang's `regexp.MatchString`. We can
take advantage of this when crafting regexes. For instance, for a case
insensitive search:

```text
is cli output stdout date like "(?i)wed"
```

## Top Level Commands

### audio

Checks against known audio attributes, if available. Returns non-zero and
prints an error message if audio cannot be discovered.

Available attributes:

* level
* muted

`level` is a value from 0 to 100.

```shell
is audio level gte 56 && echo "loud enough"
```

```shell
is audio muted && echo "nothing to hear here"
```

### arch

Checks against the arch which this binary has been compiled for.

```text
is arch eq amd64
```

```text
is arch like 64
```

Supported comparisons are:

- `eq`
- `ne`
- `in`
- `like`
- `unlike`

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

### battery

```shell
./is battery --help
Usage: is battery <attribute> <op> <val>

Check battery attributes. e.g. "is battery state eq charging"

Arguments:
  <attribute>    [charge-rate|count|current-capacity|current-charge|design-capacity|design-voltage|last-full-capacity|state|voltage]
  <op>           [eq|ne|gt|gte|in|like|lt|lte|unlike]
  <val>

Flags:
  -h, --help       Show context-sensitive help.
      --debug      turn on debugging statements
      --version    Print version to screen

      --nth=1      Specify which battery to use (1 for the first battery)
      --round      Round float values to the nearest integer
```

Compare against battery attributes for power management. This is useful for writing scripts that depend on battery state or charge level.

```bash
is battery state eq charging
```

```bash
is battery current-charge lt 20 && notify-send "Battery Low" "Battery level below 20%"
```

#### state

Check the current state of the battery:

```bash
is battery state eq charging
```

```bash
is battery state in charging,full
```

```bash
is battery state like discharg
```

Possible states include:

- charging
- discharging
- idle
- empty
- full
- unknown
- undefined

#### current-charge

Check the current charge percentage of the battery:

```bash
is battery current-charge lt 15 && echo "Battery critically low!"
```

#### count

Check the number of batteries detected in the system:

```bash
is battery count gt 0 || echo "No battery found"
```

#### charge-rate

This is the current (momentary) charge rate (in mW). It is always non-negative, check `is battery state` to see whether it means charging or discharging.

See <https://github.com/distatus/battery/blob/master/battery.go>

#### current-capacity

Check the current (momentary) capacity in mWh:

```bash
is battery current-capacity gt 7000
```

#### design-capacity

Check the design capacity in mWh:

```bash
is battery design-capacity gt 7000
```

#### design-voltage

Check the design voltage in mWh:

```bash
is battery design-voltage gt 13
```

#### last-full-capacity

Check the capacity at last full charge in mWh:

```bash
is battery last-full-capacity gt 7000
```

#### voltage

Check the current battery voltage in V:

```bash
is battery voltage gt 10
```

#### --nth

Specify which battery to check if multiple batteries are present:

```bash
is battery --nth=2 state eq charging
```

#### --round

Round float values to the nearest integer:

```bash
is battery --round current-charge eq 85
```

Without `--round`:

```bash
is battery current-charge eq 85.4
```

Supported comparisons are:

- `lt`
- `lte`
- `eq`
- `gte`
- `gt`
- `in`
- `ne`
- `like`
- `unlike`

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

- `lt`
- `gt`

Supported units are:

- `s`
- `second`
- `seconds`
- `m`
- `minute`
- `minutes`
- `h`
- `hour`
- `hours`
- `d`
- `day`
- `days`

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

- `lt`
- `lte`
- `eq`
- `gte`
- `gt`
- `in`
- `ne`
- `like`
- `unlike`

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

```bash
is cli output stdout uname --arg="-m" --arg="-n" eq "olafs-mbp-2.lan x86_64"
```

If our args don't contain special characters or spaces, we may not need to
quote them. Let's match on the results of `cat README.md`.

```bash
is cli output stdout cat --arg README.md like "an inspector for your environment"
```

##### --compare

Optional argument to command. Defaults to `optimistic`. Because comparisons
like `eq` mean different things when comparing strings, integers and floats, we
can tell `is` what sort of a comparison to perform. Our options are:

- float
- integer
- string
- version
- optimistic

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

```bash
$ is cli output stdout 'bash -c' -a 'date|wc -l' gt -1
```

> 💥 is: error: unknown flag -1, did you mean one of "-h", "-a"?

We can use `--` before the expected value to get around this. 😅

```bash
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

- `lt`
- `lte`
- `eq`
- `gte`
- `gt`
- `in`
- `ne`
- `like`
- `unlike`

👉 Nota bene: because `is` doesn't know what you're trying to match, it will,
in some cases try to do an optimistic comparison. That is, it will try a string
comparison first and then a numeric comparison. Hopefully this will "do the
right thing" for you. If not, please open an issue.

### fso

`fso` is short for filesystem object (file, directory, link, etc). This command
is very similar to `cli age`. The difference between `cli age` and `fso age` is
that `fso` will not search your `$PATH`. You may provide either a relative or
an absolute path.

#### age

Compare against the last modified date of a file.

```bash
is fso age /tmp/programs.csv lt 18 hours
```

Compare against the last modified date of a directory.

```bash
is fso age ~./local/cache gt 1 d
```

Supported comparisons are:

- `lt`
- `gt`

Supported units are:

- `s`
- `second`
- `seconds`
- `m`
- `minute`
- `minutes`
- `h`
- `hour`
- `hours`
- `d`
- `day`
- `days`

Note that `d|day|days` is shorthand for 24 hours. DST offsets are not taken
into account here.

The `--debug` flag can give us some helpful information when troubleshooting
date math.

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

- `lt`
- `lte`
- `eq`
- `gte`
- `gt`
- `in`
- `ne`
- `like`
- `unlike`

#### name

Under the hood, this returns the value of `runtime.GOOS`, a constant which is
set at compile time. So, `is` reports on on the OS name which is the target of
build rather than running `uname` or something like that. This is "good enough"
for my purposes. If it's not good enough for yours, we probably need to add
more build targets.

Available comparisons are:

- `eq`
- `ne`
- `in`
- `like`
- `unlike`

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

- `eq`
- `ne`
- `in`
- `like`
- `unlike`

##### id

Linux only.

```bash
is os id eq ubuntu
```

Available comparisons are:

- `eq`
- `ne`
- `in`
- `like`
- `unlike`

##### id-like

```bash
is os id-like eq debian
```

Available comparisons are:

- `eq`
- `ne`
- `in`
- `like`
- `unlike``

##### version-codename

```bash
is os version-codename eq jammy
```

Available comparisons are:

- `eq`
- `ne`
- `in`
- `like`
- `unlike`

On Linux, the value for `version-codename` is taken from `/etc/os-release`. For
Macs, the values are mapped inside this application.

Possible values for Mac:

- ventura
- monterey
- big sur
- catalina
- mojave
- high sierra
- sierra
- el capitan
- yosemite
- mavericks
- mountain lion

### there

Returns exit code of 0 if command exists and exit code of 1 if command cannot
be found.

```bash
is there tmux && echo "we have tmux"
```

#### --verbose

`--verbose` is only useful without the `--all` or `--json` flags, since they
imply `--verbose`. Using this flag you can see the path and version of the
first binary which is in your `$PATH`.


```shell
is there bash --verbose
┏━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━┓
┃ Path                ┃ Version ┃
┣━━━━━━━━━━━━━━━━━━━━━╋━━━━━━━━━┫
┃ /usr/local/bin/bash ┃ 5.2.37  ┃
┗━━━━━━━━━━━━━━━━━━━━━┻━━━━━━━━━┛
```

#### --all

The all switch adds command output, which is an ASCII table of all of
the binaries which have been found, as well as their versions. This is
essentially a wrapper around `which -a`, with added version parsing.

```shell
$ is there bash --all
┏━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━┓
┃ Path                ┃ Version ┃
┣━━━━━━━━━━━━━━━━━━━━━╋━━━━━━━━━┫
┃ /usr/local/bin/bash ┃ 5.2.37  ┃
┣━━━━━━━━━━━━━━━━━━━━━╋━━━━━━━━━┫
┃ /bin/bash           ┃ 3.2.57  ┃
┗━━━━━━━━━━━━━━━━━━━━━┻━━━━━━━━━┛
```

#### --json

```shell
is there bash --json
```

```json
[
    {
        "path": "/opt/homebrew/bin/bash",
        "version": "5.2.37"
    }
]
```

```shell
is there bash --all --json
```

```json
[
    {
        "path": "/opt/homebrew/bin/bash",
        "version": "5.2.37"
    },
    {
        "path": "/bin/bash",
        "version": "3.2.57"
    }
]
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

### var

Check if environment variables are set, unset or contain a specific value.

```bash
is var EDITOR set && is var EDITOR like vim

is var EDITOR set && is var EDITOR eq ""
```

`is var` can only operate on environment variables which are available in its
environment. This means that you may need to export some variables before you
can invoke `is var`. For example, if you have created a variable in the current
script, it should be exported before `is var` is used.

```shell
SLOW_HORSE=rivercartwright
export SLOW_HORSE
is var SLOW_HORSE like cart && echo "blue shirt, white tee"
```

#### set

```bash
is var EDITOR set
```

#### unset

```bash
unset EDITOR
is var EDITOR unset
```

`set` and `unset` don't require arguments, but the comparison operators do.

Supported comparisons are:

- `lt`
- `lte`
- `eq`
- `gte`
- `gt`
- `in`
- `ne`
- `like`
- `unlike`

Both `eq` and `ne` allow for comparisons with an empty string:

```shell
SET_BUT_EMPTY=
export SET_BUT_EMPTY

is var SET_BUT_EMPTY set && is var SET_BUT_EMPTY eq ""

NOT_EMPTY=lebowski
export NOT_EMPTY

is var NOT_EMPTY set && is var NOT_EMPTY ne ""
```

##### --compare

Optional argument to command. Defaults to `optimistic`. Because comparisons
like `eq` mean different things when comparing strings, integers and floats, we
can tell `is` what sort of a comparison to perform. Our options are:

- float
- integer
- string
- version
- optimistic

`optimistic` will first try a `string` comparison. If this fails, it will try a
`version` comparison. This will "Do What I Mean" in a lot of cases, but if we
want to constrain the check to a specific type, we can certainly do that.

💥 alert:

```bash
FLOATER=1.1 is var FLOATER eq 1.1 --compare integer
```

```text
is: error: wanted result must be an integer
```

### known

Prints known information about a resource to `STDOUT`. Returns `0` on success
and `1` if info cannot be found.

### audio

Prints value of some audio attributes, if available. Returns non-zero  and
prints an error message if audio cannot be discovered.

Available attributes:

- level
- muted

`level` prints a value from 0 to 100.

```shell
$ is known audio level
56
```

`muted` prints a string of `true` or `false`

```shell
$ is known audio muted
false
```

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

#### battery

Get information about the system battery without performing a check.

```text
$ is known battery state
charging
```

```text
$ is known battery current-charge
85.4
```

You can specify which battery to query with `--nth`:

```text
$ is known battery --nth=2 state
discharging
```

Round float values to the nearest integer:

```text
$ is known battery --round current-charge
85
```

Available attributes:

* state (one of)
	* 	undefined
	*  unknown
	*  empty
	*  full
	*  charging
	*  discharging
	*  idle
* current-charge (mWh)
* count (int)
* charge-rate (V)
* current-capacity (mWh)
* design-capacity (mWh)
* design-voltage (V)
* last-full-capacity (mWh)
* voltage (V)

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

#### var

```shell
$ is known var PATH
/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
```

This is not so useful on its own, but combined with `--json` it can give you a readable `$PATH`.

```
$ is known var PATH --json
[
    "/go/bin",
    "/usr/local/go/bin",
    "/usr/local/sbin",
    "/usr/local/bin",
    "/usr/sbin",
    "/usr/bin",
    "/sbin",
    "/bin"
]
```

Now we can do something like `is known var PATH --json | jq .[1]` to get the
second item in the list.

#### summary

`summary` is a special subcommand, which aggregates known data for another
subcommand and emits it to `STDOUT` as `json`. This replaces the previous
behaviour of `--debug` for `is known battery` and `is known os`.

##### battery

`is known summary battery`

```json
{
    "state": "discharging",
    "battery-number": 1,
    "count": 1,
    "charge-rate": 5700.2,
    "current-capacity": 69231.52,
    "current-charge": 91,
    "design-capacity": 74620.8,
    "design-voltage": 12.955,
    "last-full-capacity": 75942.21,
    "voltage": 12.955
}
```

##### os

```shell
$ is known summary os
┏━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━┓
┃ Attribute   ┃ Value              ┃
┣━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━┫
┃ name        ┃ linux              ┃
┣━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━┫
┃ version     ┃ 3.22.0             ┃
┣━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━┫
┃ id          ┃ alpine             ┃
┣━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━┫
┃ pretty-name ┃ Alpine Linux v3.22 ┃
┗━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━┛
```

```shell
$ is known summary os
┏━━━━━━━━━━━━━━━━━━┳━━━━━━━━━┓
┃ Attribute        ┃ Value   ┃
┣━━━━━━━━━━━━━━━━━━╋━━━━━━━━━┫
┃ name             ┃ darwin  ┃
┣━━━━━━━━━━━━━━━━━━╋━━━━━━━━━┫
┃ version          ┃ 13.7.6  ┃
┣━━━━━━━━━━━━━━━━━━╋━━━━━━━━━┫
┃ version-codename ┃ ventura ┃
┗━━━━━━━━━━━━━━━━━━┻━━━━━━━━━┛
```

`is known summary os --json`

```json
{
    "id": "alpine",
    "name": "linux",
    "pretty-name": "Alpine Linux v3.22",
    "version": "3.22.0"
}
```

##### var

```shell
is known summary var
```

This will emit your environment variables in a tabular layout. It will split
`PATH` and `MANPATH` on newlines, to make them easier to read.

```shell
┏━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━┓
┃ Name           ┃ Value             ┃
┣━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━┫
┃ GOLANG_VERSION ┃ 1.24.4            ┃
┣━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━┫
┃ GOPATH         ┃ /go               ┃
┣━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━┫
┃ GOTOOLCHAIN    ┃ local             ┃
┣━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━┫
┃ HOME           ┃ /root             ┃
┣━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━┫
┃ HOSTNAME       ┃ bb7d32c277e9      ┃
┣━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━┫
┃ PATH           ┃ /go/bin           ┃
┃                ┃ /usr/local/go/bin ┃
┃                ┃ /usr/local/sbin   ┃
┃                ┃ /usr/local/bin    ┃
┃                ┃ /usr/sbin         ┃
┃                ┃ /usr/bin          ┃
┃                ┃ /sbin             ┃
┃                ┃ /bin              ┃
┣━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━┫
┃ PWD            ┃ /workspace        ┃
┣━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━┫
┃ SHLVL          ┃ 1                 ┃
┣━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━┫
┃ TERM           ┃ xterm             ┃
┗━━━━━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━━━━┛
```

```shell
is known summary var --json
```

```json
{
    "GOLANG_VERSION": "1.24.4",
    "GOPATH": "/go",
    "GOTOOLCHAIN": "local",
    "HOME": "/root",
    "HOSTNAME": "bb7d32c277e9",
    "PATH": [
        "/go/bin",
        "/usr/local/go/bin",
        "/usr/local/sbin",
        "/usr/local/bin",
        "/usr/sbin",
        "/usr/bin",
        "/sbin",
        "/bin"
    ],
    "PWD": "/workspace",
    "SHLVL": "1",
    "TERM": "xterm"
}
```

```shell
is known summary var --md
```


| Name | Value |
|---|---|
| GOLANG_VERSION | 1.24.4 |
| GOPATH | /go |
| GOTOOLCHAIN | local |
| HOME | /root |
| HOSTNAME | 44db2d2f4c32 |
| PATH | /go/bin<br>/usr/local/go/bin<br>/usr/local/sbin<br>/usr/local/bin<br>/usr/sbin<br>/usr/bin<br>/sbin<br>/bin |
| PWD | /workspace |
| SHLVL | 1 |
| TERM | xterm |


### install-completions

This is a two step process. First, run

```bash
is install-completions
```

Then run the command which is printed to your terminal in order to get
completion in your current session.

```bash
$ is install-completions
complete -C /Users/olaf/local/bin/is is
```

Or add the command to a `.bashrc` or similar in order to get completion across
all sessions.

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

an inspector for your environment

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

  fso age <name> <op> <val> <unit>
    Check age (last modified time) of an fso (2h, 4d). e.g. "is fso age
    /tmp/log.txt gt 1 d"

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

  var <name> <op> [<val>]
    Check environment variables. e.g. "is var EDITOR eq nvim"

  install-completions
    install shell completions. e.g. "is install-completions" and then run the
    command which is printed to your terminal to get completion in your current
    session. add the command to a .bashrc or similar to get completion across
    all sessions.

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
  * `go install github.com/oalders/is@v0.10.0`
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
