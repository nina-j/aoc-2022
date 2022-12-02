package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Part1("input_test")
	want := 15
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2("input_test")
	want := 12
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
