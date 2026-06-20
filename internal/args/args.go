package args

import "fmt"

func Dispatch(osArgs []string) error {
	if len(osArgs) < 1 {
		return fmt.Errorf("not enough args. run: bump help")
	}
	command, remaining := osArgs[0], osArgs[1:]
	switch command {
	case "major":
		if err := major.Run(remaining); err != nil {
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
		if err := alpha.Run(remaining); err != nil {
			return fmt.Errorf("alpha: %w", err)
		}
	case "help":
		if err := help.Run(); err != nil {
			return fmt.Errorf("help: %w", err)
		}
	default:
		return fmt.Errorf("unknown command %q. run: bump help", command)
	}
	return nil
}
