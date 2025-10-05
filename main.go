package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ufukty/bump/internal/git"
	"github.com/ufukty/bump/internal/labels"
)

func Main() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("expected to see one argument among: %s", strings.Join(labels.Mods, ", "))
	}

	v1, err := git.Describe()
	if err != nil {
		return fmt.Errorf("git describe: %w", err)
	}

	l1, err := labels.Parse(v1)
	if err != nil {
		return fmt.Errorf("parsing current version: %w", err)
	}

	l2, err := labels.Increment(l1, os.Args[1])
	if err != nil {
		return fmt.Errorf("incrementing: %w", err)
	}

	if err := git.Tag(l2.String()); err != nil {
		return fmt.Errorf("git tag: %w", err)
	}

	return nil
}

func main() {
	if err := Main(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
