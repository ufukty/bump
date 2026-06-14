package args

import (
	"testing"
)

func TestArgs_ForceMajor(t *testing.T) {
	got, err := parse([]string{"major", "--force"})
	if err != nil {
		t.Fatalf("act, unexpected error: %v", err)
	}
	expected := Args{
		Command: "major",
		Force:   true,
	}
	if *got != expected {
		t.Errorf("assert, expected %#v got %#v", expected, got)
	}
}

func TestArgs_Major(t *testing.T) {
	got, err := parse([]string{"major"})
	if err != nil {
		t.Fatalf("act, unexpected error: %v", err)
	}
	expected := Args{
		Command: "major",
	}
	if *got != expected {
		t.Errorf("assert, expected %#v got %#v", expected, got)
	}
}

func TestArgs_Help(t *testing.T) {
	got, err := parse([]string{"help"})
	if err != nil {
		t.Fatalf("act, unexpected error: %v", err)
	}
	expected := Args{
		Command: "help",
	}
	if *got != expected {
		t.Errorf("assert, expected %#v got %#v", expected, got)
	}
}
