package args

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Args struct {
	Command          string
	Force            bool   // valid when command is major or finalize
	AlphaTrackTarget string // valid when command is alpha
}

func parse(osArgs []string) (*Args, error) {
	if len(osArgs) < 1 {
		return nil, fmt.Errorf("not enough args. run: bump help")
	}
	command, remaining := osArgs[0], osArgs[1:]
	if !slices.Contains([]string{"major", "minor", "patch", "alpha", "finalize", "help"}, command) {
		return nil, fmt.Errorf("unknown command %q. run: bump help", command)
	}

	args := &Args{Command: command}
	switch command {
	case "major", "finalize":
		if len(remaining) > 0 && strings.TrimSpace(remaining[0]) == "--force" {
			args.Force = true
		}

	case "alpha":
		if len(remaining) > 0 {
			track := remaining[0]
			if !slices.Contains([]string{"major", "minor", "patch"}, strings.TrimSpace(track)) {
				return nil, fmt.Errorf("unknown ")
			}
			args.AlphaTrackTarget = track
		}
	}

	return args, nil
}

func Command() (*Args, error) {
	return parse(os.Args[1:])
}
