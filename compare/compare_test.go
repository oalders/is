package compare_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/oalders/is/compare"
	"github.com/oalders/is/ops"
	"github.com/oalders/is/types"
)

func TestCompareCLIVersions(t *testing.T) {
	t.Parallel()
	type test struct {
		Got     string
		Op      string
		Want    string
		Success bool
	}
	tests := []test{
		{"3.3", ops.Gt, "3.3", false},
		{"3.3", ops.Ne, "3.3", false},
		{"3.3", ops.Eq, "3.3", true},
		{"3.3", ops.Gte, "3.3", true},
		{"3.3", ops.Lte, "3.3", true},
		{"3.3", ops.Lt, "3.3", false},
		{"3.3", ops.Like, "3.3", true},
		{"3.3", ops.Unlike, "4", true},

		{"3.3a", ops.Gt, "3.3a", false},
		{"3.3a", ops.Ne, "3.3a", false},
		{"3.3a", ops.Eq, "3.3a", true},
		{"3.3a", ops.Gte, "3.3a", true},
		{"3.3a", ops.Lte, "3.3a", true},
		{"3.3a", ops.Lt, "3.3a", false},
		{"3.3a", ops.Like, "3.3a", true},
		{"3.3a", ops.Unlike, "4", true},

		{"2", ops.Gt, "1", true},
		{"2", ops.Ne, "1", true},
		{"2", ops.Eq, "1", false},
		{"2", ops.Gte, "1", true},
		{"2", ops.Lte, "1", false},
		{"2", ops.Lt, "1", false},
		{"2", ops.Like, "1", false},
		{"2", ops.Unlike, "1", true},

		{"1", ops.Gt, "2", false},
		{"1", ops.Ne, "2", true},
		{"1", ops.Eq, "2", false},
		{"1", ops.Gte, "2", false},
		{"1", ops.Lte, "2", true},
		{"1", ops.Lt, "2", true},
		{"1", ops.Like, "2", false},
		{"1", ops.Unlike, "2", true},
	}
	for _, v := range tests {
		ctx := &types.Context{Debug: false}
		err := compare.CLIVersions(ctx, v.Op, v.Got, v.Want)
		assert.NoError(t, err)
		if v.Success {
			assert.True(t, ctx.Success, "success")
		} else {
			assert.False(t, ctx.Success, "failure")
		}
	}
}

func TestStrings(t *testing.T) {
	t.Parallel()
	ctx := &types.Context{}
	{
		err := compare.Strings(ctx, ops.Like, "delboy trotter", "delboy")
		assert.True(t, ctx.Success)
		assert.NoError(t, err)
	}
	{
		err := compare.Strings(ctx, ops.Unlike, "delboy trotter", "delboy")
		assert.False(t, ctx.Success)
		assert.NoError(t, err)
	}
	{
		err := compare.Strings(ctx, ops.Like, "delboy trotter", "Zdelboy")
		assert.False(t, ctx.Success)
		assert.NoError(t, err)
	}
	{
		err := compare.Strings(ctx, ops.Unlike, "delboy trotter", "Zdelboy")
		assert.True(t, ctx.Success)
		assert.NoError(t, err)
	}
	{
		err := compare.Strings(ctx, ops.Like, "delboy trotter", "/[/")
		assert.False(t, ctx.Success)
		assert.Error(t, err)
	}
	{
		err := compare.Strings(ctx, ops.Unlike, "delboy trotter", "/[/")
		assert.False(t, ctx.Success)
		assert.Error(t, err)
	}
	ctx.Debug = true
}
