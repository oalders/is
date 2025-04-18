#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

@test "is os" {
  ./is os name like ".*"
  run ! ./is os name unlike ".*"
}

@test "is os name in" {
  ./is os name in darwin,linux
}

@test "is os version gt 0 --debug" {
  ./is os version gt 0 --debug
}

@test "is os version --major gt 0" {
  ./is os version --major gt 0
}

@test "is os version --minor gt 0" {
  ./is os version --minor gt 0
}

@test "is os id --major gt 0" {
  run ./is os id --major gt 0
  [[ $status -ne 0 ]]
  [[ "$output" == *"--major can only be used with version"* ]]
}

@test "is os id --minor gt 0" {
  run ./is os id --minor gt 0
  [[ $status -ne 0 ]]
  [[ "$output" == *"--minor can only be used with version"* ]]
}

@test "is os id --patch gt 0" {
  run ./is os id --patch gt 0
  [[ $status -ne 0 ]]
  [[ "$output" == *"--patch can only be used with version"* ]]
}
