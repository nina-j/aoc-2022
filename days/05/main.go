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

func moveOne(moves, from, to int, m map[int][]string) map[int][]string {
	for i := 0; i < moves; i++ {
		var crate string
		last := len(m[from]) - 1
		crate, m[from] = m[from][last], m[from][:last]
		m[to] = append(m[to], crate)
	}
	return m
}

func moveMany(moves, from, to int, m map[int][]string) map[int][]string {
	var crate []string
	last := len(m[from]) - moves
	crate, m[from] = m[from][last:], m[from][:last]
	m[to] = append(m[to], crate...)
	return m
}

func moveCrates(procedure []string, m map[int][]string, move func(moves, from, to int, m map[int][]string) map[int][]string) map[int][]string {
	var moves, from, to int
	var err error
	moves, err = strconv.Atoi(procedure[1])
	if err != nil {
		panic(err)
	}
	from, err = strconv.Atoi(procedure[3])
	if err != nil {
		panic(err)
	}
	to, err = strconv.Atoi(procedure[5])
	if err != nil {
		panic(err)
	}

	return move(moves, from, to, m)
}

func Part1(fname string, crates map[int][]string) string {
	scanner := util.FileScanner(fname, fsys)
	res := ""
	for scanner.Scan() {
		proc := strings.Split(scanner.Text(), " ")
		crates = moveCrates(proc, crates, moveOne)

	}
	for i := 1; i <= len(crates); i++ {
		last := len(crates[i]) - 1
		res += crates[i][last]
	}
	return res
}

func Part2(fname string, crates map[int][]string) string {
	scanner := util.FileScanner(fname, fsys)
	res := ""
	for scanner.Scan() {
		proc := strings.Split(scanner.Text(), " ")
		crates = moveCrates(proc, crates, moveMany)

	}
	for i := 1; i <= len(crates); i++ {
		last := len(crates[i]) - 1
		res += crates[i][last]
	}
	return res
}

//crates
//            [Q]     [G]     [M]
//            [B] [S] [V]     [P] [R]
//    [T]     [C] [F] [L]     [V] [N]
//[Q] [P]     [H] [N] [S]     [W] [C]
//[F] [G] [B] [J] [B] [N]     [Z] [L]
//[L] [Q] [Q] [Z] [M] [Q] [F] [G] [D]
//[S] [Z] [M] [G] [H] [C] [C] [H] [Z]
//[R] [N] [S] [T] [P] [P] [W] [Q] [G]
// 1   2   3   4   5   6   7   8   9

func main() {
	crates := map[int][]string{
		1: {"R", "S", "L", "F", "Q"},
		2: {"N", "Z", "Q", "G", "P", "T"},
		3: {"S", "M", "Q", "B"},
		4: {"T", "G", "Z", "J", "H", "C", "B", "Q"},
		5: {"P", "H", "M", "B", "N", "F", "S"},
		6: {"P", "C", "Q", "N", "S", "L", "V", "G"},
		7: {"W", "C", "F"},
		8: {"Q", "H", "G", "Z", "W", "V", "P", "M"},
		9: {"G", "Z", "D", "L", "C", "N", "R"},
	}
	fmt.Printf("Part 1: %v\n", Part1("input", crates))
	// whoops, crates are modified in place.
	crates = map[int][]string{
		1: {"R", "S", "L", "F", "Q"},
		2: {"N", "Z", "Q", "G", "P", "T"},
		3: {"S", "M", "Q", "B"},
		4: {"T", "G", "Z", "J", "H", "C", "B", "Q"},
		5: {"P", "H", "M", "B", "N", "F", "S"},
		6: {"P", "C", "Q", "N", "S", "L", "V", "G"},
		7: {"W", "C", "F"},
		8: {"Q", "H", "G", "Z", "W", "V", "P", "M"},
		9: {"G", "Z", "D", "L", "C", "N", "R"},
	}
	fmt.Printf("Part 2: %v\n", Part2("input", crates))
}
