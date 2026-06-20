package labels

import (
	"fmt"
	"regexp"
	"strconv"
)

var extractor = regexp.MustCompile(`v([0-9]+)\.([0-9]+)\.([0-9]+)(?:-alpha.([0-9]+))?.*`)

type Labels [4]int // [major, minor, patch, alpha]

var V1 = Labels{1, 0, 0, 0}

func Parse(verstr string) (Labels, error) {
	ms := extractor.FindStringSubmatch(verstr)
	if len(ms) != 5 {
		return Labels{}, fmt.Errorf("expected to see at least 'major.minor.patch' format: %s", verstr)
	}
	ms = ms[1:]

	is := [4]int{}
	for i, m := range ms {
		if m != "" {
			n, err := strconv.Atoi(m)
			if err != nil {
				return Labels{}, fmt.Errorf("parsing integer: %w", err)
			}
			is[i] = n
		}
	}

	return is, nil
}

func (l Labels) String() string {
	v := fmt.Sprintf("v%d.%d.%d", l[0], l[1], l[2])
	if l[3] > 0 {
		v = fmt.Sprintf("%s-alpha.%d", v, l[3])
	}
	return v
}
