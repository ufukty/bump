package finalize

import (
	"flag"
	"fmt"

	"github.com/ufukty/bump/internal/git"
	"github.com/ufukty/bump/internal/labels"
)

type args struct {
	Force bool
}

func parse(osArgs []string) (*args, error) {
	args := &args{}
	set := flag.NewFlagSet("bump alpha finalize", flag.ExitOnError)
	set.BoolVar(&args.Force, "force", false, "disables the accidental backwards-compatibility promise prevention")
	err := set.Parse(osArgs)
	if err != nil {
		return nil, err
	}
	return args, nil
}

func issue(args *args) error {
	v1, err := git.Describe()
	if err != nil {
		return fmt.Errorf("git describe: %w", err)
	}
	l1, err := labels.Parse(v1)
	if err != nil {
		return fmt.Errorf("parsing current version: %w", err)
	}
	next, err := labels.FinalizeAlphaTrack(l1, args.Force)
	if err != nil {
		return fmt.Errorf("determining the next version tag: %w", err)
	}
	if err := git.Tag(next.String()); err != nil {
		return fmt.Errorf("git tag: %w", err)
	}
	return nil
}

func Run(osArgs []string) error {
	args, err := parse(osArgs)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}
	err = issue(args)
	if err != nil {
		return fmt.Errorf("issue: %w", err)
	}
	return nil
}
