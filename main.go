package main

import (
	_ "embed"
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
	Help  bool
}

var ErrNoLabel = fmt.Errorf("CLI args don't mention the target label name")

//go:embed synopsis.txt
var synopsis string

func args(arguments []string) (Args, error) {
	fs := flag.NewFlagSet("bump", flag.ExitOnError)
	as := Args{}
	fs.BoolVar(&as.Force, "force", false, "Use for the major command to allow incrementing the version number up to v1.0.0")
	fs.BoolVar(&as.Help, "help", false, "prints usage information and exits")
	err := fs.Parse(arguments)
	if err != nil {
		return Args{}, fmt.Errorf("parsing flags: %w", err)
	}
	if as.Help {
		fmt.Print(synopsis)
		fs.PrintDefaults()
	} else if fs.NArg() < 1 {
		return Args{}, fmt.Errorf("missing the label: %s", strings.Join(labels.Mods, ", "))
	} else if fs.NArg() > 0 {
		as.Label = fs.Arg(0)
	}
	return as, nil
}

func Main() error {
	args, err := args(os.Args[1:])
	if err != nil {
		return fmt.Errorf("reading args: %w", err)
	} else if args.Help {
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

	l2, err := labels.Increment(l1, labels.Args{Label: args.Label, Force: args.Force})
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
		case errors.Is(err, ErrNoLabel):
			fmt.Println("label name not found. provide any of: ", strings.Join(labels.Mods, ", "))
		case errors.Is(err, labels.ErrAccidentalVersionOne):
			fmt.Println("bump prevents accidentally leaving the zero versions. Run: bump --help")
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}
}
