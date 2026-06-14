package labels

import (
	"fmt"
	"slices"

	"github.com/ufukty/bump/internal/args"
)

const (
	Major = "major"
	Minor = "minor"
	Patch = "patch"
	Alpha = "alpha"
)

var Mods = []string{Major, Minor, Patch, Alpha}

func increment(version Labels, label int) Labels {
	version[label] += 1
	for j := label + 1; j < 4; j++ { // reseting digits carried over
		version[j] = 0
	}
	return version
}

func equal(a, b Labels) bool {
	return a[0] == b[0] &&
		a[1] == b[1] &&
		a[2] == b[2] &&
		a[3] == b[3]
}

var v1 = Labels{1, 0, 0, 0}

var ErrAccidentalVersionOne = fmt.Errorf("unforced leaving of zero versions")

func Increment(version Labels, args *args.Args) (Labels, error) {
	l := slices.Index(Mods, args.Command)
	if l == -1 {
		return Labels{}, fmt.Errorf("unknown label: %s", args.Command)
	}
	next := increment(version, l)
	if !args.Force && equal(next, v1) {
		return Labels{}, ErrAccidentalVersionOne
	}
	return next, nil
}
