package main

import (
	"testing"

	"github.com/ufukty/bump/internal/labels"
)

func TestArgs_ForceMajor(t *testing.T) {
	got, err := args([]string{"--force", "major"})
	if err != nil {
		t.Fatalf("act, unexpected error: %v", err)
	}
	expected := Args{
		Label: labels.Major,
		Force: true,
	}
	if got != expected {
		t.Errorf("assert, expected %#v got %#v", expected, got)
	}
}

func TestArgs_Major(t *testing.T) {
	got, err := args([]string{"major"})
	if err != nil {
		t.Fatalf("act, unexpected error: %v", err)
	}
	expected := Args{
		Label: labels.Major,
	}
	if got != expected {
		t.Errorf("assert, expected %#v got %#v", expected, got)
	}
}

func TestArgs_Help(t *testing.T) {
	got, err := args([]string{"--help"})
	if err != nil {
		t.Fatalf("act, unexpected error: %v", err)
	}
	expected := Args{
		Help: true,
	}
	if got != expected {
		t.Errorf("assert, expected %#v got %#v", expected, got)
	}
}
