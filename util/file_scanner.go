package util

import (
	"bufio"
	"embed"
)

func FileScanner(fname string, fsys embed.FS) *bufio.Scanner {
	file, err := fsys.Open(fname)
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file)
}
