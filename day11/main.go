package main

import (
	"fmt"
	"math"
	"os"
)

// Chip can be cobalt, curium, plutonium, promethium, ruthenium,
// Generator can be cobalt, curium, plutonium, promethium, ruthenium
type chip string
type generator string

type floor struct {
	id         int // note we will number them from zero to three
	chips      []chip
	generators []generator
}

type move struct {
	destFloor  int
	chips      []chip
	generators []generator
}

type status struct {
	floors   [4]floor
	elevator int // where the elevator is
	moves    int // Count of moves to get here
	previous *status
}

func matches(c chip, g generator) bool {
	return string(c) == string(g)
}

func (s *status) isTheSame(x *status) bool {
	if x.elevator != s.elevator {
		return false
	}

	for f := 0; f < 4; f++ {
		if len(x.floors[f].chips) != len(s.floors[f].chips) ||
			len(x.floors[f].generators) != len(s.floors[f].generators) {
			return false
		}

		for _, c := range x.floors[f].chips {
			here := false
			for _, cc := range s.floors[f].chips {
				if c == cc {
					here = true
				}
			}
			if !here {
				return false
			}
		}

		for _, g := range x.floors[f].generators {
			here := false
			for _, gg := range s.floors[f].generators {
				if g == gg {
					here = true
				}
			}
			if !here {
				return false
			}
		}
	}

	return true
}

func (s *status) moveOK(m move) (result *status, ok bool) {
	// Move is do-able if
	// between one and two total chips & generators move
	if len(m.chips)+len(m.generators) > 2 {
		panic("Too many items moving")
	}

	// move is up or down one
	if int(math.Abs(float64(m.destFloor-s.elevator))) != 1 {
		fmt.Printf("Moved %d floors ", int(math.Abs(float64(m.destFloor-s.elevator))))
		fmt.Printf("From %d to %d", s.elevator, m.destFloor)
		panic("Didn't move one floor")
	}

	result = s.afterMove(m)
	// chip can't be left on a floor with any other generators unless it has its own
	for f := 0; f < 4; f++ {
		protected := false
		// No need to worry if there aren't any generators here
		if len(result.floors[f].generators) == 0 {
			continue
		}
		for _, c := range result.floors[f].chips {
			for _, g := range result.floors[f].generators {
				if matches(c, g) {
					protected = true
				}
			}
			if !protected {
				// fmt.Printf("Chip %s would be unprotected on floor %d\n", c, f)
				return nil, false
			}
		}

	}

	// Don't want to go back to a  previous state
	for x := s.previous; x != nil; x = x.previous {
		if s.isTheSame(x) {
			return nil, false
		}
	}

	result.previous = s
	return result, true
}

func (s *status) finished(chipCount int) bool {
	if s.elevator != 3 || len(s.floors[3].chips) != chipCount || len(s.floors[3].generators) != chipCount {
		return false
	}

	return true
}

func makeMoves(m []move, movingChips []chip, movingGens []generator, f int) (nm []move) {
	// Move up if we can
	// fmt.Printf("Moving from floor %d", f)
	if f < 3 {
		nm = append(m, move{destFloor: f + 1, chips: movingChips, generators: movingGens})
	}
	if f > 0 {
		nm = append(m, move{destFloor: f - 1, chips: movingChips, generators: movingGens})
	}
	return nm
}

// Find all the possible combinations of moves (without checking whether they are OK)
func (s *status) possibleMoves() (m []move) {
	// We can theoretically move any pair of chips and/or generators up or down
	chipsOnThisFloor := s.floors[s.elevator].chips
	gensOnThisFloor := s.floors[s.elevator].generators

	for i, c := range chipsOnThisFloor {
		// Sending just one chip
		m = makeMoves(m, []chip{c}, []generator{}, s.elevator)

		if len(chipsOnThisFloor) > (i + 1) {
			for _, c2 := range chipsOnThisFloor[i+1:] {
				// Comnbinations of two chips
				m = makeMoves(m, []chip{c, c2}, []generator{}, s.elevator)
			}
		}

		if len(gensOnThisFloor) > 0 {
			for _, g := range gensOnThisFloor {
				// Combinations of one chip one generator
				m = makeMoves(m, []chip{c}, []generator{g}, s.elevator)
			}
		}
	}

	for i, g := range gensOnThisFloor {
		// Sending just one generator
		m = makeMoves(m, []chip{}, []generator{g}, s.elevator)

		if len(gensOnThisFloor) > (i + 1) {
			for _, g2 := range gensOnThisFloor[i+1:] {
				m = makeMoves(m, []chip{}, []generator{g, g2}, s.elevator)
			}
		}
	}

	// fmt.Printf("Generated %d possible moves from a state after %d moves\n", len(m), s.moves)
	// for i := range m {
	// 	fmt.Printf("  %#v\n", m[i])
	// }
	return m
}

