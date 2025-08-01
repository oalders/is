package main

import (
	"context"
	"testing"

	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

func TestSudoer(t *testing.T) {
	t.Parallel()
	ctx := types.Context{
		Context: context.Background(),
		Debug:   true,
	}
	cmd := UserCmd{}
	cmd.Sudoer = "sudoer"
	err := cmd.Run(&ctx)
	assert.NoError(t, err)
}
