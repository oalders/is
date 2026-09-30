//nolint:testpackage // This test validates the unexported stream reader.
package parser

import (
	"context"
	"io"
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
