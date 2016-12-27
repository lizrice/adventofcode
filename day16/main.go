package main

import "fmt"

const (
	desiredLength = 35651584
	// desiredLength = 272
	input = "10111100110001111"
	// desiredLength = 20
	// input         = "10000"
)

type output struct {
	length  int
	content [desiredLength]byte
}

func (b *output) dragon() {
	b.content[b.length] = '0'

	newLength := 2*(b.length) + 1
	if newLength > desiredLength {
		newLength = desiredLength
	}

	j := b.length + 1
	i := b.length - 1
	for j < newLength {
		if b.content[i] == '0' {
			b.content[j] = '1'
		} else {
			b.content[j] = '0'
		}
		j++
		i--
	}

	b.length = newLength
}

func (b *output) checksum() {
	var even = true

	for even == true {
		j := 0
		for i := 0; i < b.length-1; i += 2 {
			if b.content[i] == b.content[i+1] {
				b.content[j] = '1'
			} else {
				b.content[j] = '0'
			}
			j++
		}
		b.length = j
		even = (j%2 == 0)
	}
	return
}

func (b *output) print() {
	fmt.Printf("Length %d content ", b.length)
	for i := 0; i < b.length; i++ {
		fmt.Printf("%c", b.content[i])
	}
	fmt.Printf("\n")
}

func (b *output) set(input string) {
	b.length = len(input)
	for i := 0; i < b.length; i++ {
		b.content[i] = input[i]
	}
}

func main() {
	// fmt.Println(checksum("110010110100"))
	// fmt.Println(dragon("111100001010"))
	// os.Exit(1)

	var o output
	o.set(input)

	target := 100000 // For printing out when length gets this long
	for o.length < desiredLength {
		o.dragon()
		if o.length > target {
			fmt.Println(target)
			target += 100000
		}
	}

	fmt.Println("--------------")
	o.checksum()
	o.print()
}