func (s *status) afterMove(m move) (ns *status) {
	ns = &status{}
	ns.elevator = m.destFloor
	ns.moves = s.moves + 1
	for floor := 0; floor < 4; floor++ {
		switch {
		case floor == s.elevator:
			// Floor we're moving from
			for _, chip := range s.floors[floor].chips {
				chipMoved := false
				for _, movingChip := range m.chips {
					if chip == movingChip {
						chipMoved = true
					}
				}

				if chipMoved {
					ns.floors[ns.elevator].chips = append(ns.floors[ns.elevator].chips, chip)
				} else {
					ns.floors[s.elevator].chips = append(ns.floors[s.elevator].chips, chip)
				}
			}

			for _, gen := range s.floors[floor].generators {
				genMoved := false
				for _, movingGen := range m.generators {
					if gen == movingGen {
						genMoved = true
					}
				}

				if genMoved {
					ns.floors[ns.elevator].generators = append(ns.floors[ns.elevator].generators, gen)
				} else {
					ns.floors[s.elevator].generators = append(ns.floors[s.elevator].generators, gen)
				}
			}
		case floor == ns.elevator:
			// Floor we're moving to. Any existing chips or generators remain here
			ns.floors[floor].chips = append(ns.floors[floor].chips, s.floors[floor].chips...)
			ns.floors[floor].generators = append(ns.floors[floor].generators, s.floors[floor].generators...)

		default:
			ns.floors[floor].chips = s.floors[floor].chips
			ns.floors[floor].generators = s.floors[floor].generators
		}
	}

	return
}

func (s *status) possibleStates() (ns []*status) {
	candidates := s.possibleMoves()
	for _, candidate := range candidates {
		if rs, ok := s.moveOK(candidate); ok {
			// fmt.Println("Move is OK")
			ns = append(ns, rs)
		}
	}

	// fmt.Printf("State has %d possible new states\n", len(ns))
	return
}

func main() {
	// Initial status
	f0 := floor{id: 0, chips: []chip{"promethium"}, generators: []generator{"promethium"}}
	f1 := floor{id: 1, generators: []generator{"cobalt", "curium", "ruthenium", "plutonium"}}
	f2 := floor{id: 2, chips: []chip{"cobalt", "curium", "ruthenium", "plutonium"}}
	s := status{elevator: 0}
	s.floors[0] = f0
	s.floors[1] = f1
	s.floors[2] = f2

	// Test that finished is OK
	st := status{elevator: 3}
	st.floors[3] = floor{id: 3, chips: []chip{"cobalt", "curium", "ruthenium", "plutonium", "promethium"},
		generators: []generator{"cobalt", "curium", "ruthenium", "plutonium", "promethium"}}

	if s.finished(5) {
		fmt.Println("Initial status unexpectedly marked as finished")
	}

	if !st.finished(5) {
		fmt.Println("Final status unexpectedly not marked as finished")
	}

	moves := 0
	states := []*status{&s}
	for {
		newStates := []*status{}
		for _, state := range states {
			if state.finished(5) {
				fmt.Printf("Found a solution with %d moves\n", s.moves)
				os.Exit(0)
			}
			candidates := state.possibleStates()
			newStates = append(newStates, candidates...)
		}
		states = newStates
		moves++
		fmt.Printf("*** Move: %d Possible states %d\n", moves, len(states))
		// if moves > 2 {
		// 	os.Exit(1)
		// }
	}
}
