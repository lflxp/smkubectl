package utils

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"
)

func ExecCmd(in string) error {
	cmd := exec.Command("sh", "-c", in)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		slog.Error(err.Error(), "CMD", in)
		return err
	}
	return nil
}

func ExecCmdString(in string) (string, error) {
	cmd := exec.Command("sh", "-c", in)
	cmd.Stdin = os.Stdin
	var out bytes.Buffer
	cmd.Stdout = &out
	var errout bytes.Buffer
	cmd.Stderr = &errout
	err := cmd.Run()
	if err != nil {
		slog.Error(err.Error(), "CMD", in, "errout", errout.String())
		return "", err
	}
	return out.String(), nil
}
