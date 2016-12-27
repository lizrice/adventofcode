package main

import (
	"fmt"
	"strconv"
	"strings"
)

type password struct {
	content []byte
	length  int
}

func (p *password) String() string {
	return string(p.content)
}

func (p *password) move(words []string) {
	var X, Y int
	X, _ = strconv.Atoi(words[2])
	Y, _ = strconv.Atoi(words[5])

	if reverse {
		X, Y = Y, X
	}

	tmp := p.content[X]
	if X < Y {
		for i := X; i < Y; i++ {
			p.content[i] = p.content[i+1]
		}
	} else {
		for i := X; i > Y; i-- {
			p.content[i] = p.content[i-1]
		}
	}
	p.content[Y] = tmp
}

func (p *password) reverse(words []string) {
	var X, Y int
	X, _ = strconv.Atoi(words[2])
	Y, _ = strconv.Atoi(words[4])
	if X > Y {
		X, Y = Y, X
	}
	copy := make([]byte, Y-X+1)
	for i := 0; i < Y-X+1; i++ {
		copy[i] = p.content[X+i]
	}
	for x := X; x <= Y; x++ {
		p.content[x] = copy[len(copy)-(x-X)-1]
	}
}

func (p *password) rotate(words []string) {
	var F, X int
	switch words[1] {
	case "based":
		if reverse {
			// We rotated right X positions to get a letter to where it is now C from former position F
			// (F + X) % l = C => F = C - X
			// if F >= 4, X = F + 2 => F = X - 2 => C - X = X - 2 => 2X - 2 = C -> X = (C + 2)/2
			//C = (2F + 2) % l => F = (C-2)/2
			// else X = F + 1 => F = X - 1 => C - X = X - 1 => 2X - 1 = C -> X = (C + 1)/2
			//C = 2F + 1 => F = (C-1)/2
			C := strings.Index(string(p.content), words[6])
			// If C is even, F must be >= 4
			// if C%2 == 0 {
			// 	X = (C + 2) / 2
			// } else {
			// 	X = (C + 1) / 2
			// }

			switch C {
			case 0:
				X = 9
			case 1:
				X = 1
			case 2:
				X = 6
			case 3:
				X = 2
			case 4:
				X = 7
			case 5:
				X = 3
			case 6:
				X = 8
			case 7:
				X = 4
			}
			F = (C - X) % p.length

			fmt.Printf("Char currently at %d was at %d and rotated right by %d\n", C, F, X)

		} else {
			X = strings.Index(string(p.content), words[6])
			if X >= 4 {
				X++
			}
			X++
			// fmt.Printf("Based on index %d\n", X)
		}

	case "right", "left":
		X, _ = strconv.Atoi(words[2])
	}

	X = X % p.length
	// Rotating to the right by X is the same as rotating to the left by l - X
	if words[1] != "left" {
		X = p.length - X
	}

	// In reverse, we go in the opposite direction
	if reverse {
		X = p.length - X
	}

	// fmt.Printf("Rotating left %d\n", X)
	// Rotate left
	copy := make([]byte, X)
	for i := 0; i < X; i++ {
		copy[i] = p.content[i]
	}

	// fmt.Printf("Copy : %s\n", string(copy))
	for i := 0; i < p.length-X; i++ {
		// fmt.Printf("Copy rest from %d to %d\n", i+X, i)
		p.content[i] = p.content[i+X]
	}
	// fmt.Printf("Rest : %s\n", string(p.content))
	for i := 0; i < X; i++ {
		// fmt.Printf("index %d\n", p.length-X+1)
		p.content[p.length-X+i] = copy[i]
	}
}

func (p *password) swap(words []string) {
	var X, Y int
	switch words[1] {
	case "position":
		X, _ = strconv.Atoi(words[2])
		Y, _ = strconv.Atoi(words[5])
	case "letter":
		X = strings.Index(string(p.content), words[2])
		Y = strings.Index(string(p.content), words[5])
	}

	tmp := p.content[X]
	p.content[X] = p.content[Y]
	p.content[Y] = tmp
}

