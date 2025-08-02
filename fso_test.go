package main

import (
	"context"
	"testing"

	"github.com/oalders/is/ops"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

func TestFSOLastModifiedTime(t *testing.T) {
	t.Parallel()
	const tmux = "testdata/bin/tmux"
	{
		ctx := &types.Context{
			Context: context.Background(),
			Debug:   true,
		}
		cmd := FSOCmd{Age: AgeCmp{tmux, ops.Gt, "1", "s"}}
		err := cmd.Run(ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := &types.Context{
			Context: context.Background(),
			Debug:   true,
		}
		cmd := FSOCmd{Age: AgeCmp{tmux, ops.Lt, "100000", "days"}}
		err := cmd.Run(ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := &types.Context{
			Context: context.Background(),
			Debug:   true,
		}
		cmd := FSOCmd{Age: AgeCmp{tmux, ops.Lt, "1.1", "d"}}
		err := cmd.Run(ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}
	{
		ctx := &types.Context{
			Context: context.Background(),
			Debug:   true,
		}
		cmd := FSOCmd{Age: AgeCmp{"tmuxxx", ops.Lt, "1", "d"}}
		err := cmd.Run(ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success)
	}
}
