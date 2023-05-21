package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaybeReadINI(t *testing.T) {
	release, err := maybeReadINI("testdata/etc/os-release")
	assert.NoError(t, err)
	assert.Equal(t, "18.04", release.VersionID)
	fmt.Printf("%+v", release)
}
