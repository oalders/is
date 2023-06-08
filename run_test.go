package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThereCmd(t *testing.T) {
	ctx := Context{Debug: true}

	{
		cmd := ThereCmd{Name: "cat"}
		info := meta{}
		err := cmd.Run(&ctx, &info)
		assert.NoError(t, err)
		assert.True(t, info.Success)
	}
	{
		cmd := ThereCmd{Name: "catzzzzz"}
		info := meta{}
		err := cmd.Run(&ctx, &info)
		assert.NoError(t, err)
		assert.False(t, info.Success)
	}
}

// func (r *OSCmd) Run(ctx *Context, info *meta) error {

func TestOSCmd(t *testing.T) {
	ctx := Context{Debug: true}

	{
		cmd := OSCmd{}
		cmd.Name.Op = "eq"
		cmd.Name.Val = "zzz"
		info := meta{}
		err := cmd.Run(&ctx, &info)
		assert.NoError(t, err)
		assert.False(t, info.Success)
	}
	{
		cmd := OSCmd{}
		cmd.Name.Op = "ne"
		cmd.Name.Val = "zzz"
		info := meta{}
		err := cmd.Run(&ctx, &info)
		assert.NoError(t, err)
		assert.True(t, info.Success)
	}
}

func TestKnownCmd(t *testing.T) {
	ctx := Context{Debug: true}

	{
		cmd := KnownCmd{}
		cmd.Name.Name = "os"
		cmd.Name.Val = "name"
		info := meta{}
		err := cmd.Run(&ctx, &info)
		assert.NoError(t, err)
		assert.True(t, info.Success)
	}
	if runtime.GOOS == "darwin" {
		{
			cmd := KnownCmd{}
			cmd.Name.Name = "os"
			cmd.Name.Val = "version"
			info := meta{}
			err := cmd.Run(&ctx, &info)
			assert.NoError(t, err)
			assert.True(t, info.Success)
		}
	}
	{
		cmd := KnownCmd{}
		cmd.Name.Name = "command-version"
		cmd.Name.Val = "tmux"
		info := meta{}
		err := cmd.Run(&ctx, &info)
		assert.NoError(t, err)
		assert.True(t, info.Success)
	}
}

func TestCommandCmd(t *testing.T) {
	ctx := Context{Debug: true, Verbose: true}

	{
		cmd := CommandCmd{}
		cmd.Name.Name = "tmux"
		cmd.Name.Val = "1"
		cmd.Name.Op = "ne"
		info := meta{}
		err := cmd.Run(&ctx, &info)
		assert.NoError(t, err)
		assert.True(t, info.Success)
	}

	{
		cmd := CommandCmd{}
		cmd.Name.Name = "tmuxzzz"
		cmd.Name.Val = "1"
		cmd.Name.Op = "ne"
		info := meta{}
		err := cmd.Run(&ctx, &info)
		assert.Error(t, err)
		assert.False(t, info.Success)
	}

	{
		cmd := CommandCmd{}
		cmd.Name.Name = "tmux"
		cmd.Name.Val = "1"
		cmd.Name.Op = "eq"
		info := meta{}
		err := cmd.Run(&ctx, &info)
		assert.NoError(t, err)
		assert.False(t, info.Success)
	}
}

func TestMacCodeName(t *testing.T) {
	assert.Equal(t, "ventura", macCodeName("13.1"))
	assert.Equal(t, "monterey", macCodeName("12.2"))
	assert.Equal(t, "big sur", macCodeName("11.2"))
	assert.Equal(t, "catalina", macCodeName("10.15"))
	assert.Equal(t, "mojave", macCodeName("10.14"))
	assert.Equal(t, "high sierra", macCodeName("10.13"))
	assert.Equal(t, "sierra", macCodeName("10.12"))
	assert.Equal(t, "el capitan", macCodeName("10.11"))
	assert.Equal(t, "yosemite", macCodeName("10.10"))
	assert.Equal(t, "mavericks", macCodeName("10.9"))
	assert.Equal(t, "mountain lion", macCodeName("10.8"))
	assert.Equal(t, "", macCodeName("10.7"))
}
