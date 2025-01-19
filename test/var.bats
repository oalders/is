#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

setup() {
  export TEST_VAR="nvim"
}

teardown() {
  unset TEST_VAR
}

@test "is var TEST_VAR set" {
  run ./is var TEST_VAR set
  [ "$status" -eq 0 ]
}

@test "is var TEST_VAR unset" {
  run ./is var TEST_VAR unset
  [ "$status" -eq 1 ]
}

@test "is var TEST_VAR eq nvim" {
  run ./is var TEST_VAR eq nvim
  [ "$status" -eq 0 ]
}

@test "is var TEST_VAR eq vim" {
  run ./is var TEST_VAR eq vim
  [ "$status" -eq 1 ]
}

@test "is var TEST_VAR gt 1" {
  run ./is var TEST_VAR gt 1
  [ "$status" -eq 1 ]
}

@test "is var TEST_VAR like nv.*" {
  run ./is var TEST_VAR like "nv.*"
  [ "$status" -eq 0 ]
}

@test "is var TEST_VAR unlike nv.*" {
  run ./is var TEST_VAR unlike "nv.*"
  [ "$status" -eq 1 ]
}

@test "is var NON_EXISTENT_VAR set" {
  run ./is var NON_EXISTENT_VAR set
  [ "$status" -eq 1 ]
}

@test "is var NON_EXISTENT_VAR unset" {
  run ./is var NON_EXISTENT_VAR unset
  [ "$status" -eq 0 ]
}
