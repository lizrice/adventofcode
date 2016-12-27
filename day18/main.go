package main

import (
	"fmt"
	"strings"
)

func tile(l, c, r byte) byte {
	if l == '^' && c == '^' && r == '.' {
		return '^'
	}
	if l == '.' && c == '^' && r == '^' {
		return '^'
	}
	if l == '^' && c == '.' && r == '.' {
		return '^'
	}
	if l == '.' && c == '.' && r == '^' {
		return '^'
	}
	return '.'
}

func nextRow(row []byte) (newRow []byte, safe int) {
	newRow = make([]byte, len(row))
	var left, right byte

	for b := 0; b < len(row); b++ {
		if b == 0 {
			left = '.'
		} else {
			left = row[b-1]
		}
		if b == len(row)-1 {
			right = '.'
		} else {
			right = row[b+1]
		}
		newRow[b] = tile(left, row[b], right)
		if newRow[b] == '.' {
			safe++
		}
	}
	return
}

func main() {
	row := []byte(".^..^....^....^^.^^.^.^^.^.....^.^..^...^^^^^^.^^^^.^.^^^^^^^.^^^^^..^.^^^.^^..^.^^.^....^.^...^^.^.")
	maxRows := 400000
	safeTiles := 0
	safe := strings.Count(string(row), ".")

	for r := 0; r < maxRows; r++ {
		safeTiles += safe
		// fmt.Println(string(row))
		row, safe = nextRow(row)
	}

	fmt.Println(safeTiles)
}
