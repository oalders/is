package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/oalders/is/ops"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

func TestArchCmd(t *testing.T) {
	t.Parallel()
	type ArchTest struct {
		Cmd     ArchCmd
		Error   bool
		Success bool
	}

	tests := []ArchTest{
		{ArchCmd{ops.Eq, "zzz"}, false, false},
		{ArchCmd{ops.Ne, "zzz"}, false, true},
		{ArchCmd{ops.In, "amd64,arm,arm64"}, false, true},
		{ArchCmd{ops.In, "X"}, false, false},
		{ArchCmd{ops.Like, "zzz"}, false, false},
		{ArchCmd{ops.Unlike, "zzz"}, false, true},
	}

	for _, test := range tests {
		ctx := types.Context{
			Context: context.Background(),
		}
		err := test.Cmd.Run(&ctx)
		name := fmt.Sprintf("%s %s", test.Cmd.Op, test.Cmd.Val)
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
