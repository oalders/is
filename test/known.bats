#!/usr/bin/env bats

@test "is known cli" {
  ./is known cli version tmux
}

@test "is known os" {
  ./is known os name
}
