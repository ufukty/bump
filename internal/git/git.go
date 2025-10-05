package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/ufukty/bump/internal/labels"
)

func describe() (string, error) {
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
func register(v string) error {
	b := &bytes.Buffer{}
	cmd := exec.Command("git", "tag", v)
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

// CAUTION: side effects
func IncrementAndApply(label string) error {
	v, err := describe()
	if err != nil {
		return fmt.Errorf("git describe: %w", err)
	}

	v2, err := labels.Increment(v, label)
	if err != nil {
		return fmt.Errorf("incrementing: %w", err)
	}

	err = register(v2)
	if err != nil {
		return fmt.Errorf("registering the next version: %w", err)
	}

	return nil
}
