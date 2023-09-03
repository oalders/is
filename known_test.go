package main

import (
	"runtime"
	"testing"

	"github.com/oalders/is/attr"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

func TestKnownCmd(t *testing.T) {
	t.Parallel()
	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.OS.Attr = attr.Name
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	if runtime.GOOS == "darwin" {
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.OS.Attr = attr.Version
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.OS.Attr = attr.Version
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}

	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.OS.Attr = "tmuxxx"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success, "No success")
	}
}
