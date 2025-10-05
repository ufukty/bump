package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Describe() (string, error) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := exec.Command("git", "describe", "--tag")
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	if strings.Contains(stderr.String(), "cannot describe anything") {
		return "v0.0.0", nil
	}
	if err != nil {
		return "", fmt.Errorf("%s: %w", cmd.String(), err)
	}
	if stderr.Len() > 0 {
		return "", fmt.Errorf("checking output:\n%s", stderr.String())
	}
	return stdout.String(), nil
}

// CAUTION: side effects
func Register(verstr string) error {
	b := &bytes.Buffer{}
	cmd := exec.Command("git", "tag", verstr)
	cmd.Stdout = b
	cmd.Stderr = b
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%s: %w", cmd.String(), err)
	}
	if b.Len() > 0 {
		return fmt.Errorf("checking output:\n%s", b.String())
	}
	return nil
}
