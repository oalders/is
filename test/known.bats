#!/usr/bin/env bats

tmux=./testdata/bin/tmux

@test "is known cli" {
	./is known cli version $tmux
}

@test "is known os" {
	./is known os name
}

@test "is known arch" {
	./is known arch
}

bats_require_minimum_version 1.5.0

@test "ensure something is printed" {
	run -0 ./is known arch
	[ -n "${lines[0]}" ]
}