func (p *password) doLine(line string) {
	fmt.Println(line)
	words := strings.Fields(line)
	switch words[0] {
	case "rotate":
		p.rotate(words)
	case "move":
		p.move(words)
	case "reverse":
		p.reverse(words)
	case "swap":
		p.swap(words)
	}
	fmt.Printf(" %v\n", p)
}

func main() {
	p := &password{
		content: []byte("fbgdceah"),
		length:  len("fbgdceah"),
	}

	// adgfhcbe is wrong

	if reverse {
		fmt.Printf("Ended up with %v\n", p)
		lines := strings.Split(input, "\n")
		for l := len(lines); l > 0; l-- {
			p.doLine(lines[l-1])
		}

	} else {
		for _, line := range strings.Split(input, "\n") {
			p.doLine(line)
		}

	}
}

// var input = `swap position 4 with position 0 : edcba.
// swap letter d with letter b : edcba.
// reverse positions 0 through 4 : abcde.
// rotate left 1 step : bcdea.
// move position 1 to position 4 : bdeac.
// move position 3 to position 0 : abdec.
// rotate based on position of letter b : ecabd.
// rotate based on position of letter d : decab.`

var reverse = true

var input = `rotate based on position of letter d
move position 1 to position 6
swap position 3 with position 6
rotate based on position of letter c
swap position 0 with position 1
rotate right 5 steps
rotate left 3 steps
rotate based on position of letter b
swap position 0 with position 2
rotate based on position of letter g
rotate left 0 steps
reverse positions 0 through 3
rotate based on position of letter a
rotate based on position of letter h
rotate based on position of letter a
rotate based on position of letter g
rotate left 5 steps
move position 3 to position 7
rotate right 5 steps
rotate based on position of letter f
rotate right 7 steps
rotate based on position of letter a
rotate right 6 steps
rotate based on position of letter a
swap letter c with letter f
reverse positions 2 through 6
rotate left 1 step
reverse positions 3 through 5
rotate based on position of letter f
swap position 6 with position 5
swap letter h with letter e
move position 1 to position 3
swap letter c with letter h
reverse positions 4 through 7
swap letter f with letter h
rotate based on position of letter f
rotate based on position of letter g
reverse positions 3 through 4
rotate left 7 steps
swap letter h with letter a
rotate based on position of letter e
rotate based on position of letter f
rotate based on position of letter g
move position 5 to position 0
rotate based on position of letter c
reverse positions 3 through 6
rotate right 4 steps
move position 1 to position 2
reverse positions 3 through 6
swap letter g with letter a
rotate based on position of letter d
rotate based on position of letter a
swap position 0 with position 7
rotate left 7 steps
rotate right 2 steps
rotate right 6 steps
rotate based on position of letter b
rotate right 2 steps
swap position 7 with position 4
rotate left 4 steps
rotate left 3 steps
swap position 2 with position 7
move position 5 to position 4
rotate right 3 steps
rotate based on position of letter g
move position 1 to position 2
swap position 7 with position 0
move position 4 to position 6
move position 3 to position 0
rotate based on position of letter f
swap letter g with letter d
swap position 1 with position 5
reverse positions 0 through 2
swap position 7 with position 3
rotate based on position of letter g
swap letter c with letter a
rotate based on position of letter g
reverse positions 3 through 5
move position 6 to position 3
swap letter b with letter e
reverse positions 5 through 6
move position 6 to position 7
swap letter a with letter e
swap position 6 with position 2
move position 4 to position 5
rotate left 5 steps
swap letter a with letter d
swap letter e with letter g
swap position 3 with position 7
reverse positions 0 through 5
swap position 5 with position 7
swap position 1 with position 7
swap position 1 with position 7
rotate right 7 steps
swap letter f with letter a
reverse positions 0 through 7
rotate based on position of letter d
reverse positions 2 through 4
swap position 7 with position 1
swap letter a with letter h`
