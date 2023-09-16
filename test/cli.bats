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

@test "succinct output with pipe" {
    ./is cli output stdout "bash -c" -a 'date|wc -l' eq 1
}

@test "output with pipe and --compare integer" {
    ./is cli output stdout --compare integer "bash -c" -a 'date|wc -l' eq 1
}

@test "output with pipe and --compare float" {
    ./is cli output stdout --compare float "bash -c" -a 'date|wc -l' eq 1
}

@test "output with pipe and --compare string" {
    ./is cli output stdout --compare string "bash -c" -a 'date|wc -l' eq 1
}

@test "output with pipe and --compare version" {
    ./is cli output stdout --compare version "bash -c" -a 'date|wc -l' eq 1
}

@test "output with pipe and --compare optimistic" {
    ./is cli output stdout --compare optimistic "bash -c" -a 'date|wc -l' eq 1
}

@test "output gt" {
    ./is cli output stdout "bash -c" -a 'date|wc -l' gt 0
}

@test "output gte" {
    ./is cli output stdout "bash -c" -a 'date|wc -l' gte 1
}

@test "output lt" {
    ./is cli output stdout "bash -c" -a 'date|wc -l' lt 2
}

@test "output lte" {
    ./is cli output stdout "bash -c" -a 'date|wc -l' lte 1
}

@test "output like" {
    ./is cli output stdout "bash -c" -a 'date|wc -l' like 1
}

@test "output unlike" {
    ./is cli output stdout "bash -c" -a 'date|wc -l' unlike 111
}
