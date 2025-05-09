#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

setup() {
    if ./is battery count gt 0; then
        skip "Skipping these battery if battery present"
    fi
}

@test "is battery" {
    run ./is battery
    [ "${status}" -eq 1 ]
    [[ "${output}" == *"Usage: is battery <attribute> <op> <val>"* ]]
}

@test "is battery count eq 0" {
    ./is battery count eq 0
}

@test "is battery charge-rate eq 0 --round" {
    ./is battery charge-rate eq 0 --round
}

@test "is battery charge-rate eq 0" {
    ./is battery charge-rate eq 0
}

@test "is battery current-capacity eq 0" {
    ./is battery current-capacity eq 0
}

@test "is battery current-charge eq 0" {
    ./is battery current-charge eq 0
}

@test "is battery design-capacity eq 0" {
    ./is battery design-capacity eq 0
}

@test "is battery design-voltage eq 0" {
    ./is battery design-voltage eq 0
}

@test "is battery last-full-capacity eq 0" {
    ./is battery last-full-capacity eq 0
}

@test "is battery state unlike charg" {
    ./is battery state unlike char
}

@test "is battery voltage eq 0 --debug" {
    ./is battery voltage eq 0 --debug
}

# fixme
# @test "is battery --nth 77 voltage eq 0" {
#   run ./is battery --nth 77 voltage eq 0
#   [[ $status -ne 0 ]]
#   [[ "$output" == *"battery 77 requested, but only"* ]]
# }
