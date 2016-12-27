package main

// DOESN'T ENTIRELY WORK!!!!

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toggle(inst instruction) command {
	panic("Unexpected toggle")

	switch inst.cmd {
	case inc:
		return dec
	case dec:
		return inc
	case tgl:
		return inc
	case jnz:
		return cpy
	case cpy:
		return jnz
	default:
		fmt.Printf("Unexpected instruction %d\n", inst.cmd)
	}

	// fmt.Printf("Changed instruction: %s\n", *instruction)
	return tgl
}

type command int

const (
	inc command = iota
	dec
	tgl
	jnz
	cpy
	out
)

type register int

const (
	A register = iota
	B
	C
	D
)

type instruction struct {
	cmd    command
	reg    register
	param1 string
	param2 string
}

func parse(ii []string) (instructions []instruction) {
	for _, line := range ii {
		inst := instruction{}

		// fmt.Printf("p: %d i: %s\n", p, instructions[p])
		fields := strings.Fields(line)

		// Which field indicates a register whose value will change
		switch fields[0] {
		case "inc":
			inst.cmd = inc
		case "dec":
			inst.cmd = dec
		case "tgl":
			inst.cmd = tgl
		case "jnz":
			inst.cmd = jnz
		case "cpy":
			inst.cmd = cpy
		case "out":
			inst.cmd = out
		default:
			panic("Unexpected instruction")
		}

		if len(fields) > 1 {
			inst.param1 = fields[1]
		}
		if len(fields) > 2 {
			inst.param2 = fields[2]
		}

		fmt.Printf("%v %s %s\n", inst.cmd, inst.param1, inst.param2)
		instructions = append(instructions, inst)
	}
	os.Exit(1)
	return instructions

}

func (i instruction) String() string {
	switch i.cmd {
	case inc:
		return "inc"
	case dec:
		return "dec"
	case tgl:
		return "tgl"
	case jnz:
		return "jnz"
	case cpy:
		return "cpy"
	case out:
		return "out"
	}
	return ""
}

func main() {
	// c = 1
	// a = 7
	var even, finished bool
	var startA = 0
	const finishLength int = 10
	var err error

	ii := strings.Split(input, "\n")
	count := 0

	instructions := parse(ii)

	for !finished && startA < 1 {
		startA++
		var a, b, c, d int
		var p int
		a = startA

		output := make([]byte, finishLength)
		outputCount := 0
		for p < len(instructions) && outputCount < finishLength {
			// fmt.Printf("p: %d i: %s\n", p, instructions[p])
			inst := instructions[p]
			var val, val2 int
			var reg *int
			var regField string

			// fmt.Printf("%v\n", inst)

			// Which field indicates a register whose value will change
			switch inst.cmd {
			case inc, dec:
				regField = inst.param1
			case cpy:
				regField = inst.param2
			}

			// Which fields
			switch inst.cmd {
			case inc, dec, cpy:
				switch regField {
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
			switch inst.cmd {
			case cpy, jnz, tgl, out:
				switch inst.param1 {
				case "a":
					val = a
				case "b":
					val = b
				case "c":
					val = c
				case "d":
					val = d
				default:
					val, err = strconv.Atoi(inst.param1)
					if err != nil {
						panic("Bad val")
					}
				}
			}

			if inst.cmd == jnz {
				switch inst.param2 {
				case "a":
					val2 = a
				case "b":
					val2 = b
				case "c":
					val2 = c
				case "d":
					val2 = d
				default:
					val2, err = strconv.Atoi(inst.param2)
					if err != nil {
						panic("Bad val")
					}
				}
				// fmt.Printf("Jump val %d ", val2)
			}

			switch inst.cmd {
			case cpy:
				*reg = val
				p++
			case inc:
				*reg = *reg + 1
				p++
			case dec:
				*reg = *reg - 1
				p++
			case jnz:
				if val != 0 {
					// jump, err := strconv.Atoi(fields[2])
					// if err != nil {
					// 	fmt.Printf("Bad jump - skipping instruction: %v\n", err)
					// 	p++
					// }
					fmt.Printf("jump %d to %d \n", val2, p+val2)
					p += val2
				} else {
					p++
				}
			case tgl:
				if p+val >= 0 && p+val < len(instructions) {
					instructions[p+val].cmd = toggle(instructions[p+val])
					// fmt.Printf("Changes instruction at %d to %v", p+val, instructions[p+val])
				}
				p++
			case out:
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
			count++
			// if count%1000000 == 0 {
			// 	fmt.Println(count / 1000000)
			// }

		}
		fmt.Printf("      a %d, b %d, c %d, d %d", a, b, c, d)
		fmt.Printf("    %s\n", output)
		// fmt.Printf("    finished %t, outputCount %d\n", finished, outputCount)

	}

}

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
