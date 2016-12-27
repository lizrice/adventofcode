package main

import (
	"fmt"
	"strconv"
	"strings"
)

const width = 50
const height = 6

type row [width]bool
type col [height]bool
type screen [height]row

func (s *screen) print() {
	for _, r := range s {
		for _, c := range r {
			if c {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (s *screen) rect(A, B int) {
	for r := 0; r < B; r++ {
		for c := 0; c < A; c++ {
			s[r][c] = true
		}
	}
}

func (s *screen) rotateRow(A, B int) {
	var newRow row
	r := s[A]
	for c := 0; c < width; c++ {
		newRow[(c+B)%width] = r[c]
	}

	s[A] = newRow
}

func (s *screen) rotateCol(A, B int) {
	var newCol col
	for r := 0; r < height; r++ {
		newCol[(r+B)%height] = s[r][A]
	}
	for r := 0; r < height; r++ {
		s[r][A] = newCol[r]
	}
}

func (s *screen) voltage() int {
	var v int
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if s[r][c] {
				v++
			}
		}
	}
	return v
}

func main() {
	var ss screen
	s := &ss
	// s.rect(3, 2)
	// s.print()
	// s.rotateCol(1, 1)
	// s.print()
	// s.rotateRow(0, 4)
	// s.print()

	for _, instruction := range strings.Split(input, "\n") {
		ii := strings.Fields(instruction)
		switch ii[0] {
		case "rotate":
			B, _ := strconv.Atoi(ii[4])
			switch ii[1] {
			case "row":
				A, _ := strconv.Atoi(strings.TrimPrefix(ii[2], "y="))
				s.rotateRow(A, B)
			case "column":
				A, _ := strconv.Atoi(strings.TrimPrefix(ii[2], "x="))
				s.rotateCol(A, B)
			}

		case "rect":
			params := strings.Split(ii[1], "x")
			fmt.Printf("Rect %s x %s", params[0], params[1])
			A, _ := strconv.Atoi(params[0])
			B, _ := strconv.Atoi(params[1])
			s.rect(A, B)
		default:
			panic(ii[0])
		}
	}

	fmt.Printf("Voltage is %d\n", s.voltage())
	s.print()

}

var input = `rect 1x1
rotate row y=0 by 6
rect 1x1
rotate row y=0 by 3
rect 1x1
rotate row y=0 by 5
rect 1x1
rotate row y=0 by 4
rect 2x1
rotate row y=0 by 5
rect 2x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 5
rect 4x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 3
rect 1x1
rotate row y=0 by 3
rect 1x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 6
rect 4x1
rotate row y=0 by 4
rotate column x=0 by 1
rect 3x1
rotate row y=0 by 6
rotate column x=0 by 1
rect 4x1
rotate column x=10 by 1
rotate row y=2 by 16
rotate row y=0 by 8
rotate column x=5 by 1
rotate column x=0 by 1
rect 7x1
rotate column x=37 by 1
rotate column x=21 by 2
rotate column x=15 by 1
rotate column x=11 by 2
rotate row y=2 by 39
rotate row y=0 by 36
rotate column x=33 by 2
rotate column x=32 by 1
rotate column x=28 by 2
rotate column x=27 by 1
rotate column x=25 by 1
rotate column x=22 by 1
rotate column x=21 by 2
rotate column x=20 by 3
rotate column x=18 by 1
rotate column x=15 by 2
rotate column x=12 by 1
rotate column x=10 by 1
rotate column x=6 by 2
rotate column x=5 by 1
rotate column x=2 by 1
rotate column x=0 by 1
rect 35x1
rotate column x=45 by 1
rotate row y=1 by 28
rotate column x=38 by 2
rotate column x=33 by 1
rotate column x=28 by 1
rotate column x=23 by 1
rotate column x=18 by 1
rotate column x=13 by 2
rotate column x=8 by 1
rotate column x=3 by 1
rotate row y=3 by 2
rotate row y=2 by 2
rotate row y=1 by 5
rotate row y=0 by 1
rect 1x5
rotate column x=43 by 1
rotate column x=31 by 1
rotate row y=4 by 35
rotate row y=3 by 20
rotate row y=1 by 27
rotate row y=0 by 20
rotate column x=17 by 1
rotate column x=15 by 1
rotate column x=12 by 1
rotate column x=11 by 2
rotate column x=10 by 1
rotate column x=8 by 1
rotate column x=7 by 1
rotate column x=5 by 1
rotate column x=3 by 2
rotate column x=2 by 1
rotate column x=0 by 1
rect 19x1
rotate column x=20 by 3
rotate column x=14 by 1
rotate column x=9 by 1
rotate row y=4 by 15
rotate row y=3 by 13
rotate row y=2 by 15
rotate row y=1 by 18
rotate row y=0 by 15
rotate column x=13 by 1
rotate column x=12 by 1
rotate column x=11 by 3
rotate column x=10 by 1
rotate column x=8 by 1
rotate column x=7 by 1
rotate column x=6 by 1
rotate column x=5 by 1
rotate column x=3 by 2
rotate column x=2 by 1
rotate column x=1 by 1
rotate column x=0 by 1
rect 14x1
rotate row y=3 by 47
rotate column x=19 by 3
rotate column x=9 by 3
rotate column x=4 by 3
rotate row y=5 by 5
rotate row y=4 by 5
rotate row y=3 by 8
rotate row y=1 by 5
rotate column x=3 by 2
rotate column x=2 by 3
rotate column x=1 by 2
rotate column x=0 by 2
rect 4x2
rotate column x=35 by 5
rotate column x=20 by 3
rotate column x=10 by 5
rotate column x=3 by 2
rotate row y=5 by 20
rotate row y=3 by 30
rotate row y=2 by 45
rotate row y=1 by 30
rotate column x=48 by 5
rotate column x=47 by 5
rotate column x=46 by 3
rotate column x=45 by 4
rotate column x=43 by 5
rotate column x=42 by 5
rotate column x=41 by 5
rotate column x=38 by 1
rotate column x=37 by 5
rotate column x=36 by 5
rotate column x=35 by 1
rotate column x=33 by 1
rotate column x=32 by 5
rotate column x=31 by 5
rotate column x=28 by 5
rotate column x=27 by 5
rotate column x=26 by 5
rotate column x=17 by 5
rotate column x=16 by 5
rotate column x=15 by 4
rotate column x=13 by 1
rotate column x=12 by 5
rotate column x=11 by 5
rotate column x=10 by 1
rotate column x=8 by 1
rotate column x=2 by 5
rotate column x=1 by 5`
