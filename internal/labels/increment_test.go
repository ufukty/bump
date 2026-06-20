package labels

import (
	"cmp"
	"fmt"
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

func TestIncrement_majorToVersionOnePositive(t *testing.T) {
	tcs := map[Labels]Labels{
		{0, 0, 0, 0}: {1, 0, 0, 0},
		{0, 0, 0, 1}: {1, 0, 0, 0},
		{0, 0, 1, 0}: {1, 0, 0, 0},
		{0, 0, 1, 1}: {1, 0, 0, 0},
		{0, 1, 0, 0}: {1, 0, 0, 0},
		{0, 1, 0, 1}: {1, 0, 0, 0},
	}

	for _, input := range sort(maps.Keys(tcs)) {
		t.Run(input.String(), func(t *testing.T) {
			expected := tcs[input]
			got, err := NextMajor(input, true)
			if err != nil {
				t.Fatalf("act, unexpected error on forcing to v1.0.0: %v", err)
			}
			if expected != got {
				t.Errorf("expected %q got %q", expected, got)
			}
		})
	}
}

func TestIncrement_majorToVersionOneNegative(t *testing.T) {
	tcs := []Labels{
		{0, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 1, 0},
		{0, 0, 1, 1},
		{0, 1, 0, 0},
		{0, 1, 0, 1},
	}

	for _, input := range tcs {
		t.Run(input.String(), func(t *testing.T) {
			_, err := NextMajor(input, false)
			if err == nil {
				t.Fatalf("act, unexpected success. Increment should reject issuing v1.0.0 without the arg")
			} else if err != ErrBackwardsCompatibilityPromise {
				t.Fatalf("act, expected %v got %v", ErrBackwardsCompatibilityPromise, err)
			}
		})
	}
}

func TestIncrement_major(t *testing.T) {
	tcs := map[Labels]Labels{
		{1, 0, 0, 0}: {2, 0, 0, 0},
		{1, 0, 0, 1}: {2, 0, 0, 0},
		{1, 0, 1, 0}: {2, 0, 0, 0},
		{1, 0, 1, 1}: {2, 0, 0, 0},
		{1, 1, 0, 0}: {2, 0, 0, 0},
		{1, 1, 0, 1}: {2, 0, 0, 0},
		{2, 0, 0, 0}: {3, 0, 0, 0},
		{2, 0, 0, 1}: {3, 0, 0, 0},
	}

	for _, input := range sort(maps.Keys(tcs)) {
		t.Run(input.String(), func(t *testing.T) {
			expected := tcs[input]
			got, err := NextMajor(input, false)
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}
			if expected != got {
				t.Errorf("expected %q got %q", expected, got)
			}
		})
	}
}

func TestIncrement_minor(t *testing.T) {
	tcs := map[Labels]Labels{
		{0, 0, 0, 0}: {0, 1, 0, 0},
		{0, 0, 0, 1}: {0, 1, 0, 0},
		{0, 0, 1, 0}: {0, 1, 0, 0},
		{0, 0, 1, 1}: {0, 1, 0, 0},
		{0, 1, 0, 0}: {0, 2, 0, 0},
		{0, 1, 0, 1}: {0, 2, 0, 0},
		{1, 0, 0, 0}: {1, 1, 0, 0},
		{1, 0, 0, 1}: {1, 1, 0, 0},
	}

	for _, input := range sort(maps.Keys(tcs)) {
		t.Run(input.String(), func(t *testing.T) {
			expected := tcs[input]
			got := NextMinor(input)
			if expected != got {
				t.Errorf("expected %q got %q", expected, got)
			}
		})
	}
}

