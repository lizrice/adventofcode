package main

import (
	"fmt"
	"os"
)

// If I drop at time t
// Disc n pos = (pos0 + n) % size
//.          Disc 2 pos =

type disc struct {
	initial int
	size    int
	index   int
}

func pos(d disc, t int) int {
	return (d.initial + d.index + t) % d.size
}

var discs = []disc{
	{initial: 0, size: 7, index: 1},
	{initial: 0, size: 13, index: 2},
	{initial: 2, size: 3, index: 3},
	{initial: 2, size: 5, index: 4},
	{initial: 0, size: 17, index: 5},
	{initial: 7, size: 19, index: 6},
	{initial: 0, size: 11, index: 7},
}

func main() {
	// discs = []disc{
	// 	{initial: 4, size: 5, index: 1},
	// 	{initial: 1, size: 2, index: 2},
	// }

	t := 0
	var finished bool
	for {
		t++
		fmt.Printf("\nTime %d: ", t)
		for _, d := range discs {
			p := pos(d, t)
			fmt.Printf("Disc #%d pos %d ", d.index, p)
			if p != 0 {
				finished = false
				break
			}
			finished = true
		}
		if finished {
			fmt.Printf("\nStart at time %d\n", t)
			os.Exit(1)
		}
	}
}
