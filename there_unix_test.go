//go:build unix

package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A slow CLI must neither fail nor starve the version lookups of the others.
func TestCLIVersionsSurviveSlowCLI(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	slow := filepath.Join(dir, "slowcli")
	fast := filepath.Join(dir, "fastcli")
	for path, script := range map[string]string{
		slow: "#!/bin/sh\n[ -z \"$1\" ] || sleep 30\n",
		fast: "#!/bin/sh\necho 'fastcli 1.2.3'\n",
	} {
		require.NoError(t, os.WriteFile(path, []byte(script), 0o700)) //nolint:gosec
	}
	// macOS can take seconds to first exec a new file; don't count that.
	// Without --version, slowcli exits at once.
	for _, path := range []string{slow, fast} {
		require.NoError(t, exec.CommandContext(context.Background(), path).Run())
	}

	timeout, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ctx := types.Context{Context: timeout}

	versions, err := cliVersions(&ctx, []string{slow, fast})

	require.NoError(t, err)
	assert.Equal(t, []string{"", "1.2.3"}, versions)
}
