package main

// DOESN'T ENTIRELY WORK!!!!

import (
	"fmt"
	"strconv"
	"strings"
)

func toggle(inst instruction) command {

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
		default:
			panic("Unexpected instruction")
		}

		if len(fields) > 1 {
			inst.param1 = fields[1]
		}
		if len(fields) > 2 {
			inst.param2 = fields[2]
		}

		instructions = append(instructions, inst)
	}
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
	}
	return ""
}

func main() {
	var a, b, c, d int
	var p int
	// c = 1
	// a = 7

	ii := strings.Split(input, "\n")
	count := 0

	instructions := parse(ii)

	for p < len(instructions) {
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
		case cpy, jnz, tgl:
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
				val, _ = strconv.Atoi(inst.param1)
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
				val2, _ = strconv.Atoi(inst.param2)
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
				// fmt.Printf("jump %d to %d \n", val2, p+val2)
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
