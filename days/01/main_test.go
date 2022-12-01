package main

import (
	"reflect"
	"testing"
)

func TestCalorieSums(t *testing.T) {
	got := calorieSums("input_test.txt")
	want := []int{24000, 11000, 10000, 6000, 4000}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func TestPart1(t *testing.T) {
	got := Part1("input_test.txt")
	want := 24000
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
func TestPart2(t *testing.T) {
	got := Part2("input_test.txt")
	want := 45000
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
