package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// c = 1
	var startA = 0
	const finishLength int = 100

	var even, finished bool

	instructions := strings.Split(input, "\n")

	for !finished && startA < 1000 {
		var a, b, c, d int
		var p int
		startA++
		a = startA
		output := make([]byte, finishLength)
		outputCount := 0

		for p < len(instructions) && outputCount < finishLength {

			// fmt.Printf("p: %d i: %s\n", p, instructions[p])
			fields := strings.Fields(instructions[p])
			var val int
			// var val2 int
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
			case "cpy", "jnz", "out":
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

			// if fields[0] == "jnz" {
			// 	switch fields[2] {
			// 	case "a":
			// 		val2 = a
			// 	case "b":
			// 		val2 = b
			// 	case "c":
			// 		val2 = c
			// 	case "d":
			// 		val2 = d
			// 	default:
			// 		val2, _ = strconv.Atoi(fields[2])
			// 	}
			// 	// fmt.Printf("Jump val %d ", val2)
			// }

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
						fmt.Printf("Bad jump - skipping instruction: %v\n", err)
						p++
					}
					// fmt.Printf("jump %d to %d \n", jump, p+jump)
					p += jump
				} else {
					p++
				}
			case "out":
				if outputCount == 0 {
					even = (val % 2) == 0
				} else {
					// even tells us whether it was even last time around, so this time we want it to be different
					if (val%2 == 0) == even {
						fmt.Printf("No good after %d chars with A set to %d", outputCount, startA)
						outputCount = finishLength - 1
					} else {
						// fmt.Printf(".%d", outputCount)
						if outputCount >= finishLength-1 {
							finished = true
							fmt.Printf("\nDone! started with %d: %s\n", startA, output)
							os.Exit(1)
							break
						}
					}
					even = !even
				}
				output[outputCount] = strconv.Itoa(val)[0]
				outputCount++
				p++

			default:
				panic("Unexpected instruction")
			}
			// fmt.Printf("      a %d, b %d, c %d, d %d\n", a, b, c, d)

		}
		// fmt.Printf("      a %d, b %d, c %d, d %d\n", a, b, c, d)
		fmt.Printf("    %s\n", output)
	}
}

// var input = `cpy 2 a
// tgl a
// tgl a
// tgl a
// cpy 1 a
// dec a
// dec a`

var input = `cpy a d
cpy 4 c
cpy 643 b
inc d
dec b
jnz b -2
dec c
jnz c -5
cpy d a
jnz 0 0
cpy a b
cpy 0 a
cpy 2 c
jnz b 2
jnz 1 6
dec b
dec c
jnz c -4
inc a
jnz 1 -7
cpy 2 b
jnz c 2
jnz 1 4
dec b
dec c
jnz 1 -4
jnz 0 0
out b
jnz a -19
jnz 1 -21`
