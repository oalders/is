# is (is) an inspector for your environment

## Why?

I want easier version parsing and I'm tired of having to remember the different
incantations of software versioning.

```
go version
perl --version
tmux -V
```

## Check OS Name

### Equality

```text
is os name eq darwin
```

### Inequality

```text
is os name ne debian
```

## Check Command version

```text
is command go lt 1.20.5

is command go eq 1.20.4

is command go gt 1.20.3

is command go ne 1.20.2
```

## Print Information Without Testing It

```text
is known os name
is known command tmux version
```

## Get Hints in Debug Mode

```text
is os name eq darwins --debug
Comparison failed: darwin eq darwins
```
