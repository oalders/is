package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/oalders/is/attr"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

//nolint:unparam
func prependPath(path string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(wd, path) + string(os.PathListSeparator) + os.Getenv("PATH")
}

//nolint:paralleltest,nolintlint
func TestKnownCmd(t *testing.T) {
	t.Setenv("PATH", prependPath("testdata/bin"))

	const command = "semver"
	type testableOS struct {
		Attr    string
		Error   bool
		Success bool
	}

	osTests := []testableOS{
		{attr.Name, false, true},
		{attr.Version, false, true},
		{"tmuxxx", false, false},
		{"tmuxxx", false, false},
	}

	if runtime.GOOS == "darwin" {
		osTests = append(osTests, testableOS{attr.Version, false, true})
	}

	for _, test := range osTests {
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.OS.Attr = test.Attr
		err := cmd.Run(&ctx)
		name := fmt.Sprintf("%s err: %t success: %t", test.Attr, test.Error, test.Success)
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

	type testableCLI struct {
		Cmd     KnownCmd
		Error   bool
		Success bool
	}
	cliTests := []testableCLI{
		{KnownCmd{CLI: KnownCLI{attr.Version, "gitzzz"}}, false, false},
		{KnownCmd{CLI: KnownCLI{attr.Version, command}}, false, true},
		{KnownCmd{CLI: KnownCLI{attr.Version, command}, Major: true}, false, true},
		{KnownCmd{CLI: KnownCLI{attr.Version, command}, Minor: true}, false, true},
		{KnownCmd{CLI: KnownCLI{attr.Version, command}, Patch: true}, false, true},
	}

	for _, test := range cliTests {
		ctx := types.Context{Debug: true}
		err := test.Cmd.Run(&ctx)

		switch test.Error {
		case true:
			assert.Error(t, err)
		default:
			assert.NoError(t, err)
		}

		switch test.Success {
		case true:
			assert.True(t, ctx.Success)
		default:
			assert.False(t, ctx.Success)
		}
	}

	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{}
		cmd.Arch.Attr = "arch"
		err := cmd.Run(&ctx)
		assert.NoError(t, err)
		assert.True(t, ctx.Success, "success")
	}
	{
		ctx := types.Context{Debug: true}
		cmd := KnownCmd{Major: true}
		cmd.OS.Attr = "name"
		err := cmd.Run(&ctx)
		assert.Error(t, err)
		assert.False(t, ctx.Success, "success")
	}
}
