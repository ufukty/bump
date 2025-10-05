package labels

import (
	"testing"
)

func TestParse(t *testing.T) {
	type tc struct {
		input    string
		expected Labels
	}
	tcs := []tc{
		{input: "v0.0.0", expected: Labels{0, 0, 0, 0}},
		{input: "v0.0.1", expected: Labels{0, 0, 1, 0}},
		{input: "v0.1.0", expected: Labels{0, 1, 0, 0}},
		{input: "v1.0.0", expected: Labels{1, 0, 0, 0}},

		{input: "v0.0.0.1", expected: Labels{0, 0, 0, 1}},
		{input: "v0.0.1.1", expected: Labels{0, 0, 1, 1}},
		{input: "v0.1.0.1", expected: Labels{0, 1, 0, 1}},
		{input: "v1.0.0.1", expected: Labels{1, 0, 0, 1}},
	}

	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {

			got, err := Parse(tc.input)
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}
			if tc.expected != got {
				t.Errorf("expected %q got %q", tc.expected, got)
			}
		})
	}
}

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
		t.Run(tc.expected, func(t *testing.T) {
			got := tc.input.String()
			if tc.expected != got {
				t.Errorf("expected %q got %q", tc.expected, got)
			}
		})
	}
}
