package labels

import (
	"cmp"
	"fmt"
	"iter"
	"maps"
	"slices"
	"testing"

	"github.com/ufukty/bump/internal/args"
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
			got, err := Increment(input, &args.Args{Command: Major, Force: true})
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
			_, err := Increment(input, &args.Args{Command: Major})
			if err == nil {
				t.Fatalf("act, unexpected success. Increment should reject issuing v1.0.0 without the arg")
			} else if err != ErrAccidentalVersionOne {
				t.Fatalf("act, expected %v got %v", ErrAccidentalVersionOne, err)
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
			got, err := Increment(input, &args.Args{Command: Major})
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
			got, err := Increment(input, &args.Args{Command: Minor})
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}
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
			got, err := Increment(input, &args.Args{Command: Patch})
			if err != nil {
				t.Fatalf("act, unexpected error: %v", err)
			}
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
			got, err := Increment(input, &args.Args{Command: Alpha})
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
			_, err := Increment(input, &args.Args{Command: Alpha})
			if err == nil {
				t.Fatalf("act, unexpected success")
			}
		})
	}
}

func TestIncrement_initalizeAlphaTrack(t *testing.T) {
	type input struct {
		Current     Labels
		AlphaTarget string
	}
	type tc struct {
		input
		expected Labels
	}
	tcs := []tc{
		{input: input{Current: Labels{0, 0, 0, 0}, AlphaTarget: Major}, expected: Labels{1, 0, 0, 1}},
		{input: input{Current: Labels{0, 0, 0, 0}, AlphaTarget: Minor}, expected: Labels{0, 1, 0, 1}},
		{input: input{Current: Labels{0, 0, 0, 0}, AlphaTarget: Patch}, expected: Labels{0, 0, 1, 1}},
		{input: input{Current: Labels{1, 2, 3, 0}, AlphaTarget: Major}, expected: Labels{2, 0, 0, 1}},
		{input: input{Current: Labels{1, 2, 3, 0}, AlphaTarget: Minor}, expected: Labels{1, 3, 0, 1}},
		{input: input{Current: Labels{1, 2, 3, 0}, AlphaTarget: Patch}, expected: Labels{1, 2, 4, 1}},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%s on %s", tc.AlphaTarget, tc.input.Current), func(t *testing.T) {
			got, err := Increment(tc.input.Current, &args.Args{Command: Alpha, AlphaTrackTarget: tc.AlphaTarget})
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
	got, err := Increment(current, &args.Args{Command: Finalize})
	if err != nil {
		t.Fatalf("act, unexpected error: %v", err)
	}
	if expected != got {
		t.Errorf("expected %q got %q", expected, got)
	}
}

func TestIncrement_finalizeWithoutAlphaTrack(t *testing.T) {
	got, err := Increment(Labels{0, 1, 0, 0}, &args.Args{Command: Finalize})
	if err == nil {
		t.Fatalf("act, unexpected success: %q", got)
	}
}
