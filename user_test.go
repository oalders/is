package main

import (
	"testing"

	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

func TestSudoer(t *testing.T) {
	ctx := types.Context{Debug: true}
	cmd := UserCmd{}
	cmd.Sudoer = "sudoer"
	err := cmd.Run(&ctx)
	assert.NoError(t, err)
}
