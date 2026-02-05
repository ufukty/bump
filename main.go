package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ufukty/bump/internal/git"
	"github.com/ufukty/bump/internal/labels"
)

type Args struct {
	Label string
	Force bool
}

func args() (Args, error) {
	args := Args{}
	flag.BoolVar(&args.Force, "force", false, "force incrementing the version to v1.0.0 with major command")
	flag.Parse()
	if flag.NArg() != 1 {
		return Args{}, fmt.Errorf("expected to see one argument among: %s", strings.Join(labels.Mods, ", "))
	}
	args.Label = flag.Arg(1)
	return args, nil
}

func Main() error {
	args, err := args()
	if err != nil {
		return fmt.Errorf("reading args: %w", err)
	}

	v1, err := git.Describe()
	if err != nil {
		return fmt.Errorf("git describe: %w", err)
	}

	l1, err := labels.Parse(v1)
	if err != nil {
		return fmt.Errorf("parsing current version: %w", err)
	}

	l2, err := labels.Increment(l1, labels.Args(args))
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
