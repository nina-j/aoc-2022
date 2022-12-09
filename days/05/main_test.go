package main

import "testing"

// test crates
//    [D]
//[N] [C]
//[Z] [M] [P]
// 1   2   3

func TestPart1(t *testing.T) {
	crates := map[int][]string{
		1: {"Z", "N"},
		2: {"M", "C", "D"},
		3: {"P"},
	}
	got := Part1("input_test", crates)
	want := "CMZ"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	crates := map[int][]string{
		1: {"Z", "N"},
		2: {"M", "C", "D"},
		3: {"P"},
	}
	got := Part2("input_test", crates)
	want := "MCD"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
