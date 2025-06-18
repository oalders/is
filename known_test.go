package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/oalders/is/attr"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		{KnownCmd{CLI: KnownCLI{
			Attr: attr.Version,
			Name: "gitzzz",
		}}, false, false},
		{KnownCmd{
			CLI: KnownCLI{
				Attr: attr.Version,
				Name: command,
			},
		}, false, true},
		{KnownCmd{
			CLI: KnownCLI{
				Attr:    attr.Version,
				Name:    command,
				Version: Version{Major: true},
			},
		}, false, true},
		{KnownCmd{CLI: KnownCLI{Attr: attr.Version, Name: command, Version: Version{
			Minor: true,
		}}}, false, true},
		{KnownCmd{CLI: KnownCLI{Attr: attr.Version, Name: command, Version: Version{
			Patch: true,
		}}}, false, true},
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
}

func Test_getEnv(t *testing.T) {
	t.Run("regular environment variable", func(t *testing.T) {
		ctx := types.Context{}
		// Setup
		testVarName := "TEST_ENV_VAR"
		testValue := "test_value"
		t.Setenv(testVarName, testValue)

		// Test non-JSON retrieval
		value, err := getEnv(&ctx, testVarName, false)
		require.True(t, ctx.Success)
		require.NoError(t, err)
		assert.Equal(t, testValue, value)
	})

	t.Run("path environment variable as JSON", func(t *testing.T) {
		ctx := types.Context{}
		// Setup
		pathValue := "/usr/bin:/usr/local/bin:/bin"
		t.Setenv(path, pathValue)

		// Test JSON retrieval
		value, err := getEnv(&ctx, path, true)
		require.NoError(t, err)
		assert.Contains(t, value, "/usr/bin")
		assert.Contains(t, value, "/usr/local/bin")
		assert.Contains(t, value, "/bin")

		// Verify it's valid JSON
		assert.True(t, strings.HasPrefix(value, "["))
		assert.True(t, strings.HasSuffix(value, "]"))
	})

	t.Run("manpath environment variable as JSON", func(t *testing.T) {
		ctx := types.Context{}
		// Setup
		manpathValue := "/usr/share/man:/usr/local/share/man"
		t.Setenv(manpath, manpathValue)

		// Test JSON retrieval
		value, err := getEnv(&ctx, manpath, true)
		require.NoError(t, err)
		assert.Contains(t, value, "/usr/share/man")
		assert.Contains(t, value, "/usr/local/share/man")

		// Verify it's valid JSON
		assert.True(t, strings.HasPrefix(value, "["))
		assert.True(t, strings.HasSuffix(value, "]"))
	})

	t.Run("non-path variable as JSON returns empty array", func(t *testing.T) {
		ctx := types.Context{}
		// Setup
		testVarName := "REGULAR_VAR"
		testValue := "something"
		t.Setenv(testVarName, testValue)

		// Test JSON retrieval for non-path/manpath variable
		value, err := getEnv(&ctx, testVarName, true)
		require.NoError(t, err)
		assert.Equal(t, "[\n    \"something\"\n]", strings.TrimSpace(value))
	})

	t.Run("non-existent variable", func(t *testing.T) { //nolint:paralleltest,nolintlint
		ctx := types.Context{}
		// Test non-existent variable with non-JSON mode
		value, err := getEnv(&ctx, "NON_EXISTENT_VAR", false)
		require.NoError(t, err)
		assert.Equal(t, "", value)

		// Test non-existent variable with JSON mode
		value, err = getEnv(&ctx, "NON_EXISTENT_VAR", true)
		require.NoError(t, err)
		assert.Equal(t, "null", strings.TrimSpace(value))
		require.False(t, ctx.Success)
	})
}

func Test_envSummary(t *testing.T) {
	t.Run("tabular output", func(t *testing.T) {
		// Set up test environment
		t.Setenv("TEST_VAR", "test_value")
		t.Setenv("PATH", "/usr/bin:/usr/local/bin")

		// Create context and capture stdout
		ctx := &types.Context{}
		originalStdout := os.Stdout
		r, w, err := os.Pipe() //nolint:varnamelen
		require.NoError(t, err)
		os.Stdout = w

		// Test the function
		err = envSummary(ctx, false)
		require.NoError(t, err)

		// Restore stdout
		w.Close()
		os.Stdout = originalStdout

		// Read output
		var output strings.Builder
		_, err = io.Copy(&output, r)
		require.NoError(t, err)

		// Basic validations
		assert.True(t, ctx.Success)
		assert.Contains(t, output.String(), "TEST_VAR")
		assert.Contains(t, output.String(), "test_value")
		assert.Contains(t, output.String(), "PATH")
	})

	t.Run("JSON output", func(t *testing.T) {
		// Set up test environment
		t.Setenv("TEST_VAR", "test_value")
		t.Setenv("PATH", "/usr/bin:/usr/local/bin")

		// Create context and capture stdout
		ctx := &types.Context{}
		originalStdout := os.Stdout
		r, w, err := os.Pipe() //nolint:varnamelen
		require.NoError(t, err)
		os.Stdout = w

		// Test the function
		err = envSummary(ctx, true)
		require.NoError(t, err)

		// Restore stdout
		w.Close()
		os.Stdout = originalStdout

		// Read output
		var output strings.Builder
		_, err = io.Copy(&output, r)
		require.NoError(t, err)

		// Basic validations
		assert.True(t, ctx.Success)
		assert.Contains(t, output.String(), "TEST_VAR")
		assert.Contains(t, output.String(), "test_value")
		assert.Contains(t, output.String(), "PATH")
		assert.Contains(t, output.String(), "/usr/bin")
	})
}
