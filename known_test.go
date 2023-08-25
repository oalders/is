package main

import (
	"runtime"
	"testing"

	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

func TestKnownCmd(t *testing.T) {
	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.OS.Attr = "name"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	if runtime.GOOS == "darwin" {
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.OS.Attr = "version"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.CLI.Attr = "version"
		cmd.CLI.Name = "tmux"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}

	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.CLI.Attr = "version"
		cmd.CLI.Name = "tmuxxx"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success, "No success")
	}
}
