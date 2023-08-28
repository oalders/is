package compare_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/oalders/is/compare"
	"github.com/oalders/is/types"
)

const (
	like   = "like"
	unlike = "unlike"
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
		{"3.3", "gt", "3.3", false},
		{"3.3", "ne", "3.3", false},
		{"3.3", "eq", "3.3", true},
		{"3.3", "gte", "3.3", true},
		{"3.3", "lte", "3.3", true},
		{"3.3", "lt", "3.3", false},
		{"3.3", "like", "3.3", true},
		{"3.3", "unlike", "4", true},

		{"3.3a", "gt", "3.3a", false},
		{"3.3a", "ne", "3.3a", false},
		{"3.3a", "eq", "3.3a", true},
		{"3.3a", "gte", "3.3a", true},
		{"3.3a", "lte", "3.3a", true},
		{"3.3a", "lt", "3.3a", false},
		{"3.3a", "like", "3.3a", true},
		{"3.3a", "unlike", "4", true},

		{"2", "gt", "1", true},
		{"2", "ne", "1", true},
		{"2", "eq", "1", false},
		{"2", "gte", "1", true},
		{"2", "lte", "1", false},
		{"2", "lt", "1", false},
		{"2", "like", "1", false},
		{"2", "unlike", "1", true},

		{"1", "gt", "2", false},
		{"1", "ne", "2", true},
		{"1", "eq", "2", false},
		{"1", "gte", "2", false},
		{"1", "lte", "2", true},
		{"1", "lt", "2", true},
		{"1", "like", "2", false},
		{"1", "unlike", "2", true},
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
		err := compare.Strings(ctx, like, "delboy trotter", "delboy")
		assert.True(t, ctx.Success)
		assert.NoError(t, err)
	}
	{
		err := compare.Strings(ctx, unlike, "delboy trotter", "delboy")
		assert.False(t, ctx.Success)
		assert.NoError(t, err)
	}
	{
		err := compare.Strings(ctx, like, "delboy trotter", "Zdelboy")
		assert.False(t, ctx.Success)
		assert.NoError(t, err)
	}
	{
		err := compare.Strings(ctx, unlike, "delboy trotter", "Zdelboy")
		assert.True(t, ctx.Success)
		assert.NoError(t, err)
	}
	{
		err := compare.Strings(ctx, like, "delboy trotter", "/[/")
		assert.False(t, ctx.Success)
		assert.Error(t, err)
	}
	{
		err := compare.Strings(ctx, unlike, "delboy trotter", "/[/")
		assert.False(t, ctx.Success)
		assert.Error(t, err)
	}
	ctx.Debug = true
}
