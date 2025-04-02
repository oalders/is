#!/usr/bin/env bats

semver=./testdata/bin/semver
tmux=./testdata/bin/tmux

bats_require_minimum_version 1.5.0

@test "is known cli" {
    ./is known cli version $tmux
}

@test "is known os" {
    ./is known os name
}

@test "is known arch" {
    ./is known arch
}

@test "ensure something is printed" {
    run -0 ./is known arch
    [ -n "${lines[0]}" ]
}

@test "is known cli version semver" {
    run -0 ./is known cli version $semver
    [ "${lines[0]}" = "1.2.3" ]
}

@test "is known cli version --major semver" {
    run -0 ./is known cli version --major $semver
    [ "${lines[0]}" = "1" ]
}

@test "is known cli version --minor semver" {
    run -0 ./is known cli version --minor $semver
    [ "${lines[0]}" = "2" ]
}

@test "is known cli version --patch semver" {
    run -0 ./is known cli version --patch $semver
    [ "${lines[0]}" = "3" ]
}

@test "is known cli version --major --minor semver" {
    run ! ./is known cli version --major --minor $semver
}

@test "is known os version --major" {
    ./is known os version --major
}

@test "! is known os name --minor" {
    run ! ./is known os name --minor
}

@test "! is known os name --patch" {
    run ! ./is known os name --patch
}

@test "! is known os version --major" {
    run ! ./is known os name --major
}

@test "is known os version --minor" {
    ./is known os version --minor
}

@test "is known os version --patch" {
    ./is known os version --patch
}
