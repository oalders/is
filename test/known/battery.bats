#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

setup() {
    if ./is battery count eq 0; then
        skip "Skipping battery tests as no battery found"
    fi
}

@test "is known battery" {
    run ./is known battery
    [ "${status}" -eq 1 ]
    [[ "${output}" == *"Usage: is known battery <attribute>"* ]]
}

@test "is known battery charge-rate --nth 0" {
    run ./is known battery charge-rate --nth 0
    [ "${status}" -eq 1 ]
    [[ "${output}" == *"use --nth 1 to get the first battery"* ]]
}

@test "is known battery count" {
    ./is known battery count
}

@test "is known battery charge-rate --round" {
    ./is known battery charge-rate --round
}

@test "is known battery current-capacity" {
    ./is known battery current-capacity
}

@test "is known battery current-charge" {
    ./is known battery current-charge
}

@test "is known battery design-capacity" {
    ./is known battery design-capacity
}

@test "is known battery design-voltage" {
    ./is known battery design-voltage
}

@test "is known battery last-full-capacity" {
    ./is known battery last-full-capacity
}

@test "is known battery state" {
    ./is battery count gt 0 && ./is known battery state
}

@test "is known battery voltage" {
    ./is known battery voltage
}

@test "is known summary battery" {
    ./is known summary battery
}

@test "is known summary battery --json" {
    ./is known summary battery --json
}
