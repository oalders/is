#!/bin/bash

set -eux -o pipefail

from="$1"
to="$2"

perl -i -pE "s/$from/$to/g" README.md main.go test/is.bats
