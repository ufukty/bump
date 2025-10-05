package labels

import (
	"cmp"
	"iter"
	"maps"
	"slices"
	"testing"
)

// to stablize ordering of test cases
func sort(cases iter.Seq[Labels]) []Labels {
	return slices.SortedFunc(cases, func(a, b Labels) int {
		return cmp.Or(
			cmp.Compare(a[0], b[0]),
			cmp.Compare(a[1], b[1]),
			cmp.Compare(a[2], b[2]),
			cmp.Compare(a[3], b[3]),
		)
	})
}

func TestIncrement_Major(t *testing.T) {
	tcs := map[Labels]Labels{
		{0, 0, 0}: {1, 0, 0},
		{0, 0, 1}: {1, 0, 0},
		{0, 1, 0}: {1, 0, 0},
		{1, 0, 0}: {2, 0, 0},
	}

	for _, input := range sort(maps.Keys(tcs)) {
		t.Run(input.String(), func(t *testing.T) {
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
	tcs := map[Labels]Labels{
		{0, 0, 0}: {0, 1, 0},
		{0, 0, 1}: {0, 1, 0},
		{0, 1, 0}: {0, 2, 0},
		{1, 0, 0}: {1, 1, 0},
	}

	for _, input := range sort(maps.Keys(tcs)) {
		t.Run(input.String(), func(t *testing.T) {
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
	tcs := map[Labels]Labels{
		{0, 0, 0}: {0, 0, 1},
		{0, 0, 1}: {0, 0, 2},
		{0, 1, 0}: {0, 1, 1},
		{1, 0, 0}: {1, 0, 1},
	}

	for _, input := range sort(maps.Keys(tcs)) {
		t.Run(input.String(), func(t *testing.T) {
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
