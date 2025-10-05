package labels

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var extractor = regexp.MustCompile(`v([0-9]+)\.([0-9]+)\.([0-9]+).*`)

type Labels [4]int

func Parse(verstr string) (Labels, error) {
	ms := extractor.FindStringSubmatch(verstr)
	if len(ms) != 4 {
		return Labels{}, fmt.Errorf("expected to see 'major.minor.patch' format: %s", verstr)
	}
	ms = ms[1:]

	is := [4]int{}
	for i, m := range ms {
		n, err := strconv.Atoi(m)
		if err != nil {
			return Labels{}, fmt.Errorf("parsing integer: %w", err)
		}
		is[i] = n
	}

	return is, nil
}

func (l Labels) String() string {
	is := []string{}
	for i := range 4 {
		if i == 3 && l[3] == 0 { // omit no-alpha
			continue
		}
		is = append(is, strconv.Itoa(l[i]))
	}
	return "v" + strings.Join(is, ".")
}