func TestIncrement_patch(t *testing.T) {
	tcs := map[Labels]Labels{
		{0, 0, 0, 0}: {0, 0, 1, 0},
		{0, 0, 0, 1}: {0, 0, 1, 0},
		{0, 0, 1, 0}: {0, 0, 2, 0},
		{0, 0, 1, 1}: {0, 0, 2, 0},
		{0, 1, 0, 0}: {0, 1, 1, 0},
		{0, 1, 0, 1}: {0, 1, 1, 0},
		{1, 0, 0, 0}: {1, 0, 1, 0},
		{1, 0, 0, 1}: {1, 0, 1, 0},
	}

	for _, input := range sort(maps.Keys(tcs)) {
		t.Run(input.String(), func(t *testing.T) {
			expected := tcs[input]
			got := NextPatch(input)
			if expected != got {
				t.Errorf("expected %q got %q", expected, got)
			}
		})
	}
}

func TestIncrement_iterateAlpha(t *testing.T) {
	tcs := map[Labels]Labels{
		{0, 0, 0, 1}: {0, 0, 0, 2},
		{0, 0, 0, 8}: {0, 0, 0, 9},
		{0, 0, 1, 1}: {0, 0, 1, 2},
		{0, 0, 1, 8}: {0, 0, 1, 9},
		{0, 1, 0, 1}: {0, 1, 0, 2},
		{0, 1, 0, 8}: {0, 1, 0, 9},
		{1, 0, 0, 1}: {1, 0, 0, 2},
		{1, 0, 0, 8}: {1, 0, 0, 9},
	}

	for _, input := range sort(maps.Keys(tcs)) {
		t.Run(input.String(), func(t *testing.T) {
			expected := tcs[input]
			got, err := IterateAlphaTrack(input)
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}
			if expected != got {
				t.Errorf("expected %q got %q", expected, got)
			}
		})
	}
}

func TestIncrement_iterateAlphaWithoutTrack(t *testing.T) {
	tcs := []Labels{
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 1, 0, 0},
		{1, 0, 0, 0},
	}

	for _, input := range tcs {
		t.Run(input.String(), func(t *testing.T) {
			_, err := IterateAlphaTrack(input)
			if err == nil {
				t.Fatalf("act, unexpected success")
			}
		})
	}
}

func TestIncrement_initalizeAlphaTrack(t *testing.T) {
	type input struct {
		Current Labels
		Target  Label
	}
	type tc struct {
		input
		expected Labels
	}
	tcs := []tc{
		{input: input{Current: Labels{0, 0, 0, 0}, Target: Major}, expected: Labels{1, 0, 0, 1}},
		{input: input{Current: Labels{0, 0, 0, 0}, Target: Minor}, expected: Labels{0, 1, 0, 1}},
		{input: input{Current: Labels{0, 0, 0, 0}, Target: Patch}, expected: Labels{0, 0, 1, 1}},
		{input: input{Current: Labels{1, 2, 3, 0}, Target: Major}, expected: Labels{2, 0, 0, 1}},
		{input: input{Current: Labels{1, 2, 3, 0}, Target: Minor}, expected: Labels{1, 3, 0, 1}},
		{input: input{Current: Labels{1, 2, 3, 0}, Target: Patch}, expected: Labels{1, 2, 4, 1}},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%s on %s", tc.Target, tc.input.Current), func(t *testing.T) {
			got, err := InitAlphaTrack(tc.input.Current, tc.Target, true)
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}
			if tc.expected != got {
				t.Errorf("expected %q got %q", tc.expected, got)
			}
		})
	}
}

func TestIncrement_finalizeAlphaTrack(t *testing.T) {
	current, expected := Labels{0, 1, 0, 1}, Labels{0, 1, 0, 0}
	got, err := FinalizeAlphaTrack(current, false)
	if err != nil {
		t.Fatalf("act, unexpected error: %v", err)
	}
	if expected != got {
		t.Errorf("expected %q got %q", expected, got)
	}
}

func TestIncrement_finalizeWithoutAlphaTrack(t *testing.T) {
	got, err := FinalizeAlphaTrack(Labels{0, 1, 0, 0}, false)
	if err == nil {
		t.Fatalf("act, unexpected success: %q", got)
	}
}
