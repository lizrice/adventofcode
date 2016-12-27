package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// facing: 0-N 1-E, 2-S, 3-W
// x is E
// y is N
type position struct {
	x      int
	y      int
	facing int
}

const (
	maxSize = 1000
	origin  = 500
)

func takeStep(p *position, step string) bool {
	// Turn right or left
	if string(step[0]) == "R" {
		fmt.Printf("  turn right")
		p.facing++
		if p.facing == 4 {
			p.facing = 0
		}
	} else {
		fmt.Printf("  turn left")
		p.facing--
		if p.facing == -1 {
			p.facing = 3
		}
	}

	d, err := strconv.Atoi(step[1:])
	if err != nil {
		panic(err)
	}

	fmt.Printf("  distance %d\n", d)
	for d > 0 {
		switch p.facing {
		case 0:
			p.y++
		case 1:
			p.x++
		case 2:
			p.y--
		case 3:
			p.x--
		default:
			fmt.Printf("Direction %d", p.facing)
			panic("unexpected direction")
		}
		d--

		if visited(p) {
			return true
		}
	}

	return false
}

func distance(p position) int {
	return int(math.Abs(float64(p.x-origin))) + int(math.Abs(float64(p.y-origin)))
}

func visited(p *position) bool {
	index := p.y*maxSize + p.x
	if grid[index] {
		fmt.Printf("Pos index %d already visited!\n", index)
		return true
	}

	grid[index] = true
	return false
}

var grid map[int]bool

func main() {
	stepsString := "R3, L5, R2, L1, L2, R5, L2, R2, L2, L2, L1, R2, L2, R4, R4, R1, L2, L3, R3, L1, R2, L2, L4, R4, R5, L3, R3, L3, L3, R4, R5, L3, R3, L5, L1, L2, R2, L1, R3, R1, L1, R187, L1, R2, R47, L5, L1, L2, R4, R3, L3, R3, R4, R1, R3, L1, L4, L1, R2, L1, R4, R5, L1, R77, L5, L4, R3, L2, R4, R5, R5, L2, L2, R2, R5, L2, R194, R5, L2, R4, L5, L4, L2, R5, L3, L2, L5, R5, R2, L3, R3, R1, L4, R2, L1, R5, L1, R5, L1, L1, R3, L1, R5, R2, R5, R5, L4, L5, L5, L5, R3, L2, L5, L4, R3, R1, R1, R4, L2, L4, R5, R5, R4, L2, L2, R5, R5, L5, L2, R4, R4, L4, R1, L3, R1, L1, L1, L1, L4, R5, R4, L4, L4, R5, R3, L2, L2, R3, R1, R4, L3, R1, L4, R3, L3, L2, R2, R2, R2, L1, L4, R3, R2, R2, L3, R2, L3, L2, R4, L2, R3, L4, R5, R4, R1, R5, R3"
	steps := strings.Split(stepsString, ",")

	pos := position{
		x: origin, y: origin, facing: 0,
	}

	grid = make(map[int]bool, origin*origin)

	for _, step := range steps {
		fmt.Printf("  %#v\n", pos)
		if takeStep(&pos, strings.TrimSpace(step)) {
			break
		}
	}

	fmt.Printf("Position is %#v\n", pos)
	fmt.Printf("Distance is %#v\n", distance(pos))
}
