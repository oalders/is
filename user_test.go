package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSudoer(t *testing.T) {
	ctx := Context{Debug: true}
	cmd := UserCmd{}
	cmd.Sudoer = "sudoer"
	err := cmd.Run(&ctx)
	assert.NoError(t, err)
}
