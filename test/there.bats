#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

tmux=./testdata/bin/tmux

@test "is there tmux" {
    ./is there $tmux
    run ! ./is there tmuxxx
}

@test "non-zero when cli does not exist" {
    ./is there $tmux
    run ! ./is there tmuxxx
}

@test "is there bash" {
    run ./is there bash --verbose
    echo $status
    [ "$status" -eq 0 ]
}

@test "is there bash --verbose" {
    run ./is there bash --verbose
    echo $status
    [ "$status" -eq 0 ]
}

@test "is there bash --json" {
    run ./is there bash --json
    echo $status
    [ "$status" -eq 0 ]
}

@test "is there bash --all" {
    run ./is there bash --all
    echo $status
    [ "$status" -eq 0 ]
}

@test "is there bash --json --all" {
    run ./is there bash --json --all
    echo $status
    [ "$status" -eq 0 ]
}
