package main

import (
	"aoc/util"
	"embed"
	"fmt"
	"index/suffixarray"
	"unicode/utf8"
)

//go:embed input
//go:embed input_test
var fsys embed.FS

func substrings(str1, str2 string) []string {
	index := suffixarray.New([]byte(str1))
	var substrings []string
	for _, char := range str2 {
		offsets := index.Lookup([]byte(string(char)), -1)
		for _, offset := range offsets { // can this be done better?
			substrings = append(substrings, string(str1[offset]))
		}
	}
	return util.Unique(substrings)
}

func priority(char []byte) int {
	r, _ := utf8.DecodeRune(char)
	switch {
	case r >= 'A' && r <= 'Z':
		return int(r-'A'+1) + 26
	case r >= 'a' && r <= 'z':
		return int(r - 'a' + 1)
	}
	return 0
}

func Part1(fname string) int {
	scanner := util.FileScanner(fname, fsys)
	sum := 0
	for scanner.Scan() {
		items := scanner.Text()
		half := len(items) / 2
		item1, item2 := items[:half], items[half:]
		subs := substrings(item1, item2)
		for _, sub := range subs {
			sum += priority([]byte(sub))
		}
	}
	return sum
}

func Part2(fname string) int {
	scanner := util.FileScanner(fname, fsys)
	sum := 0
	var items []string
	for scanner.Scan() {
		items = append(items, scanner.Text())
		if len(items) == 3 {
			compMap := make(map[string]bool)
			for _, str := range substrings(items[0], items[1]) {
				compMap[str] = true
			}
			for _, str := range substrings(items[1], items[2]) {
				if compMap[str] {
					sum += priority([]byte(str))
				}
			}
			items = []string{}

		}
	}
	return sum
}

func main() {
	fmt.Printf("Part 1: %d\n", Part1("input"))
	fmt.Printf("Part 2: %d\n", Part2("input"))
}
