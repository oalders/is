//nolint:testpackage // This test validates the unexported stream reader.
package parser

import (
	"context"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type synchronizedReader struct {
	output       string
	started      chan<- struct{}
	otherStarted <-chan struct{}
	hasStarted   bool
}

func (r *synchronizedReader) Read(buffer []byte) (int, error) {
	if !r.hasStarted {
		close(r.started)
		r.hasStarted = true
	}
	<-r.otherStarted

	if r.output == "" {
		return 0, io.EOF
	}

	n := copy(buffer, r.output)
	r.output = r.output[n:]
	return n, nil
}

func TestReadCLIOutputDrainsBothStreams(t *testing.T) {
	t.Parallel()

	stdoutStarted := make(chan struct{})
	stderrStarted := make(chan struct{})
	stdout := &synchronizedReader{
		output:       "noisy-version 1.2.3",
		started:      stdoutStarted,
		otherStarted: stderrStarted,
	}
	stderr := &synchronizedReader{
		started:      stderrStarted,
		otherStarted: stdoutStarted,
	}

	type result struct {
		output string
		err    error
	}
	resultCh := make(chan result, 1)
	go func() {
		output, err := readCLIOutput(
			&types.Context{Context: context.Background()},
			[]string{"noisy-version", "--version"},
			stdout,
			stderr,
		)
		resultCh <- result{output: output, err: err}
	}()

	select {
	case result := <-resultCh:
		require.NoError(t, result.err)
		assert.Equal(t, "noisy-version 1.2.3", result.output)
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for both streams to be read")
	}
}

func TestReadCLIOutputLimitsEachStream(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		stdout string
		stderr string
	}{
		{
			name:   "stdout",
			stdout: strings.Repeat("s", maxVersionBytes+1),
			stderr: "stderr",
		},
		{
			name:   "stderr",
			stderr: strings.Repeat("e", maxVersionBytes+1),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			output, err := readCLIOutput(
				&types.Context{Context: context.Background()},
				[]string{"noisy-version", "--version"},
				strings.NewReader(test.stdout),
				strings.NewReader(test.stderr),
			)

			require.NoError(t, err)
			expected := test.stdout
			if expected == "" {
				expected = test.stderr
			}
			assert.Equal(t, expected[:maxVersionBytes], output)
		})
	}
}
