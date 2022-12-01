package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"sort"
	"strconv"
)

//go:embed input.txt
//go:embed input_test.txt
var fsys embed.FS

func fileLines(fname string) (lines []string) {
	content, err := fsys.ReadFile(fname)
	if err != nil {
		panic(err) // Not great, better error handling!
	}
	scanner := bufio.NewScanner(bytes.NewReader(content))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func calorieSums(fname string) (calorie_sums []int) {
	lines := fileLines(fname)
	calorie_sums = append(calorie_sums, 0)
	for _, line := range lines {
		text := line
		value, err := strconv.Atoi(text)
		if err != nil {
			// This is probably not okay :'D
			calorie_sums = append(calorie_sums, 0)
		}
		calorie_sums[len(calorie_sums)-1] += value

	}
	sort.Slice(calorie_sums, func(i, j int) bool {
		return calorie_sums[i] > calorie_sums[j]
	})
	return calorie_sums
}

func Part1(fname string) (calories int) {
	return calorieSums(fname)[0]
}

func Part2(fname string) (calories int) {
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
