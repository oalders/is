package command

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func Output(cmd *exec.Cmd, stream string) (string, error) {
	var pipe io.ReadCloser
	var err error
	var output []byte

	switch stream {
	case "stdout":
		pipe, err = cmd.StdoutPipe()
		if err != nil {
			return "", fmt.Errorf("stdout pipe: %w", err)
		}
		defer pipe.Close()
	case "stderr":
		pipe, err = cmd.StderrPipe()
		if err != nil {
			return "", fmt.Errorf("stderr pipe: %w", err)
		}
		defer pipe.Close()
	case "combined":
		output, err = cmd.CombinedOutput()
		if err != nil {
			return "", fmt.Errorf("combined output: %w", err)
		}
	}

	// This means it's not combined output
	if len(output) == 0 {
		if err := cmd.Start(); err != nil {
			return "", fmt.Errorf("starting command: %w", err)
		}

		output, err = io.ReadAll(pipe)
		if err != nil {
			return "", fmt.Errorf("read output: %w", err)
		}
	}
	return strings.TrimSpace(string(output)), nil
}
