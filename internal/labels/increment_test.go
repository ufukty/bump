package labels

import (
	"maps"
	"slices"
	"testing"
)

func TestIncrement_Major(t *testing.T) {
	tcs := map[string]string{
		"v0.0.0": "v1.0.0",
		"v0.0.1": "v1.0.0",
		"v0.1.0": "v1.0.0",
		"v1.0.0": "v2.0.0",
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			expected := tcs[input]
			got, err := Increment(input, Major)
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}
			if expected != got {
				t.Errorf("expected %q got %q", expected, got)
			}
		})
	}
}

func TestIncrement_Minor(t *testing.T) {
	tcs := map[string]string{
		"v0.0.0": "v0.1.0",
		"v0.0.1": "v0.1.0",
		"v0.1.0": "v0.2.0",
		"v1.0.0": "v1.1.0",
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			expected := tcs[input]
			got, err := Increment(input, Minor)
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}
			if expected != got {
				t.Errorf("expected %q got %q", expected, got)
			}
		})
	}
}

func TestIncrement_Patch(t *testing.T) {
	tcs := map[string]string{
		"v0.0.0": "v0.0.1",
		"v0.0.1": "v0.0.2",
		"v0.1.0": "v0.1.1",
		"v1.0.0": "v1.0.1",
	}

	for _, input := range slices.Sorted(maps.Keys(tcs)) {
		t.Run(input, func(t *testing.T) {
			expected := tcs[input]
			got, err := Increment(input, Patch)
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}
			if expected != got {
				t.Errorf("expected %q got %q", expected, got)
			}
		})
	}
}
