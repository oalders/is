package main

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/oalders/is/os"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

func TestOSInfo(t *testing.T) {
	t.Parallel()
	tests := []string{"name", "version", "version-codename"}

	for _, v := range tests {
		ctx := types.Context{Debug: true}
		found, err := os.Info(&ctx, v)
		assert.NoError(t, err, v)
		assert.True(t, ctx.Success, v)
		assert.NotEmpty(t, found, v)
	}

	// id-like not present in Debian 11, so can't be in a blanket Linux test
	if runtime.GOOS == "linux" {
		tests := []string{"id", "pretty-name"}

		for _, v := range tests {
			ctx := types.Context{Debug: true}
			found, err := os.Info(&ctx, v)
			assert.NoError(t, err, v)
			assert.True(t, ctx.Success, v)
			assert.NotEmpty(t, found, v)
		}
	}
}

func TestOSCmd(t *testing.T) {
	t.Parallel()
	type OSTest struct {
		Cmd     OSCmd
		Error   bool
		Success bool
	}

	tests := []OSTest{
		{
			Cmd:     OSCmd{"name", "eq", "zzz"},
			Error:   false,
			Success: false,
		},
		{
			Cmd:     OSCmd{"name", "ne", "zzz"},
			Error:   false,
			Success: true,
		},
		{
			Cmd:     OSCmd{"version", "eq", "1"},
			Error:   false,
			Success: false,
		},
		{
			Cmd:     OSCmd{"version", "ne", "1"},
			Error:   false,
			Success: true,
		},
		{
			Cmd:     OSCmd{"name", "like", "zzzzz"},
			Error:   false,
			Success: false,
		},
		{
			Cmd:     OSCmd{"name", "like", ".*"},
			Error:   false,
			Success: true,
		},
		{
			Cmd:     OSCmd{"name", "unlike", "zzzzz"},
			Error:   false,
			Success: true,
		},
		{
			Cmd:     OSCmd{"name", "unlike", ".*"},
			Error:   false,
			Success: false,
		},
		{
			Cmd:     OSCmd{"name", "unlike", "["},
			Error:   true,
			Success: false,
		},
		{
			Cmd:     OSCmd{"version", "like", "xxx"},
			Error:   false,
			Success: false,
		},
		{
			Cmd:     OSCmd{"version", "like", ".*"},
			Error:   false,
			Success: true,
		},
		{
			Cmd:     OSCmd{"version", "like", "[+"},
			Error:   true,
			Success: false,
		},
		{
			Cmd:     OSCmd{"version", "unlike", "xxX"},
			Error:   false,
			Success: true,
		},
		{
			Cmd:     OSCmd{"version", "unlike", ".*"},
			Error:   false,
			Success: false,
		},
		{
			Cmd:     OSCmd{"version", "unlike", "[+"},
			Error:   true,
			Success: false,
		},
	}

	for _, test := range tests {
		ctx := types.Context{Debug: false}
		err := test.Cmd.Run(&ctx)
		name := fmt.Sprintf("%s %s %s", test.Cmd.Attr, test.Cmd.Op, test.Cmd.Val)
		if test.Error {
			assert.Error(t, err, name)
		} else {
			assert.NoError(t, err, name)
		}
		if test.Success {
			assert.True(t, ctx.Success, name)
		} else {
			assert.False(t, ctx.Success, name)
		}
	}
}
