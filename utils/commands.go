package utils

import (
	"bytes"
	"os/exec"
)

// Exec exec the command + args and returns stdout and stderr
func Exec(cmd string, args ...string) ([]byte, []byte, error) {
	c := exec.Command(cmd, args...)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	c.Stdout = stdout
	c.Stderr = stderr

	err := c.Run()
	if err != nil {
		return stdout.Bytes(), stderr.Bytes(), err
	}
	return stdout.Bytes(), stderr.Bytes(), nil
}
