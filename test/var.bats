#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

setup() {
  export TEST_VAR="nvim"
  export NUM_VAR="42"
  export FLOAT_VAR="3.14"
  export BOOL_TRUE_VAR="true"
  export BOOL_FALSE_VAR="false"
  export BOOL_ONE_VAR="1"
  export BOOL_ZERO_VAR="0"
  export BOOL_T_VAR="t"
  export BOOL_F_VAR="f"
  export BOOL_TRUE_UPPER_VAR="TRUE"
  export BOOL_FALSE_UPPER_VAR="FALSE"
  export BOOL_INVALID_VAR="maybe"
}

teardown() {
  unset TEST_VAR
  unset NUM_VAR
  unset FLOAT_VAR
  unset BOOL_TRUE_VAR
  unset BOOL_FALSE_VAR
  unset BOOL_ONE_VAR
  unset BOOL_ZERO_VAR
  unset BOOL_T_VAR
  unset BOOL_F_VAR
  unset BOOL_TRUE_UPPER_VAR
  unset BOOL_FALSE_UPPER_VAR
  unset BOOL_INVALID_VAR
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

# Tests for the new 'true' subcommand
@test "is var BOOL_TRUE_VAR true should pass" {
  run ./is var BOOL_TRUE_VAR true
  [ "$status" -eq 0 ]
}

@test "is var BOOL_FALSE_VAR true should fail" {
  run ./is var BOOL_FALSE_VAR true
  [ "$status" -eq 1 ]
}

@test "is var BOOL_ONE_VAR true should pass" {
  run ./is var BOOL_ONE_VAR true
  [ "$status" -eq 0 ]
}

@test "is var BOOL_ZERO_VAR true should fail" {
  run ./is var BOOL_ZERO_VAR true
  [ "$status" -eq 1 ]
}

@test "is var BOOL_T_VAR true should pass" {
  run ./is var BOOL_T_VAR true
  [ "$status" -eq 0 ]
}

@test "is var BOOL_F_VAR true should fail" {
  run ./is var BOOL_F_VAR true
  [ "$status" -eq 1 ]
}

@test "is var BOOL_TRUE_UPPER_VAR true should pass" {
  run ./is var BOOL_TRUE_UPPER_VAR true
  [ "$status" -eq 0 ]
}

@test "is var BOOL_FALSE_UPPER_VAR true should fail" {
  run ./is var BOOL_FALSE_UPPER_VAR true
  [ "$status" -eq 1 ]
}

@test "is var BOOL_INVALID_VAR true should fail with error" {
  run ./is var BOOL_INVALID_VAR true
  [ "$status" -eq 1 ]
  [[ "$output" == *"cannot be parsed as boolean"* ]]
}

@test "is var NON_EXISTENT_VAR true should fail with error" {
  run ./is var NON_EXISTENT_VAR true
  [ "$status" -eq 1 ]
  [[ "$output" == *"is not set"* ]]
}

# Tests for the new 'false' subcommand
@test "is var BOOL_TRUE_VAR false should fail" {
  run ./is var BOOL_TRUE_VAR false
  [ "$status" -eq 1 ]
}

@test "is var BOOL_FALSE_VAR false should pass" {
  run ./is var BOOL_FALSE_VAR false
  [ "$status" -eq 0 ]
}

@test "is var BOOL_ONE_VAR false should fail" {
  run ./is var BOOL_ONE_VAR false
  [ "$status" -eq 1 ]
}

@test "is var BOOL_ZERO_VAR false should pass" {
  run ./is var BOOL_ZERO_VAR false
  [ "$status" -eq 0 ]
}

@test "is var BOOL_T_VAR false should fail" {
  run ./is var BOOL_T_VAR false
  [ "$status" -eq 1 ]
}

@test "is var BOOL_F_VAR false should pass" {
  run ./is var BOOL_F_VAR false
  [ "$status" -eq 0 ]
}

@test "is var BOOL_TRUE_UPPER_VAR false should fail" {
  run ./is var BOOL_TRUE_UPPER_VAR false
  [ "$status" -eq 1 ]
}

@test "is var BOOL_FALSE_UPPER_VAR false should pass" {
  run ./is var BOOL_FALSE_UPPER_VAR false
  [ "$status" -eq 0 ]
}

@test "is var BOOL_INVALID_VAR false should fail with error" {
  run ./is var BOOL_INVALID_VAR false
  [ "$status" -eq 1 ]
  [[ "$output" == *"cannot be parsed as boolean"* ]]
}

@test "is var NON_EXISTENT_VAR false should fail with error" {
  run ./is var NON_EXISTENT_VAR false
  [ "$status" -eq 1 ]
  [[ "$output" == *"is not set"* ]]
}
