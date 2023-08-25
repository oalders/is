package main

import (
	"runtime"
	"testing"

	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

func TestOSInfo(t *testing.T) {
	t.Parallel()
	tests := []string{"name", "version", "version-codename"}

	for _, v := range tests {
		ctx := types.Context{Debug: true}
		found, err := osInfo(&ctx, v)
		assert.NoError(t, err, v)
		assert.True(t, ctx.Success, v)
		assert.NotEmpty(t, found, v)
	}

	// id-like not present in Debian 11, so can't be in a blanket Linux test
	if runtime.GOOS == "linux" {
		tests := []string{"id", "pretty-name"}

		for _, v := range tests {
			ctx := types.Context{Debug: true}
			found, err := osInfo(&ctx, v)
			assert.NoError(t, err, v)
			assert.True(t, ctx.Success, v)
			assert.NotEmpty(t, found, v)
		}
	}
}

func TestOSCmd(t *testing.T) {
	t.Parallel()
	{
		ctx := types.Context{Debug: true}
		cmd := OSCmd{
			Attr: "name",
			Op:   "eq",
			Val:  "zzz",
		}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: true}
		cmd := OSCmd{
			Attr: "name",
			Op:   "ne",
			Val:  "zzz",
		}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: true}
		cmd := OSCmd{
			Attr: "version",
			Op:   "eq",
			Val:  "1",
		}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: true}
		cmd := OSCmd{
			Attr: "version",
			Op:   "ne",
			Val:  "1",
		}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: false}
		cmd := OSCmd{
			Attr: "name",
			Op:   "gte",
			Val:  "1",
		}
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: false}
		cmd := OSCmd{
			Attr: "name",
			Op:   "like",
			Val:  "zzzzz",
		}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: false}
		cmd := OSCmd{
			Attr: "name",
			Op:   "like",
			Val:  ".*",
		}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: false}
		cmd := OSCmd{
			Attr: "name",
			Op:   "unlike",
			Val:  "zzzzz",
		}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: false}
		cmd := OSCmd{
			Attr: "name",
			Op:   "unlike",
			Val:  ".*",
		}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: false}
		cmd := OSCmd{
			Attr: "name",
			Op:   "unlike",
			Val:  "/[/",
		}
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: false}
		cmd := OSCmd{
			Attr: "version",
			Op:   "like",
			Val:  "xxX",
		}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: false}
		cmd := OSCmd{
			Attr: "version",
			Op:   "like",
			Val:  ".*",
		}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: false}
		cmd := OSCmd{
			Attr: "version",
			Op:   "like",
			Val:  "[+",
		}
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: false}
		cmd := OSCmd{
			Attr: "version",
			Op:   "unlike",
			Val:  "xxX",
		}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: false}
		cmd := OSCmd{
			Attr: "version",
			Op:   "unlike",
			Val:  ".*",
		}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: false}
		cmd := OSCmd{
			Attr: "version",
			Op:   "unlike",
			Val:  "[+",
		}
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}
}
