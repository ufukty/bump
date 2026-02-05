package labels

import (
	"fmt"
	"slices"
)

const (
	Major = "major"
	Minor = "minor"
	Patch = "patch"
	Alpha = "alpha"
)

var Mods = []string{Major, Minor, Patch, Alpha}

type Args struct {
	Label string
	Force bool
}

func increment(version Labels, label int) Labels {
	version[label] += 1
	for j := label + 1; j < 4; j++ { // reseting digits carried over
		version[j] = 0
	}
	return version
}

func Increment(version Labels, args Args) (Labels, error) {
	l := slices.Index(Mods, args.Label)
	if l == -1 {
		return Labels{}, fmt.Errorf("unknown label: %s", args.Label)
	}
	return increment(version, l), nil
}
