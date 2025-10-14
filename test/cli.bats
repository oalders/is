#!/usr/bin/env bats

bats_require_minimum_version 1.5.0

semver=./testdata/bin/semver
tmux=./testdata/bin/tmux

@test 'cli age' {
    ./is cli age $tmux gt 1 s
    run ! ./is cli age $tmux lt 1 s
}

@test 'cli unimplemented subcommand' {
    run ! ./is cli Zage $tmux gt 1 s
}

@test 'output' {
    ./is cli output stdout date like '\d'
}

@test 'output with pipe' {
    ./is cli output stdout bash --arg='-c' --arg='date|wc -l' eq 1
}

@test 'succinct output with pipe' {
    ./is cli output stdout 'bash -c' -a 'date|wc -l' eq 1
}

@test 'output with pipe and --compare integer' {
    ./is cli output stdout --compare integer 'bash -c' -a 'date|wc -l' eq 1
}

@test 'output with pipe and --compare float' {
    ./is cli output stdout --compare float 'bash -c' -a 'date|wc -l' eq 1
}

@test 'output with pipe and --compare string' {
    ./is cli output stdout --compare string 'bash -c' -a 'date|wc -l' eq 1
}

@test 'output with pipe and --compare version' {
    ./is cli output stdout --compare version 'bash -c' -a 'date|wc -l' eq 1
}

@test 'output with pipe and --compare optimistic' {
    ./is cli output stdout --compare optimistic 'bash -c' -a 'date|wc -l' eq 1
}

@test 'output gt' {
    ./is cli output stdout 'bash -c' -a 'date|wc -l' gt 0
}

@test 'output gte' {
    ./is cli output stdout 'bash -c' -a 'date|wc -l' gte 1
}

@test 'output lt' {
    ./is cli output stdout 'bash -c' -a 'date|wc -l' lt 2
}

@test 'output lte' {
    ./is cli output stdout 'bash -c' -a 'date|wc -l' lte 1
}

@test 'output like' {
    ./is cli output stdout 'bash -c' -a 'date|wc -l' like 1
}

@test 'output unlike' {
    ./is cli output stdout 'bash -c' -a 'date|wc -l' unlike 111
}

@test 'output gte negative integer' {
    ./is cli output stdout 'bash -c' -a 'date|wc -l' gte --compare integer -- -1
}

@test 'output in' {
    ./is cli output --debug stdout $semver --arg="--version" in 1.2.3,3.2.1
}

@test 'major' {
    ./is cli version --major $semver eq 1
}

@test 'minor' {
    ./is cli version --minor $semver eq 2
}

@test 'patch' {
    ./is cli version --patch $semver eq 3
}

@test 'version in (float)' {
    ./is cli version $tmux in "3.2,3.3a"
}

@test 'version in' {
    ./is cli version $semver in 1.2.3,1.2.4,1.2.5
}

@test 'version --major in' {
    ./is cli version --major $semver in 1,4,5
}

@test 'unspecified patch in output' {
    ./is cli version --patch $tmux eq 0
}

@test 'string in' {
    ./is cli output stdout date --arg="+%a" in Mon,Tue,Wed,Thu,Fri,Sat,Sun
}

@test 'command with arguments - uname -a' {
    ./is os name ne linux && skip "Linux-only test"
    ./is cli output stdout "uname -a" like "Linux"
}

@test 'command with arguments - uname -m' {
    ./is os name ne linux && skip "Linux-only test"
    ./is cli output stdout "uname -m" like "x86_64"
}
