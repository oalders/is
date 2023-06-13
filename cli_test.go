package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCliCmd(t *testing.T) {
	{
		ctx := Context{Debug: true}
		cmd := CLICmd{}
		cmd.Version.Name = "tmux"
		cmd.Version.Op = "ne"
		cmd.Version.Val = "1"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}

	{
		ctx := Context{Debug: true}
		cmd := CLICmd{}
		cmd.Version.Name = "tmuxzzz"
		cmd.Version.Op = "ne"
		cmd.Version.Val = "1"
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}

	{
		ctx := Context{Debug: true}
		cmd := CLICmd{}
		cmd.Version.Name = "tmux"
		cmd.Version.Op = "eq"
		cmd.Version.Val = "1"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success)
	}

	{
		ctx := Context{Debug: true}
		cmd := CLICmd{}
		cmd.Version.Name = "tmux"
		cmd.Version.Op = "eq"
		cmd.Version.Val = "zzz"
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}
}
