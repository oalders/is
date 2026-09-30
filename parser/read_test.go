//go:build unix

//nolint:testpackage // This test validates unexported output handling.
package parser

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// writeCLI writes an executable shell script named "orphaner" and returns
// its path.
func writeCLI(t *testing.T, body string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "orphaner")
	require.NoError(t, os.WriteFile(path, []byte("#!/bin/sh\n"+body), 0o700)) //nolint:gosec
	return path
}

type cliResult struct {
	output string
	err    error
}

func runCLIOutput(ctx context.Context, t *testing.T, cli string) cliResult {
	t.Helper()
	resultCh := make(chan cliResult, 1)
	go func() {
		output, err := cliOutput(&types.Context{Context: ctx}, cli)
		resultCh <- cliResult{output: output, err: err}
	}()

	select {
	case result := <-resultCh:
		return result
	case <-time.After(5 * time.Second):
		t.Fatal("timed out: an orphaned grandchild kept the output pipe open")
		return cliResult{}
	}
}

// A CLI that exits after leaving a background process attached to its
// stdout must not block us until that background process exits.
func TestCLIOutputIgnoresOrphanedGrandchild(t *testing.T) {
	t.Parallel()

	cli := writeCLI(t, "sleep 10 &\necho 'orphaner 1.2.3'\n")
	result := runCLIOutput(context.Background(), t, cli)

	require.NoError(t, result.err)
	assert.Equal(t, "orphaner 1.2.3\n", result.output)
}

// When cancellation kills a CLI whose child still holds the pipe (e.g. a
// PyInstaller bootloader and its worker), we must return rather than hang,
// and must not leave that child running.
func TestCLIOutputKillsGrandchildOnCancel(t *testing.T) {
	t.Parallel()

	pidFile := filepath.Join(t.TempDir(), "pid")
	cli := writeCLI(t, "sleep 10 &\necho $! > "+pidFile+"\nsleep 10\n")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var pid int
	go func() {
		for {
			contents, err := os.ReadFile(pidFile)
			if err == nil && strings.HasSuffix(string(contents), "\n") {
				pid, _ = strconv.Atoi(strings.TrimSpace(string(contents)))
				cancel()
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
	result := runCLIOutput(ctx, t, cli)

	require.ErrorIs(t, result.err, context.Canceled)
	require.NotZero(t, pid)
	assert.Eventually(t, func() bool {
		return errors.Is(syscall.Kill(pid, 0), syscall.ESRCH)
	}, 2*time.Second, 50*time.Millisecond, "orphaned grandchild %d is still running", pid)
}

func TestCLIOutputFallsBackToStderr(t *testing.T) {
	t.Parallel()

	cli := writeCLI(t, "echo 'orphaner 1.2.3' >&2\n")
	result := runCLIOutput(context.Background(), t, cli)

	require.NoError(t, result.err)
	assert.Equal(t, "orphaner 1.2.3\n", result.output)
}

func TestCappedBufferLimitsOutput(t *testing.T) {
	t.Parallel()

	var buffer cappedBuffer
	chunk := strings.Repeat("s", maxVersionBytes-1)
	for range 3 {
		n, err := buffer.Write([]byte(chunk))
		require.NoError(t, err)
		assert.Equal(t, len(chunk), n, "writes past the cap are discarded, not rejected")
	}
	assert.Equal(t, strings.Repeat("s", maxVersionBytes), buffer.String())
}
