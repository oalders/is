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

@test "is known battery count" {
    ./is known battery count
}

@test "is known battery charge-rate" {
    ./is known battery charge-rate
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
    ./is known battery state
}

@test "is known battery voltage" {
    ./is known battery voltage
}
