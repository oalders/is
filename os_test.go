package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOSInfo(t *testing.T) {
	tests := []string{"name", "version", "version-codename"}

	for _, v := range tests {
		ctx := Context{Debug: true}
		found, err := osInfo(&ctx, v)
		assert.NoError(t, err, v)
		assert.True(t, ctx.Success, v)
		assert.NotEmpty(t, found, v)
	}

	// id-like not present in Debian 11, so can't be in a blanket Linux test
	if runtime.GOOS == "linux" {
		tests := []string{"id", "pretty-name"}

		for _, v := range tests {
			ctx := Context{Debug: true}
			found, err := osInfo(&ctx, v)
			assert.NoError(t, err, v)
			assert.True(t, ctx.Success, v)
			assert.NotEmpty(t, found, v)
		}
	}
}

func TestOSCmd(t *testing.T) {
	{
		ctx := Context{Debug: true}
		cmd := OSCmd{}
		cmd.Attr = "name"
		cmd.Op = "eq"
		cmd.Val = "zzz"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := Context{Debug: true}
		cmd := OSCmd{}
		cmd.Attr = "name"
		cmd.Op = "ne"
		cmd.Val = "zzz"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := Context{Debug: true}
		cmd := OSCmd{}
		cmd.Attr = "version"
		cmd.Op = "eq"
		cmd.Val = "1"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := Context{Debug: true}
		cmd := OSCmd{}
		cmd.Attr = "version"
		cmd.Op = "ne"
		cmd.Val = "1"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := Context{Debug: false}
		cmd := OSCmd{}
		cmd.Attr = "name"
		cmd.Op = "gte"
		cmd.Val = "1"
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}
}
