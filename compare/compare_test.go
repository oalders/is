package compare

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/hashicorp/go-version"
	"github.com/oalders/is/types"
)

func TestCompareCLIVersions(t *testing.T) {
	t.Parallel()
	{
		want, _ := version.NewVersion("3.3")
		got, _ := version.NewVersion("3.3")
		assert.False(t, CLIVersions("gt", got, want))
		assert.False(t, CLIVersions("ne", got, want))
		assert.True(t, CLIVersions("eq", got, want))
		assert.True(t, CLIVersions("gte", got, want))
		assert.False(t, CLIVersions("lt", got, want))
		assert.True(t, CLIVersions("lte", got, want))
	}

	{
		want, _ := version.NewVersion("3.3a")
		got, _ := version.NewVersion("3.3a")
		assert.False(t, CLIVersions("gt", got, want))
		assert.False(t, CLIVersions("ne", got, want))
		assert.True(t, CLIVersions("eq", got, want))
		assert.True(t, CLIVersions("gte", got, want))
		assert.True(t, CLIVersions("lte", got, want))
		assert.False(t, CLIVersions("lt", got, want))
	}

	{
		want, _ := version.NewVersion("2")
		got, _ := version.NewVersion("1")
		assert.False(t, CLIVersions("gt", got, want))
		assert.True(t, CLIVersions("ne", got, want))
		assert.False(t, CLIVersions("eq", got, want))
		assert.False(t, CLIVersions("gte", got, want))
		assert.True(t, CLIVersions("lte", got, want))
		assert.True(t, CLIVersions("lt", got, want))
	}

	{
		want, _ := version.NewVersion("1")
		got, _ := version.NewVersion("2")
		assert.True(t, CLIVersions("gt", got, want))
		assert.True(t, CLIVersions("ne", got, want))
		assert.False(t, CLIVersions("eq", got, want))
		assert.True(t, CLIVersions("gte", got, want))
		assert.False(t, CLIVersions("lte", got, want))
		assert.False(t, CLIVersions("lt", got, want))
	}
}

func TestStrings(t *testing.T) {
	t.Parallel()
	ctx := &types.Context{}
	{
		err := Strings(ctx, "like", "delboy trotter", "delboy")
		assert.True(t, ctx.Success)
		assert.NoError(t, err)
	}
	{
		err := Strings(ctx, "unlike", "delboy trotter", "delboy")
		assert.False(t, ctx.Success)
		assert.NoError(t, err)
	}
	{
		err := Strings(ctx, "like", "delboy trotter", "Zdelboy")
		assert.False(t, ctx.Success)
		assert.NoError(t, err)
	}
	{
		err := Strings(ctx, "unlike", "delboy trotter", "Zdelboy")
		assert.True(t, ctx.Success)
		assert.NoError(t, err)
	}
	{
		err := Strings(ctx, "like", "delboy trotter", "/[/")
		assert.False(t, ctx.Success)
		assert.Error(t, err)
	}
	{
		err := Strings(ctx, "unlike", "delboy trotter", "/[/")
		assert.False(t, ctx.Success)
		assert.Error(t, err)
	}
	{
		err := Strings(ctx, "Zunlike", "x", "x")
		assert.Error(t, err)
	}
}
