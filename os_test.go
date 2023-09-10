package main

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/oalders/is/attr"
	"github.com/oalders/is/ops"
	"github.com/oalders/is/os"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

func TestOSInfo(t *testing.T) {
	t.Parallel()
	tests := []string{attr.Name, attr.Version, attr.VersionCodename}

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
		{OSCmd{attr.Name, ops.Eq, "zzz"}, false, false},
		{OSCmd{attr.Name, ops.Ne, "zzz"}, false, true},
		{OSCmd{attr.Version, ops.Eq, "1"}, false, false},
		{OSCmd{attr.Version, ops.Ne, "1"}, false, true},
		{OSCmd{attr.Version, ops.Eq, "[*&1.1.1.1.1"}, true, false},
		{OSCmd{attr.Name, ops.Like, "zzz"}, false, false},
		{OSCmd{attr.Name, ops.Like, ".*"}, false, true},
		{OSCmd{attr.Name, ops.Unlike, "zzz"}, false, true},
		{OSCmd{attr.Name, ops.Unlike, ".*"}, false, false},
		{OSCmd{attr.Name, ops.Unlike, "["}, true, false},
		{OSCmd{attr.Version, ops.Like, "xxx"}, false, false},
		{OSCmd{attr.Version, ops.Like, ".*"}, false, true},
		{OSCmd{attr.Version, ops.Like, "[+"}, true, false},
		{OSCmd{attr.Version, ops.Unlike, "xxX"}, false, true},
		{OSCmd{attr.Version, ops.Unlike, ".*"}, false, false},
		{OSCmd{attr.Version, ops.Unlike, "[+"}, true, false},
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
