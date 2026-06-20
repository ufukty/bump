package minor

import (
	"fmt"

	"github.com/ufukty/bump/internal/git"
	"github.com/ufukty/bump/internal/labels"
)

func issue() error {
	v1, err := git.Describe()
	if err != nil {
		return fmt.Errorf("git describe: %w", err)
	}
	l1, err := labels.Parse(v1)
	if err != nil {
		return fmt.Errorf("parsing current version: %w", err)
	}
	next, err := labels.NextMinor(l1)
	if err != nil {
		return fmt.Errorf("determining the next vesion tag: %w", err)
	}
	if err := git.Tag(next.String()); err != nil {
		return fmt.Errorf("git tag: %w", err)
	}
	return nil
}

func Run() error {
	if err := issue(); err != nil {
		return fmt.Errorf("issue: %w", err)
	}
	return nil
}
