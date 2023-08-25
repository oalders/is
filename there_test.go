package main

import (
	"testing"

	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

func TestThereCmd(t *testing.T) {
	{
		ctx := types.Context{Debug: true}
		cmd := ThereCmd{Name: "cat"}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success)
	}
	{
		ctx := types.Context{Debug: true}
		cmd := ThereCmd{Name: "catzzzzz"}
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.False(t, ctx.Success)
	}
}
