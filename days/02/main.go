package main

import (
	util "aoc/util"
	"embed"
	"fmt"
	"strings"
)

//go:embed input
//go:embed input_test
var fsys embed.FS
var combinations = map[string]map[string]int{
	"A": { // Rock
		"A": 3, // Rock
		"B": 0, // Paper
		"C": 6, // Scissors
		"p": 1, // choice points
	},
	"B": { // Paper
		"A": 6, // Rock
		"B": 3, // Paper
		"C": 0, // Scissors
		"p": 2, // choice points
	},
	"C": { // Scissors
		"A": 0, // Rock
		"B": 6, // Paper
		"C": 3, // Scissors
		"p": 3, // choice points
	}}

func getStratMap(key string) map[string]int {
	switch key {
	case "X": // rock
		return combinations["A"]
	case "Y": // paper
		return combinations["B"]
	case "Z": // scissors
		return combinations["C"]
	default: // return key
		return combinations[key]
	}
}

func Part1(fname string) int {
	scanner := util.FileScanner(fname, fsys)
	points := 0
	for scanner.Scan() {
		strat := strings.Split(scanner.Text(), " ")
		stratMap := getStratMap(strat[1])
		points += stratMap[strat[0]] + stratMap["p"]

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return points
}

func getPoints(points int, mapping map[string]int) int {
	result := 0
	for key, value := range mapping {
		if value == points && key != "p" {
			result = getStratMap(key)["p"]

		}
	}
	return result
}

func Part2(fname string) int {
	scanner := util.FileScanner(fname, fsys)
	sum := 0
	for scanner.Scan() {
		strat := strings.Split(scanner.Text(), " ")
		stratMap := getStratMap(strat[0])
		switch strat[1] {
		case "X": // lose
			// they win (6)
			sum += getPoints(6, stratMap) + 0
		case "Y": //draw
			// draw (3)
			sum += getPoints(3, stratMap) + 3
		case "Z": // win
			// they lose (0)
			sum += getPoints(0, stratMap) + 6
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return sum
}

func main() {
	fmt.Printf("Part 1: %d\n", Part1("input"))
	fmt.Printf("Part 2: %d\n", Part2("input"))
}
