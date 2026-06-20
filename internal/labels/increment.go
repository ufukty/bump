package labels

import (
	"fmt"
)

type Label string

const (
	Major = Label("major")
	Minor = Label("minor")
	Patch = Label("patch")
	Alpha = Label("alpha")
)

func increment(version Labels, label int) Labels {
	version[label] += 1
	for j := label + 1; j < 4; j++ { // resetting digits carried over
		version[j] = 0
	}
	return version
}

var (
	ErrLandingOnV1WithoutForce = fmt.Errorf("landing on v1.0.0 requires --force flag")
	ErrTargetingV1WithoutForce = fmt.Errorf("targeting v1.0.0 requires --force flag")
)

func NextMajor(version Labels, forced bool) (Labels, error) {
	next := increment(version, index(Major))
	if next == V1 && !forced {
		return Labels{}, ErrLandingOnV1WithoutForce
	}
	return next, nil
}

func NextMinor(version Labels) Labels {
	return increment(version, index(Minor))
}

func NextPatch(version Labels) Labels {
	return increment(version, index(Patch))
}

var ErrCommandRequiresAlphaTrack = fmt.Errorf("command requires an active alpha-track")

func IterateAlphaTrack(version Labels) (Labels, error) {
	if version[3] == 0 {
		return Labels{}, ErrCommandRequiresAlphaTrack
	}
	return increment(version, index(Alpha)), nil
}

func index(label Label) int {
	return map[Label]int{
		Major: 0,
		Minor: 1,
		Patch: 2,
		Alpha: 3,
	}[label]
}

func InitAlphaTrack(version Labels, target Label, forced bool) (Labels, error) {
	if version[3] > 0 {
		return Labels{}, fmt.Errorf("cannot initiate an alpha-track from an alpha version")
	}
	next := increment(version, index(target))
	if next == V1 && !forced {
		return Labels{}, ErrTargetingV1WithoutForce
	}
	next = increment(next, index(Alpha))
	return next, nil
}

func FinalizeAlphaTrack(version Labels, forced bool) (Labels, error) {
	if version[3] == 0 {
		return Labels{}, ErrCommandRequiresAlphaTrack
	}
	next := version
	next[3] = 0
	if next == V1 && !forced {
		return Labels{}, ErrLandingOnV1WithoutForce
	}
	return next, nil
}
