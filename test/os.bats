#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

@test "is os" {
  ./is os name like ".*"
  run ! ./is os name unlike ".*"
}

@test "is os name in" {
  ./is os name in darwin,linux
}
