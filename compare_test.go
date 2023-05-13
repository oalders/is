package main

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/hashicorp/go-version"
)

func TestCompareCLIVersions(t *testing.T) {
	{
		want, _ := version.NewVersion("3.3")
		got, _ := version.NewVersion("3.3")
		assert.False(t, compareCLIVersions("gt", got, want))
		assert.False(t, compareCLIVersions("ne", got, want))
		assert.True(t, compareCLIVersions("eq", got, want))
		assert.True(t, compareCLIVersions("gte", got, want))
		assert.False(t, compareCLIVersions("lt", got, want))
		assert.True(t, compareCLIVersions("lte", got, want))
	}

	{
		want, _ := version.NewVersion("3.3a")
		got, _ := version.NewVersion("3.3a")
		assert.False(t, compareCLIVersions("gt", got, want))
		assert.False(t, compareCLIVersions("ne", got, want))
		assert.True(t, compareCLIVersions("eq", got, want))
		assert.True(t, compareCLIVersions("gte", got, want))
		assert.True(t, compareCLIVersions("lte", got, want))
		assert.False(t, compareCLIVersions("lt", got, want))
	}

	{
		want, _ := version.NewVersion("2")
		got, _ := version.NewVersion("1")
		assert.False(t, compareCLIVersions("gt", got, want))
		assert.True(t, compareCLIVersions("ne", got, want))
		assert.False(t, compareCLIVersions("eq", got, want))
		assert.False(t, compareCLIVersions("gte", got, want))
		assert.True(t, compareCLIVersions("lte", got, want))
		assert.True(t, compareCLIVersions("lt", got, want))
	}

	{
		want, _ := version.NewVersion("1")
		got, _ := version.NewVersion("2")
		assert.True(t, compareCLIVersions("gt", got, want))
		assert.True(t, compareCLIVersions("ne", got, want))
		assert.False(t, compareCLIVersions("eq", got, want))
		assert.True(t, compareCLIVersions("gte", got, want))
		assert.False(t, compareCLIVersions("lte", got, want))
		assert.False(t, compareCLIVersions("lt", got, want))
	}
}
