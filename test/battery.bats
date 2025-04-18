#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

setup() {
    if ./is battery count eq 0; then
        skip "Skipping battery tests as no battery found"
    fi
}

@test "is battery" {
    run ./is battery
    [ "${status}" -eq 1 ]
    [[ "${output}" == *"Usage: is battery <attribute> <op> <val>"* ]]
}

@test "is battery count gt 0" {
    ./is battery count gt 0
}

@test "is battery charge-rate gt 0 --round" {
    ./is battery charge-rate gt 0 --round
}

@test "is battery charge-rate gt 0" {
    ./is battery charge-rate gt 0
}

@test "is battery current-capacity gt 0" {
    ./is battery current-capacity gt 0
}

@test "is battery current-charge gt 0" {
    ./is battery current-charge gt 0
}

@test "is battery design-capacity gt 0" {
    ./is battery design-capacity gt 0
}

@test "is battery design-voltage gt 0" {
    ./is battery design-voltage gt 0
}

@test "is battery last-full-capacity gt 0" {
    ./is battery last-full-capacity gt 0
}

@test "is battery state like charg" {
    ./is battery state like char
}

@test "is battery voltage gt 0 --debug" {
    ./is battery voltage gt 0 --debug
}

@test "is battery --nth 77 voltage gt 0" {
  run ./is battery --nth 77 voltage gt 0
  [[ $status -ne 0 ]]
  [[ "$output" == *"battery 77 requested, but only"* ]]
}
