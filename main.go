package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/ufukty/bump/internal/args"
	"github.com/ufukty/bump/internal/git"
	"github.com/ufukty/bump/internal/labels"
)

//go:embed synopsis.txt
var synopsis string

func Main() error {
	args, err := args.Command()
	if err != nil {
		return fmt.Errorf("args: %w", err)
	}

	if args.Command == "help" {
		fmt.Println(synopsis)
		return nil
	}

	v1, err := git.Describe()
	if err != nil {
		return fmt.Errorf("git describe: %w", err)
	}

	l1, err := labels.Parse(v1)
	if err != nil {
		return fmt.Errorf("parsing current version: %w", err)
	}

	l2, err := labels.Increment(l1, args)
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
