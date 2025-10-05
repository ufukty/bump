package labels

import (
	"testing"
)

func TestLabels_String(t *testing.T) {
	type tc struct {
		input    Labels
		expected string
	}
	tcs := []tc{
		{input: Labels{0, 0, 0, 0}, expected: "v0.0.0"},
		{input: Labels{0, 0, 1, 0}, expected: "v0.0.1"},
		{input: Labels{0, 1, 0, 0}, expected: "v0.1.0"},
		{input: Labels{1, 0, 0, 0}, expected: "v1.0.0"},

		{input: Labels{0, 0, 0, 1}, expected: "v0.0.0.1"},
		{input: Labels{0, 0, 1, 1}, expected: "v0.0.1.1"},
		{input: Labels{0, 1, 0, 1}, expected: "v0.1.0.1"},
		{input: Labels{1, 0, 0, 1}, expected: "v1.0.0.1"},
	}

	for _, tc := range tcs {
		got := tc.input.String()
		if tc.expected != got {
			t.Errorf("expected %q got %q", tc.expected, got)
		}
	}
}
