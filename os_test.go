package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOSInfo(t *testing.T) {
	tests := []string{"arch", "name", "version", "version-codename"}

	for _, v := range tests {
		ctx := Context{Debug: true}
		found, err := osInfo(&ctx, v)
		assert.NoError(t, err, v)
		assert.True(t, ctx.Success, v)
		assert.NotEmpty(t, found, v)
	}

	if runtime.GOOS == "linux" {
		tests := []string{"id", "id-like", "pretty-name"}

		for _, v := range tests {
			ctx := Context{Debug: true}
			found, err := osInfo(&ctx, v)
			assert.NoError(t, err, v)
			assert.True(t, ctx.Success, v)
			assert.NotEmpty(t, found, v)
		}
	}
}
