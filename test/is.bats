#!/usr/bin/env bats

@test "is --version" {
  run ./is --version
  [ "$status" -eq 0 ]
  [ "$output" = "0.7.0" ]
}

@test "is --help" {
  ./is --help
}

@test "is -h" {
  ./is -h
}
