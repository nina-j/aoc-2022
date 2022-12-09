package main

import (
	"aoc/util"
	"bufio"
	"embed"
	"fmt"
	"path"
	"sort"
	"strconv"
	"strings"
)

//go:embed input
//go:embed input_test
var fsys embed.FS

func getDirSizes(scanner *bufio.Scanner) map[string]int {
	curPath := ""
	files := []string{}
	dirSizes := make(map[string]int)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "$ ")
		cmd := strings.Split(line, " ")
		switch cmd[0] {
		case "cd":
			curPath = path.Join(curPath, cmd[1])
			dirSizes[curPath] = 0
		case "ls":
			continue
		case "dir":
			continue
		default: // a file size
			files = append(files, path.Join(curPath, cmd[0]))
		}

	}
	for k := range dirSizes {
		for _, f := range files {
			fsize, err := strconv.Atoi(path.Base(f))
			if err != nil {
				panic(err)
			}
			if strings.Contains(f, k) {
				dirSizes[k] += fsize
			}
		}
	}
	return dirSizes
}

func Part1(fname string) int {
	scanner := util.FileScanner(fname, fsys)
	res := 0
	dirSizes := getDirSizes(scanner)
	for _, v := range dirSizes {
		if v <= 100000 {
			res += v
		}
	}
	return res
}

func Part2(fname string) int {
	scanner := util.FileScanner(fname, fsys)
	dirSizes := getDirSizes(scanner)
	totalDisk := 70000000
	needSpace := 30000000
	curSpace := totalDisk - dirSizes["/"]
	valid := []int{}
	for _, v := range dirSizes {
		if curSpace+v >= needSpace {
			valid = append(valid, v)
		}
	}
	sort.Slice(valid, func(i, j int) bool { return valid[i] < valid[j] })
	return valid[0]
}

func main() {
	fmt.Printf("Part 1: %v\n", Part1("input"))
	fmt.Printf("Part 2: %v\n", Part2("input"))
}
