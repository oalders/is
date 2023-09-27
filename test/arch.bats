#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

@test "is arch ne xxx" {
  ./is arch ne xxx
}

@test "is arch unlike xxx" {
  ./is arch unlike xxx
}

@test 'is arch like ".*"' {
  ./is arch like ".*"
}

@test 'is arch eq beos' {
  run ! ./is arch eq beos
}

@test 'is arch in' {
  run ./is arch in amd64,arm,arm64
}

@test 'in trims whitespace' {
  run ./is arch in " amd64 , arm , arm64 "
}
