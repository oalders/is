#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

setup() {
    if ./is battery count gt 0; then
        skip "Skipping the battery tests if a battery has been found"
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

@test "is known battery voltage" {
    ./is known battery voltage
}
