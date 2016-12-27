package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a, b, c, d int
	var p int
	c = 1

	instructions := strings.Split(input, "\n")

	for p < len(instructions) {
		// fmt.Printf("p: %d i: %s\n", p, instructions[p])
		fields := strings.Fields(instructions[p])
		var val int
		var reg *int
		var regField int

		// Which field indicates a register whose value will change
		switch fields[0] {
		case "inc", "dec":
			regField = 1
		case "cpy":
			regField = 2
		}

		// Which fields
		switch fields[0] {
		case "inc", "dec", "cpy":
			switch fields[regField] {
			case "a":
				reg = &a
			case "b":
				reg = &b
			case "c":
				reg = &c
			case "d":
				reg = &d
			default:
				panic("Unexpected register")
			}
		}

		// Which fields indicates a value
		switch fields[0] {
		case "cpy", "jnz":
			switch fields[1] {
			case "a":
				val = a
			case "b":
				val = b
			case "c":
				val = c
			case "d":
				val = d
			default:
				val, _ = strconv.Atoi(fields[1])
			}
		}

		switch fields[0] {
		case "cpy":
			*reg = val
			p++
		case "inc":
			*reg = *reg + 1
			p++
		case "dec":
			*reg = *reg - 1
			p++
		case "jnz":
			if val != 0 {
				jump, err := strconv.Atoi(fields[2])
				if err != nil {
					panic("Bad jump")
				}
				p += jump
			} else {
				p++
			}
		default:
			panic("Unexpected instruction")
		}

	}
	fmt.Printf("      a %d, b %d, c %d, d %d\n", a, b, c, d)
}

var input = `cpy 1 a
cpy 1 b
cpy 26 d
jnz c 2
jnz 1 5
cpy 7 c
inc d
dec c
jnz c -2
cpy a c
inc a
dec b
jnz b -2
cpy c b
dec d
jnz d -6
cpy 16 c
cpy 17 d
inc a
dec d
jnz d -2
dec c
jnz c -5`
