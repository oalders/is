#!/usr/bin/env bash

set -eu -o pipefail
version=3.3.0

if is os name ne 'darwin'; then
    exit
fi

if ! is there rbenv; then
    brew install rbenv
fi

# Might need to initialize rbenv in the shell
if [ -z "${RBENV_SHELL:-}" ]; then
    eval "$(rbenv init - bash)"
fi

if is cli output stdout rbenv --arg version like "^$version\b"; then
    echo "Ruby version $version is already installed"
    exit
fi

if ! is cli output stdout rbenv --arg versions like "\b$version\b" --debug; then
    rbenv install $version
fi

rbenv global $version

if ! is cli version ruby eq $version; then
    echo "Ruby version $version is not available"
fi
