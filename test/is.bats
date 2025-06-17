#!/usr/bin/env bats

@test "is --version" {
  run ./is --version
  [ "$status" -eq 0 ]
  [ "$output" = "0.8.2" ]
}

@test "is --help" {
  ./is --help
}

@test "is -h" {
  ./is -h
}
