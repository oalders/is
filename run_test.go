package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThereCmd(t *testing.T) {
	{
		ctx := Context{Debug: true}
		cmd := ThereCmd{Name: "cat"}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := Context{Debug: true}
		cmd := ThereCmd{Name: "catzzzzz"}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success)
	}
}

// func (r *OSCmd) Run(ctx *Context, info *meta) error {

func TestOSCmd(t *testing.T) {
	{
		ctx := Context{Debug: true}
		cmd := OSCmd{}
		cmd.Name.Op = "eq"
		cmd.Name.Val = "zzz"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := Context{Debug: true}
		cmd := OSCmd{}
		cmd.Name.Op = "ne"
		cmd.Name.Val = "zzz"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
}

func TestKnownCmd(t *testing.T) {
	{
		ctx := Context{Debug: true}
		cmd := KnownCmd{}
		cmd.OS.Attr = "name"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	if runtime.GOOS == "darwin" {
		ctx := Context{Debug: true}
		cmd := KnownCmd{}
		cmd.OS.Attr = "version"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := Context{Debug: true}
		cmd := KnownCmd{}
		cmd.CLI.Attr = "version"
		cmd.CLI.Name = "tmux"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}

	{
		ctx := Context{Debug: true}
		cmd := KnownCmd{}
		cmd.CLI.Attr = "version"
		cmd.CLI.Name = "tmuxxx"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success, "No success")
	}
}

func TestCommandCmd(t *testing.T) {
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

func TestMacVersion(t *testing.T) {
	version, err := macVersion()
	if runtime.GOOS == "darwin" {
		assert.NotEmpty(t, version)
		assert.NoError(t, err)
	} else {
		assert.Empty(t, version)
		assert.Error(t, err)
	}
}
