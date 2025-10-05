package labels

import (
	"fmt"
	"slices"
)

const (
	Major = "major"
	Minor = "minor"
	Patch = "patch"
)

var Mods = []string{Major, Minor, Patch}

func Increment(version Labels, label string) (Labels, error) {
	l := slices.Index(Mods, label)
	if l == -1 {
		return Labels{}, fmt.Errorf("unknown label: %s", label)
	}

	version[l] += 1
	for j := l + 1; j < 4; j++ { // reseting digits carried over
		version[j] = 0
	}
	return version, nil
}
