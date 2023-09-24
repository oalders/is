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
	const tmux = "testdata/bin/tmux"
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

	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.CLI.Attr = attr.Version
		cmd.CLI.Name = "gitzzz"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success, "No success")
	}

	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.Arch.Attr = "arch"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success, "success")
	}

	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.CLI.Attr = attr.Version
		cmd.CLI.Name = tmux
		cmd.Major = true
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success, "success")
	}

	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.CLI.Attr = attr.Version
		cmd.CLI.Name = tmux
		cmd.Minor = true
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success, "success")
	}

	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.CLI.Attr = attr.Version
		cmd.CLI.Name = tmux
		cmd.Patch = true
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success, "success")
	}

	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.CLI.Attr = attr.Version
		cmd.CLI.Name = tmux
		cmd.Major = true
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success, "success")
	}

	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.CLI.Attr = attr.Name
		cmd.CLI.Name = tmux
		cmd.Major = true
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success, "success")
	}
}
