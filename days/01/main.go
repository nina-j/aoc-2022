package main

import (
	util "aoc/util"
	"embed"
	"fmt"
	"sort"
	"strconv"
)

//go:embed input.txt
//go:embed input_test.txt
var fsys embed.FS

func calorieSums(fname string) []int {
	scanner := util.FileScanner(fname, fsys)
	calorie_sums := []int{0}
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			// This is probably not okay :'D
			calorie_sums = append(calorie_sums, 0)
		}
		calorie_sums[len(calorie_sums)-1] += value

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	sort.Slice(calorie_sums, func(i, j int) bool {
		return calorie_sums[i] > calorie_sums[j]
	})
	return calorie_sums
}

func Part1(fname string) int {
	return calorieSums(fname)[0]
}

func Part2(fname string) int {
	sum := 0
	for _, cals := range calorieSums(fname)[:3] {
		sum += cals
	}
	return sum
}

func main() {
	fmt.Printf("Part 1: %d\n", Part1("input.txt"))
	fmt.Printf("Part 2: %d\n", Part2("input.txt"))
}
