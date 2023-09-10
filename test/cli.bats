#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

@test "cli age" {
    ./is cli age tmux gt 1 s
    run ! ./is cli age tmux lt 1 s
}

@test "cli unimplemented subcommand" {
    run ! ./is cli Zage tmux gt 1 s
}

@test "output" {
    ./is cli output stdout date like "\d"
}

@test "output with pipe" {
    ./is cli output stdout bash --arg='-c' --arg='date|wc -l' eq 1
}
