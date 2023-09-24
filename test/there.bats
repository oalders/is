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
