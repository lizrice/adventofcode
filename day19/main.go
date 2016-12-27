package main

import "fmt"

type elves []elf

type elf struct {
	elfToLeft int
	presents  int
	index     int
}

func (e elves) init() {
	for i := 0; i < len(e); i++ {
		e[i].index = i + 1 // In the problem elves are numbered starting with 1
		e[i].elfToLeft = i + 1
		e[i].presents = 1
	}

	// They're sitting in a circle
	e[len(e)-1].elfToLeft = 0
}

func (e elves) circles(firstElf int, size int) (last int) {
	i := firstElf
	limit := int(float64(size) / 4)
	if limit == 0 {
		limit++
	}
	// fmt.Printf("Circle size %d, we'll remove %d\n", size, limit)
	removed := 0
	// half := int(float64(size) / 2)
	stealFrom := (i + size/2) % size
	fmt.Printf("Starting at elf %d, stealing from elf %d ", e[i].index, e[stealFrom].index)
	even := (size%2 == 0)
	for removed < limit {
		// stealFrom := (i + (size-removed)/2) % size
		// fmt.Printf("Size %d, removed %d: Elf %d steals from %d\n", size, removed, e[i].index, e[stealFrom].index)
		e[i].presents += e[stealFrom].presents
		e[stealFrom].presents = 0

		// Next elf to the left
		i = e[i].elfToLeft
		if i >= size {
			panic("nope!")
		}
		removed++
		if !even {
			stealFrom = (stealFrom + 2) % size
			even = true
		} else {
			stealFrom = (stealFrom + 1) % size
			even = false
		}
	}

	return i
}

func (e elves) rebalance(size int, lastElf int) (newSize int, nextElf int) {
	j := 0
	totalPresents := 0
	for i := 0; i < size; i++ {
		if e[i].presents > 0 {
			totalPresents += e[i].presents
			e[j] = e[i]
			e[j].elfToLeft = j + 1
			if i == lastElf {
				nextElf = j
			}
			// fmt.Printf("Keep %d, %d presents\n", e[j].index, e[j].presents)
			j++
		}
	}

	// They're sitting in a circle
	e[j-1].elfToLeft = 0
	fmt.Printf("Rebalanced %d down to %d elves, %d presents\n", size, j, totalPresents)
	newSize = j

	return
}

func main() {
	size := 3018458
	// size := 100
	// size := 11
	var e elves
	e = make([]elf, size)
	e.init()
	nextElf := 0
	// lastElf := e.presents(0)
	lastElf := 0
	for size > 1 {
		lastElf = e.circles(nextElf, size)
		// fmt.Printf("Elf %d will go next\n", e[lastElf].index)
		size, nextElf = e.rebalance(size, lastElf)
		// fmt.Printf("Is it still elf %d will go next?\n", e[nextElf].index)
	}

	// fmt.Println(lastElf + 1)
	fmt.Printf("%#v\n", e[nextElf])
}

func (e elves) presents(firstElf int) (last int) {
	i := firstElf
	j := 1000000
	for {
		// i is the elf doing the stealing
		elfToLeft := e[i].elfToLeft
		j--
		if j == 0 {
			j = 1000000
			fmt.Printf("Elf %d steals from elf %d\n", i, elfToLeft)
		}

		// Finish when there is only one elf left
		if elfToLeft == i {
			return elfToLeft
		}

		// Elf at i steals presents from elf to its left
		// e[i].presents += e[elfToLeft].presents
		// e[elfToLeft].presents = 0

		// Now it's the turn of the next elf on the left to become i
		e[i].elfToLeft = e[elfToLeft].elfToLeft
		i = e[elfToLeft].elfToLeft
	}
}
