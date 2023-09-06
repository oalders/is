#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

@test "cli age" {
  ./is cli age tmux gt 1 s
  run ! ./is cli age tmux lt 1 s
}

