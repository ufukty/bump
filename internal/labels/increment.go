package labels

import (
	"fmt"
	"slices"

	"github.com/ufukty/bump/internal/args"
)

const (
	Major    = "major"    // label + command
	Minor    = "minor"    // label + command
	Patch    = "patch"    // label + command
	Alpha    = "alpha"    // label + command
	Finalize = "finalize" // command
)

func increment(version Labels, label int) Labels {
	version[label] += 1
	for j := label + 1; j < 4; j++ { // resetting digits carried over
		version[j] = 0
	}
	return version
}

func isStandard(label string) bool { return slices.Contains([]string{Major, Minor, Patch}, label) }
func index(label string) int       { return slices.Index([]string{Major, Minor, Patch, Alpha}, label) }

func iteratesReleaseVersion(a *args.Args) bool { return isStandard(a.Command) }
func initsAlphaTrack(a *args.Args) bool        { return a.Command == Alpha && isStandard(a.AlphaTrackTarget) }
func iteratesAlphaTrack(a *args.Args) bool     { return a.Command == Alpha && a.AlphaTrackTarget == "" }
func finalizesAlphaTrack(a *args.Args) bool    { return a.Command == Finalize }

func iterateReleaseVersion(version Labels, args *args.Args) (Labels, error) {
	return increment(version, index(args.Command)), nil
}

var ErrCommandRequiresAlphaTrack = fmt.Errorf("command requires an active alpha-track")

func iterateAlphaTrack(version Labels) (Labels, error) {
	if version[3] == 0 {
		return Labels{}, ErrCommandRequiresAlphaTrack
	}
	return increment(version, 3), nil
}

func initAlphaTrack(version Labels, args *args.Args) (Labels, error) {
	if version[3] > 0 {
		return Labels{}, fmt.Errorf("cannot initiate an alpha-track from an alpha version")
	}
	next := increment(version, index(args.AlphaTrackTarget))
	next = increment(next, 3)
	return next, nil
}

func finalizeAlphaTrack(version Labels) (Labels, error) {
	if version[3] == 0 {
		return Labels{}, ErrCommandRequiresAlphaTrack
	}
	next := version
	next[3] = 0
	return next, nil
}

func dispatch(version Labels, args *args.Args) (Labels, error) {
	switch {
	case iteratesReleaseVersion(args):
		next, err := iterateReleaseVersion(version, args)
		if err != nil {
			return Labels{}, fmt.Errorf("iterating release version: %w", err)
		}
		return next, nil

	case initsAlphaTrack(args):
		next, err := initAlphaTrack(version, args)
		if err != nil {
			return Labels{}, fmt.Errorf("initializing alpha track: %w", err)
		}
		return next, nil

	case iteratesAlphaTrack(args):
		next, err := iterateAlphaTrack(version)
		if err != nil {
			return Labels{}, fmt.Errorf("iterating alpha track: %w", err)
		}
		return next, nil

	case finalizesAlphaTrack(args):
		next, err := finalizeAlphaTrack(version)
		if err != nil {
			return Labels{}, fmt.Errorf("finalizing alpha track: %w", err)
		}
		return next, nil

	default:
		return Labels{}, fmt.Errorf("unsupported combination of command (%q) and arguments", args.Command)
	}
}

var v1 = Labels{1, 0, 0, 0}

var ErrAccidentalVersionOne = fmt.Errorf("landing on v1.0.0 requires --force flag")

func Increment(version Labels, args *args.Args) (Labels, error) {
	next, err := dispatch(version, args)
	if err != nil {
		return Labels{}, fmt.Errorf("dispatch: %w", err)
	}
	if !args.Force && next == v1 {
		return Labels{}, ErrAccidentalVersionOne
	}
	return next, nil
}
