package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaybeReadINI(t *testing.T) {
	{
		release, err := maybeReadINI("testdata/etc/os-release")
		assert.NoError(t, err)
		assert.Equal(t, "18.04", release.VersionID)
	}
	{
		// if the file does not exist on this system, that's not an error
		release, err := maybeReadINI("testdata/etc/os-releasezzz")
		assert.NoError(t, err)
		assert.Nil(t, release)
	}
	{
		// if the file cannot be parsed, that's an error
		release, err := maybeReadINI("testdata/etc/not-an-ini-file")
		assert.Error(t, err)
		assert.Nil(t, release)
	}
}
