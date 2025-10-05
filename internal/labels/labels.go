package labels

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var Mods = []string{"major", "minor", "patch"}

var extractor = regexp.MustCompile(`v([0-9]+)\.([0-9]+)\.([0-9]+).*`)

func Increment(verstr, label string) (string, error) {
	ms := extractor.FindStringSubmatch(verstr)
	if len(ms) != 4 {
		return "", fmt.Errorf("expected to see 'major.minor.patch' format: %s", verstr)
	}
	ms = ms[1:]

	i := slices.Index(Mods, label)
	if i == -1 {
		return "", fmt.Errorf("invalid argument. available arguments: %s", strings.Join(Mods, ", "))
	}

	n, err := strconv.Atoi(ms[i])
	if err != nil {
		return "", fmt.Errorf("parsing integer: %w", err)
	}
	ms[i] = fmt.Sprintf("%d", n+1)
	for j := i + 1; j < 3; j++ {
		ms[j] = "0"
	}

	return "v" + strings.Join(ms, "."), nil
}
