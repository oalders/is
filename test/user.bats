#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

@test "is user xxx" {
  run ! ./is user xxx
}
