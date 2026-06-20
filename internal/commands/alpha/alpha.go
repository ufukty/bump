package alpha

import (
	"fmt"

	"github.com/ufukty/bump/internal/commands/alpha/finalize"
	"github.com/ufukty/bump/internal/commands/alpha/iterate"
	"github.com/ufukty/bump/internal/commands/alpha/major"
	"github.com/ufukty/bump/internal/commands/alpha/minor"
	"github.com/ufukty/bump/internal/commands/alpha/patch"
)

func Dispatch(args []string) error {
	switch subcommand, rest := args[0], args[1:]; subcommand {
	case "":
		if err := iterate.Run(); err != nil {
			return fmt.Errorf("iterate: %w", err)
		}
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
	case "finalize":
		if err := finalize.Run(rest); err != nil {
			return fmt.Errorf("finalize: %w", err)
		}
	default:
		return fmt.Errorf("unknown subcommand: %s", subcommand)
	}
	return nil
}
