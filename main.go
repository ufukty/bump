package main

import (
	_ "embed"
	"errors"
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
		switch {
		case errors.Is(err, labels.ErrAccidentalVersionOne):
			fmt.Println("bump prevents accidentally leaving the zero versions. Run: bump --help")
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}
}
