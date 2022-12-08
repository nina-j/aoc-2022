package main

import (
	"aoc/util"
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
//go:embed input_test
var fsys embed.FS

func rangeToMap(input []string) map[int]bool {
	m := make(map[int]bool)
	min, err := strconv.Atoi(input[0])
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(input[1])
	if err != nil {
		panic(err)
	}
	for i := min; i <= max; i++ {
		m[i] = true
	}
	return m
}

func isSubsetMap(m, cmp map[int]bool) bool {
	if len(m) < len(cmp) {
		return false
	}
	for key, value := range cmp {
		mval, found := m[key]
		if !found || mval != value {
			return false
		}
	}
	return true

}

func intersection(m, cmp map[int]bool) bool {
	for key, value := range cmp {
		mval, found := m[key]
		if found || mval == value {
			return true
		}
	}
	return false
}



func Part1(fname string) int {
	scanner := util.FileScanner(fname, fsys)
	sum := 0
	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")
		map1 := rangeToMap(strings.Split(pairs[0], "-"))
		map2 := rangeToMap(strings.Split(pairs[1], "-"))
		if isSubsetMap(map1, map2) || isSubsetMap(map2, map1) {
			sum += 1
		}

	}
	return sum
}

func Part2(fname string) int {
	scanner := util.FileScanner(fname, fsys)
	sum := 0
	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")
		map1 := rangeToMap(strings.Split(pairs[0], "-"))
		map2 := rangeToMap(strings.Split(pairs[1], "-"))
		if intersection(map1, map2) || intersection(map2, map1) {
			sum += 1
		}

	}
	return sum
}

func main() {
	fmt.Printf("Part 1: %d\n", Part1("input"))
	fmt.Printf("Part 2: %d\n", Part2("input"))
}
