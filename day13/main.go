package main

import "fmt"

var input = 1362

// var input = 10

type pos struct {
	x int
	y int
}

var dest = pos{x: 31, y: 39}
var maxX int
var maxY int

type row []rune

// var dest = pos{x: 7, y: 4}
var previous = make(map[pos]bool)

func same(a, b pos) bool {
	return (a.x == b.x && a.y == b.y)
}

func bitCount(val int) (bits int) {
	for val > 0 {
		if val&1 == 1 {
			bits++
		}
		val = int(val / 2)
	}
	return bits
}

func isOpen(p pos) bool {
	val := p.x*p.x + 3*p.x + 2*p.x*p.y + p.y + p.y*p.y + input
	bits := bitCount(val)
	return (bits%2 == 0)
}

func isOK(p pos) bool {
	var visited bool
	if p.x < 0 || p.y < 0 {
		return false
	}

	if _, visited = previous[p]; visited {
		return false
	}

	previous[p] = isOpen(p)
	if p.x > maxX {
		maxX = p.x
	}
	if p.y > maxY {
		maxY = p.y
	}
	return previous[p]
}

func possibleStates(s []pos) (possible []pos) {
	for _, p := range s {
		north := pos{x: p.x, y: p.y - 1}
		south := pos{x: p.x, y: p.y + 1}
		east := pos{x: p.x + 1, y: p.y}
		west := pos{x: p.x - 1, y: p.y}
		for _, d := range []pos{north, south, east, west} {
			if isOK(d) {
				possible = append(possible, d)
			}
		}
	}
	return possible
}

func locationsVisited() (locations int) {
	for _, v := range previous {
		if v {
			locations++
		}
	}
	return locations
}

func main() {
	// fmt.Printf("Bitcount of 7: %d\n", bitCount(7))
	// fmt.Printf("Bitcount of 16: %d\n", bitCount(16))

	// Starting at 0,0 we'll do a breadth-first search of possible routes
	start := pos{x: 1, y: 1}
	states := []pos{start}
	previous[start] = true // start point is ok

	var moves = 0
	// var locations = 1

	for moves < 100 {
		moves++
		newStates := possibleStates(states)
		for _, n := range newStates {
			if same(n, dest) {
				fmt.Printf("Finished in %d moves\n", moves)
				moves = 1000
				break
			}
		}
		states = newStates
		fmt.Printf("After step %d - possible states %d, locations so far %d\n", moves, len(states), locationsVisited())
	}

	// fmt.Printf("Total locations %d\n", locations)

	fmt.Printf("maxX %d, maxY %d", maxX, maxY)
	grid := make([]row, maxY+1)
	for i := 0; i < maxY+1; i++ {
		grid[i] = make([]rune, maxX+1)
		for j := 0; j < maxX+1; j++ {
			grid[i][j] = ' '
		}
	}

	for k, v := range previous {
		// fmt.Printf("k: %#v, v: %v\n", k, v)
		if v {
			grid[k.y][k.x] = '.'
		} else {
			grid[k.y][k.x] = '#'
		}
	}

	for i, y := range grid {
		fmt.Printf("%d ", i)
		for _, x := range y {
			fmt.Printf("%c", x)
		}
		fmt.Printf("\n")
	}
}
