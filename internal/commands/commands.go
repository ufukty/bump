package commands

import (
	"fmt"

	"github.com/ufukty/bump/internal/commands/alpha"
	"github.com/ufukty/bump/internal/commands/help"
	"github.com/ufukty/bump/internal/commands/major"
	"github.com/ufukty/bump/internal/commands/minor"
	"github.com/ufukty/bump/internal/commands/patch"
)

func Dispatch(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("not enough args")
	}
	switch command, rest := args[0], args[1:]; command {
	case "major":
		if err := major.Run(rest); err != nil {
			return fmt.Errorf("major: %w", err)
		}
	case "minor":
		if err := minor.Run(); err != nil {
			return fmt.Errorf("minor: %w", err)
		}
	case "patch":
		if err := patch.Run(); err != nil {
			return fmt.Errorf("patch: %w", err)
		}
	case "alpha":
		if err := alpha.Dispatch(rest); err != nil {
			return fmt.Errorf("alpha: %w", err)
		}
	case "help":
		if err := help.Run(); err != nil {
			return fmt.Errorf("help: %w", err)
		}
	default:
		return fmt.Errorf("unknown command: %q", command)
	}
	return nil
}
