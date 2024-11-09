package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func describe() (string, error) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := exec.Command("git", "describe", "--tag")
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%s: %w", cmd.String(), err)
	}
	if stderr.Len() > 0 {
		return "", fmt.Errorf("checking output:\n%s", stderr.String())
	}
	return stdout.String(), nil
}

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

func Main() error {
	mods := []string{"major", "minor", "fix"}

	if len(os.Args) != 2 {
		return fmt.Errorf("expected to see one argument among: %s", strings.Join(mods, ", "))
	}

	v, err := describe()
	if err != nil {
		return fmt.Errorf("git describe: %w", err)
	}

	r, err := regexp.Compile(`v([0-9]+)\.([0-9]+)\.([0-9]+).*`)
	if err != nil {
		return fmt.Errorf("compiling regex: %w", err)
	}
	ms := r.FindStringSubmatch(v)
	if len(ms) != 4 {
		return fmt.Errorf("expected to see 'major.minor.fix' format: %s", v)
	}
	ms = ms[1:]

	i := slices.Index(mods, os.Args[1])
	if i == -1 {
		return fmt.Errorf("invalid argument. available arguments: %s", strings.Join(mods, ", "))
	}

	n, err := strconv.Atoi(ms[i])
	if err != nil {
		return fmt.Errorf("parsing integer: %w", err)
	}
	ms[i] = fmt.Sprintf("%d", n+1)
	for j := i + 1; j < 3; j++ {
		ms[j] = "0"
	}

	err = register("v" + strings.Join(ms, "."))
	if err != nil {
		return fmt.Errorf("registering the next version: %w", err)
	}

	return nil
}

func main() {
	if err := Main(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
