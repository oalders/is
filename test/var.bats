#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

setup() {
  export TEST_VAR="nvim"
  export NUM_VAR="42"
  export FLOAT_VAR="3.14"
}

teardown() {
  unset TEST_VAR
  unset NUM_VAR
  unset FLOAT_VAR
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

@test "is var NUM_VAR eq 42 --compare integer" {
  run ./is var NUM_VAR eq 42 --compare integer
  [ "$status" -eq 0 ]
}

@test "is var NUM_VAR gt 40 --compare integer" {
  run ./is var NUM_VAR gt 40 --compare integer
  [ "$status" -eq 0 ]
}

@test "is var NUM_VAR lt 50 --compare integer" {
  run ./is var NUM_VAR lt 50 --compare integer
  [ "$status" -eq 0 ]
}

@test "is var FLOAT_VAR eq 3.14 --compare float" {
  run ./is var FLOAT_VAR eq 3.14 --compare float
  [ "$status" -eq 0 ]
}

@test "is var FLOAT_VAR gt 3.0 --compare float" {
  run ./is var FLOAT_VAR gt 3.0 --compare float
  [ "$status" -eq 0 ]
}

@test "is var FLOAT_VAR lt 4.0 --compare float" {
  run ./is var FLOAT_VAR lt 4.0 --compare float
  [ "$status" -eq 0 ]
}

@test "is var FLOAT_VAR eq 3.14 --compare integer should fail" {
  run ./is var FLOAT_VAR eq 3.14 --compare integer
  [ "$status" -eq 1 ]
  [[ "$output" == *"wanted result must be an integer"* ]]
}

@test "is var EMPTY_VAR eq '' should pass" {
  export EMPTY_VAR=""
  run ./is var EMPTY_VAR eq ""
  [ "$status" -eq 0 ]
}

# we can't determine the difference between eq "" and eq with no trailing arg
@test "'is var EMPTY_VAR eq' should pass" {
  export EMPTY_VAR=""
  run ./is var EMPTY_VAR eq
  [ "$status" -eq 0 ]
}

@test "is var EMPTY_VAR eq \"\" should fail" {
  run ./is var EMPTY_VAR eq ""
  [ "$status" -eq 1 ]
}

@test "is var EMPTY_VAR set should pass" {
  export EMPTY_VAR=""
  run ./is var EMPTY_VAR set
  [ "$status" -eq 0 ]
}

@test "is var EMPTY_VAR unset should fail" {
  export EMPTY_VAR=""
  run ./is var EMPTY_VAR unset
  [ "$status" -eq 1 ]
}

@test "is var EMPTY_VAR ne '' should fail" {
  export EMPTY_VAR=""
  run ./is var EMPTY_VAR ne ""
  [ "$status" -eq 1 ]
}

# we can't determine the difference between ne "" and ne with no trailing arg
@test "is var EMPTY_VAR ne should fail" {
  export EMPTY_VAR=""
  run ./is var EMPTY_VAR ne
  [ "$status" -eq 1 ]
}

@test "is var EMPTY_VAR ne 'non-empty' should pass" {
  export EMPTY_VAR=""
  run ./is var EMPTY_VAR ne "non-empty"
  [ "$status" -eq 0 ]
}
