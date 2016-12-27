package main

import (
	"fmt"
	"strconv"
	"strings"
)

func toggle(instruction *string) {
	fields := strings.Fields(*instruction)

	switch fields[0] {
	case "inc":
		*instruction = strings.Replace(*instruction, "inc", "dec", 1)
	case "dec":
		*instruction = strings.Replace(*instruction, "dec", "inc", 1)
	case "tgl":
		*instruction = strings.Replace(*instruction, "tgl", "inc", 1)
	case "jnz":
		*instruction = strings.Replace(*instruction, "jnz", "cpy", 1)
	case "cpy":
		*instruction = strings.Replace(*instruction, "cpy", "jnz", 1)
	default:
		fmt.Printf("Unexpected instruction %s\n", fields[0])
	}

	// fmt.Printf("Changed instruction: %s\n", *instruction)
}

func main() {
	var a, b, c, d int
	var p int
	// c = 1
	a = 12

	instructions := strings.Split(input, "\n")
	count := 0

	for p < len(instructions) {
		// fmt.Printf("p: %d i: %s\n", p, instructions[p])
		fields := strings.Fields(instructions[p])
		var val, val2 int
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
				fmt.Println("Unexpected register - skipping instruction")
				continue
			}
		}

		// Which fields indicates a value
		switch fields[0] {
		case "cpy", "jnz", "tgl":
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

		if fields[0] == "jnz" {
			switch fields[2] {
			case "a":
				val2 = a
			case "b":
				val2 = b
			case "c":
				val2 = c
			case "d":
				val2 = d
			default:
				val2, _ = strconv.Atoi(fields[2])
			}
			// fmt.Printf("Jump val %d ", val2)
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
				// jump, err := strconv.Atoi(fields[2])
				// if err != nil {
				// 	fmt.Printf("Bad jump - skipping instruction: %v\n", err)
				// 	p++
				// }
				// fmt.Printf("jump %d to %d \n", val2, p+val2)
				p += val2
			} else {
				p++
			}
		case "tgl":
			if p+val >= 0 && p+val < len(instructions) {
				toggle(&instructions[p+val])
			}
			p++
		default:
			panic("Unexpected instruction")
		}
		// fmt.Printf("      a %d, b %d, c %d, d %d\n", a, b, c, d)
		count++
		if count%1000000 == 0 {
			fmt.Println(count / 1000000)
		}

	}
	fmt.Printf("      a %d, b %d, c %d, d %d\n", a, b, c, d)
}

// var input = `cpy 2 a
// tgl a
// tgl a
// tgl a
// cpy 1 a
// dec a
// dec a`

var input = `cpy a b
dec b
cpy a d
cpy 0 a
cpy b c
inc a
dec c
jnz c -2
dec d
jnz d -5
dec b
cpy b c
cpy c d
dec d
inc c
jnz d -2
tgl c
cpy -16 c
jnz 1 c
cpy 76 c
jnz 84 d
inc a
inc d
jnz d -2
inc c
jnz c -5`
