package main

import (
	"context"
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
	tests := []string{attr.Name, attr.Version}

	if runtime.GOOS == "linux" {
		tests = append(tests, "id", "pretty-name")
		ctx := types.Context{
			Context: context.Background(),
			Debug:   true,
		}
		found, err := os.Info(&ctx, "name")
		assert.NoError(t, err)

		// id-like not present in Debian 11, so can't be in a blanket Linux
		// test
		if found == "ubuntu" {
			tests = append(tests, "id-like")
		}
	}

	for _, v := range tests {
		ctx := types.Context{
			Context: context.Background(),
			Debug:   true,
		}
		found, err := os.Info(&ctx, v)
		assert.NoError(t, err, v)
		assert.NotEmpty(t, found, v)
	}
}

func TestOSCmd(t *testing.T) {
	t.Parallel()
	type OSTest struct {
		Cmd     OSCmd
		Error   bool
		Success bool
	}

	const major = false
	const minor = false
	const patch = false

	tests := []OSTest{
		{OSCmd{attr.Name, ops.Eq, "zzz", major, minor, patch}, false, false},
		{OSCmd{attr.Name, ops.Ne, "zzz", major, minor, patch}, false, true},
		{OSCmd{attr.Version, ops.Eq, "1", major, minor, patch}, false, false},
		{OSCmd{attr.Version, ops.Ne, "1", major, minor, patch}, false, true},
		{OSCmd{attr.Version, ops.Eq, "[*&1.1.1.1.1", major, minor, patch}, true, false},
		{OSCmd{attr.Name, ops.Like, "zzz", major, minor, patch}, false, false},
		{OSCmd{attr.Name, ops.Like, ".*", major, minor, patch}, false, true},
		{OSCmd{attr.Name, ops.Unlike, "zzz", major, minor, patch}, false, true},
		{OSCmd{attr.Name, ops.Unlike, ".*", major, minor, patch}, false, false},
		{OSCmd{attr.Name, ops.Unlike, "[", major, minor, patch}, true, false},
		{OSCmd{attr.Version, ops.Like, "xxx", major, minor, patch}, false, false},
		{OSCmd{attr.Version, ops.Like, ".*", major, minor, patch}, false, true},
		{OSCmd{attr.Version, ops.Like, "[+", major, minor, patch}, true, false},
		{OSCmd{attr.Version, ops.Unlike, "xxX", major, minor, patch}, false, true},
		{OSCmd{attr.Version, ops.Unlike, ".*", major, minor, patch}, false, false},
		{OSCmd{attr.Version, ops.Unlike, "[+", major, minor, patch}, true, false},
		{OSCmd{attr.Version, ops.Gt, "0", true, minor, patch}, false, true},
		{OSCmd{attr.Version, ops.Gte, "0", major, true, patch}, false, true},
		{OSCmd{attr.Version, ops.Gte, "0", major, minor, true}, false, true},
		{OSCmd{attr.Name, ops.Gt, "0", true, minor, patch}, true, false},
		{OSCmd{attr.Name, ops.Gt, "0", major, true, patch}, true, false},
		{OSCmd{attr.Name, ops.Gt, "0", major, minor, true}, true, false},
	}

	for _, test := range tests {
		ctx := types.Context{
			Context: context.Background(),
			Debug:   true,
		}
		err := test.Cmd.Run(&ctx)
		name := fmt.Sprintf(
			"%s %s %s major: %t minor: %t patch: %t",
			test.Cmd.Attr,
			test.Cmd.Op,
			test.Cmd.Val,
			test.Cmd.Major,
			test.Cmd.Minor,
			test.Cmd.Patch,
		)
		if test.Error {
			assert.Error(t, err, "has error "+name)
		} else {
			assert.NoError(t, err, "has no error "+name)
		}
		if test.Success {
			assert.True(t, ctx.Success, "has success "+name)
		} else {
			assert.False(t, ctx.Success, "has no success "+name)
		}
	}
}
