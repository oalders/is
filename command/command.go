package command

import (
	"errors"
	"io"
	"os/exec"
	"strings"
)

func Cmd(command string, args ...string) *exec.Cmd {
	return exec.Command(command, args...)
}

func Output(cmd *exec.Cmd, stream string) (string, error) {
	var pipe io.ReadCloser
	var err error
	var output []byte

	switch stream {
	case "stdout":
		pipe, err = cmd.StdoutPipe()
		if err != nil {
			return "", errors.Join(errors.New("stdout pipe"), err)
		}
	case "stderr":
		pipe, err = cmd.StderrPipe()
		if err != nil {
			return "", errors.Join(errors.New("stderr pipe"), err)
		}
	case "combined":
		output, err = cmd.CombinedOutput()
		if err != nil {
			return "", errors.Join(errors.New("combined output"), err)
		}
	}

	// This means it's not combined output
	if len(output) == 0 {
		if err := cmd.Start(); err != nil {
			return "", errors.Join(errors.New("starting command"), err)
		}

		output, err = io.ReadAll(pipe)
		if err != nil {
			return "", errors.Join(errors.New("read output"), err)
		}
	}
	return strings.TrimSpace(string(output)), nil
}
