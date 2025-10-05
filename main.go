package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ufukty/bump/internal/git"
)

func Main() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("expected to see one argument among: %s", strings.Join(git.Mods, ", "))
	}

	label := os.Args[1]
	if err := git.IncrementAndApply(label); err != nil {
		return fmt.Errorf("incrementing %s: %w", label, err)
	}

	return nil
}

func main() {
	if err := Main(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
