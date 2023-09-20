package main

import (
	"testing"

	"github.com/oalders/is/ops"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

const tmux = "./testdata/bin/tmux"

func TestCliVersion(t *testing.T) {
	t.Parallel()
	type test struct {
		Cmp     VersionCmp
		Error   bool
		Success bool
	}

	major := false
	minor := false
	patch := false

	//nolint:godox
	tests := []test{
		{VersionCmp{tmux, ops.Ne, "1", major, minor, patch}, false, true},
		{VersionCmp{"tmuxzzz", ops.Ne, "1", major, minor, patch}, true, false},
		{VersionCmp{tmux, ops.Eq, "1", major, minor, patch}, false, false},
		{VersionCmp{tmux, ops.Eq, "zzz", major, minor, patch}, true, false},
		{VersionCmp{tmux, ops.Unlike, "zzz", major, minor, patch}, false, true},
		{VersionCmp{tmux, ops.Like, "", major, minor, patch}, false, true}, // FIXME
		{VersionCmp{tmux, ops.Like, "3.*", major, minor, patch}, false, true},
		{VersionCmp{tmux, ops.Eq, "3", true, minor, patch}, false, true},
		{VersionCmp{tmux, ops.Eq, "3", major, true, patch}, false, true},
		{VersionCmp{tmux, ops.Eq, "0", major, minor, true}, false, true},
	}

	for _, test := range tests {
		ctx := types.Context{Debug: true}
		cmd := CLICmd{Version: test.Cmp}
		err := cmd.Run(&ctx)
		if test.Error {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		if test.Success {
			assert.True(t, ctx.Success)
		} else {
			assert.False(t, ctx.Success)
		}
	}
}

func TestCliAge(t *testing.T) {
	t.Parallel()
	{
		ctx := types.Context{Debug: true}
		cmd := CLICmd{Age: AgeCmp{tmux, ops.Gt, "1", "s"}}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: true}
		cmd := CLICmd{Age: AgeCmp{tmux, ops.Lt, "100000", "days"}}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: true}
		cmd := CLICmd{Age: AgeCmp{tmux, ops.Lt, "1.1", "d"}}
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: true}
		cmd := CLICmd{Age: AgeCmp{"tmuxxx", ops.Lt, "1", "d"}}
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}
}

func TestCliOutput(t *testing.T) {
	t.Parallel()
	type test struct {
		Cmp     OutputCmp
		Error   bool
		Success bool
	}

	command := "./testdata/bin/tmux"
	args := []string{"-V"}
	const optimistic = "optimistic"

	tests := []test{
		{OutputCmp{"stdout", command, ops.Ne, "1", args, optimistic}, false, true},
		{OutputCmp{"stdout", command, ops.Eq, "1", args, optimistic}, false, false},
		{OutputCmp{"stderr", command, ops.Like, "xxx", args, optimistic}, false, false},
		{OutputCmp{"stderr", command, ops.Unlike, "xxx", args, optimistic}, false, true},
		{OutputCmp{"combined", command, ops.Like, "xxx", args, optimistic}, false, false},
		{OutputCmp{"combined", command, ops.Unlike, "xxx", args, optimistic}, false, true},
		{OutputCmp{"stdout", command, ops.Ne, "1", args, "string"}, false, true},
		{OutputCmp{"stdout", command, ops.Ne, "1", args, "integer"}, true, false},
		{OutputCmp{"stdout", command, ops.Ne, "1", args, "version"}, true, false},
		{OutputCmp{"stdout", command, ops.Ne, "1", args, "float"}, true, false},
		{OutputCmp{"stdout", "bash -c", ops.Eq, "1", []string{"date|wc -l"}, "integer"}, false, true},
	}

	for _, test := range tests {
		ctx := types.Context{Debug: true}
		cmd := CLICmd{Output: test.Cmp}
		err := cmd.Run(&ctx)
		if test.Error {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		if test.Success {
			assert.True(t, ctx.Success)
		} else {
			assert.False(t, ctx.Success)
		}
	}
}
