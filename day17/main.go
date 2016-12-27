package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type state struct {
	x        int
	y        int
	sequence string
}

var passcode = "qtetzkpl"

func getMD5Hash(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (s *state) possibleMoves() (possible []*state) {
	hash := getMD5Hash(passcode + s.sequence)
	// up
	if hash[0] >= 'b' && hash[0] <= 'f' && s.y > 0 {
		possible = append(possible, &state{x: s.x, y: s.y - 1, sequence: s.sequence + "U"})
	}
	// down
	if hash[1] >= 'b' && hash[0] <= 'f' && s.y < 3 {
		possible = append(possible, &state{x: s.x, y: s.y + 1, sequence: s.sequence + "D"})
	}
	// left
	if hash[2] >= 'b' && hash[0] <= 'f' && s.x > 0 {
		possible = append(possible, &state{x: s.x - 1, y: s.y, sequence: s.sequence + "L"})
	}
	// down
	if hash[3] >= 'b' && hash[0] <= 'f' && s.x < 3 {
		possible = append(possible, &state{x: s.x + 1, y: s.y, sequence: s.sequence + "R"})
	}
	return possible
}

func (s *state) finished() bool {
	return (s.x == 3 && s.y == 3)
}

func main() {
	initial := state{}
	s := []*state{}
	s = append(s, &initial)
	longest := &initial
	for len(s) > 0 {
		ns := []*state{}
		for _, p := range s {
			if p.finished() {
				if len(p.sequence) > len(longest.sequence) {
					fmt.Println(len(p.sequence))
					longest = p
				}
			} else {
				ns = append(ns, p.possibleMoves()...)
			}
		}
		s = ns
	}
	fmt.Println("No possible outcomes")
}
