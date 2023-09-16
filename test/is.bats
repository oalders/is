#!/usr/bin/env bats

@test "is --version" {
  ./is --version
}

@test "is --help" {
  ./is --help
}

@test "is -h" {
  ./is -h
}
