package main

import (
	"aoc/util"
	"embed"
	"fmt"
	"strings"
)

//go:embed input
//go:embed input_test
var fsys embed.FS

func detectMarker(buffer []string, windowSize int) int {
	for i := 0; i < len(buffer)-windowSize-1; i++ {
		end := i + windowSize
		window := buffer[i:end]
		if len(window) == len(util.Unique(window)) {
			return end
		}
	}
	return -1
}

func Part1(fname string) int {
	scanner := util.FileScanner(fname, fsys)
	res := 0
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), "")
		res = detectMarker(buf, 4)
	}
	return res
}
func Part2(fname string) int {
	scanner := util.FileScanner(fname, fsys)
	res := 0
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), "")
		res = detectMarker(buf, 14)
	}
	return res
}

func main() {
	fmt.Printf("Part 1: %v\n", Part1("input"))
	fmt.Printf("Part 2: %v\n", Part2("input"))
}
