package main

import (
	"fmt"
	"math"
	"os"
)

// Chip can be cobalt, curium, plutonium, promethium, ruthenium,
// Generator can be cobalt, curium, plutonium, promethium, ruthenium
type chips int
type generators int

const (
	cobalt     = 0x01
	hydrogen   = 0x01
	curium     = 0x02
	lithium    = 0x02
	plutonium  = 0x04
	promethium = 0x08
	ruthenium  = 0x10
	elerium    = 0x20
	dilithium  = 0x40
)

var all int
var allElements []int

type floor struct {
	id         int // note we will number them from zero to three
	chips      int
	generators int
}

type move struct {
	destFloor  int
	chips      int
	generators int
}

type status struct {
	floors   [4]floor
	elevator int // where the elevator is
}

var knownStates map[status]int

func chipHasGenerator(chip int, generators int) bool {
	return (chip&generators == chip)
}

func allChipsHaveGenerators(chips int, generators int) bool {
	// if !(chips&generators == chips) {
	// 	fmt.Printf("%x is unprotected - ", chips-chips&generators)
	// }
	return (chips&generators == chips)
}

func (s *status) isTheSame(x *status) bool {
	if x.elevator != s.elevator {
		return false
	}

	for f := 0; f < 4; f++ {
		if x.floors[f].chips != s.floors[f].chips ||
			x.floors[f].generators != s.floors[f].generators {
			return false
		}
	}

	fmt.Printf("%v is the same as %v\n", s, x)
	return true
}

func (s *status) moveOK(m move) (result *status, ok bool) {
	// Move is do-able if
	// between one and two total chips & generators move

	// move is up or down one
	if int(math.Abs(float64(m.destFloor-s.elevator))) != 1 {
		fmt.Printf("Moved %d floors ", int(math.Abs(float64(m.destFloor-s.elevator))))
		fmt.Printf("From %d to %d\n", s.elevator, m.destFloor)
		panic("Didn't move one floor")
	}

	result = s.afterMove(m)
	// chip can't be left on a floor with any other generators unless it has its own
	for f := 0; f < 4; f++ {
		// No need to worry if there aren't any generators here
		if result.floors[f].generators == 0 {
			continue
		}

		if !allChipsHaveGenerators(result.floors[f].chips, result.floors[f].generators) {
			// fmt.Printf("Move %v would leave chip unprotected on floor %d\n", m, f)
			return nil, false
		}
	}

	// // Don't want to go back to a  previous state
	// for x := s.previous; x != nil; x = x.previous {
	// 	if s.isTheSame(x) {
	// 		// fmt.Printf("Matched previous state\n")
	// 		return nil, false
	// 	}
	// }

	// result.previous = s
	// fmt.Printf("Move is OK\n")
	return result, true
}

func (s *status) finished(chipCount int) bool {
	if s.elevator != 3 || s.floors[3].chips != all || s.floors[3].generators != all {
		return false
	}

	return true
}

func makeMoves(m []move, movingChips int, movingGens int, f int) (nm []move) {
	// Move up if we can
	// fmt.Printf("Moving from floor %d - ", f)
	nm = m
	if f < 3 {
		// fmt.Printf(" up ")
		nm = append(nm, move{destFloor: f + 1, chips: movingChips, generators: movingGens})
	}
	if f > 0 {
		// fmt.Printf(" down ")
		nm = append(nm, move{destFloor: f - 1, chips: movingChips, generators: movingGens})
	}

	// fmt.Printf("\n")
	return nm
}

// Find all the possible combinations of moves (without checking whether they are OK)
func (s *status) possibleMoves() (m []move) {
	// We can theoretically move any pair of chips and/or generators up or down
	chipsOnThisFloor := s.floors[s.elevator].chips
	gensOnThisFloor := s.floors[s.elevator].generators

	for i, element := range allElements {
		if element&chipsOnThisFloor == element {
			// Send just one chip and just one generator
			m = makeMoves(m, element, 0, s.elevator)

			// Look for another chip we could send with it, or non-matching generators
			if i < len(allElements)-1 {
				for _, element2 := range allElements[i+1:] {
					if element2&chipsOnThisFloor == element2 {
						m = makeMoves(m, element|element2, 0, s.elevator)
					}
					if element2&gensOnThisFloor == element2 {
						m = makeMoves(m, element, element2, s.elevator)
					}
				}
			}
		}

		// Send the matching pair of chip & generator
		if element&chipsOnThisFloor&gensOnThisFloor == element {
			m = makeMoves(m, element, element, s.elevator)
		}

		if element&gensOnThisFloor == element {
			m = makeMoves(m, 0, element, s.elevator)

			// Look for another generator we could send with it , or non-matching chips
			for _, element2 := range allElements[i+1:] {
				if element2&chipsOnThisFloor == element2 {
					m = makeMoves(m, element2, element, s.elevator)
				}
				if element2&gensOnThisFloor == element2 {
					m = makeMoves(m, 0, element|element2, s.elevator)
				}
			}
		}
	}

	// fmt.Printf("Generated %d possible moves from state %v\n", len(m), s)
	// for i := range m {
	// 	fmt.Printf("  %#v\n", m[i])
	// }
	return m
}

// TODO!! From here down make it integer matches
// Have a list of previous states that is keyed by state?

func (s *status) afterMove(m move) (ns *status) {
	ns = &status{}
	ns.elevator = m.destFloor
	for floor := 0; floor < 4; floor++ {
		ns.floors[floor].id = floor
		switch {
		case floor == s.elevator:
			// Floor we're moving from
			ns.floors[floor].chips = s.floors[floor].chips & (0xFF - m.chips)
			ns.floors[floor].generators = s.floors[floor].generators & (0xFF - m.generators)
		case floor == ns.elevator:
			// Floor we're moving to. Any existing chips or generators remain here
			ns.floors[floor].chips = s.floors[floor].chips | m.chips
			ns.floors[floor].generators = s.floors[floor].generators | m.generators
		default:
			ns.floors[floor].chips = s.floors[floor].chips
			ns.floors[floor].generators = s.floors[floor].generators
		}
	}

	// fmt.Printf("-> Move %v gives %v\n", m, ns)
	return ns
}

func (s *status) possibleStates(move int) (ns []*status) {
	candidates := s.possibleMoves()
	// fmt.Printf("%d moves under consideration\n", len(candidates))
	for _, candidate := range candidates {
		if rs, ok := s.moveOK(candidate); ok {
			if rs.finished(5) {
				fmt.Printf("Found a finisher so let's return\n")
				ns = append([]*status{}, rs)
				return
			}

			if _, ok := knownStates[*rs]; !ok {
				knownStates[*rs] = move
				ns = append(ns, rs)
			}
		}
	}

	// fmt.Printf("State has %d possible new states:\n", len(ns))
	// for _, n := range ns {
	// 	fmt.Printf(". %v\n", n)
	// }
	return
}

func main() {
	// Initial status
	// f0 := floor{id: 0, chips: []chip{"promethium"}, generators: []generator{"promethium"}}
	// f1 := floor{id: 1, generators: []generator{"cobalt", "curium", "ruthenium", "plutonium"}}
	// f2 := floor{id: 2, chips: []chip{"cobalt", "curium", "ruthenium", "plutonium"}}
	s := status{elevator: 0}
	s.floors[0] = floor{id: 0, chips: promethium | elerium | dilithium, generators: promethium | elerium | dilithium}
	s.floors[1] = floor{id: 1, generators: cobalt | curium | ruthenium | plutonium}
	s.floors[2] = floor{id: 2, chips: cobalt | curium | ruthenium | plutonium}
	s.floors[3] = floor{id: 3}
	all = cobalt | curium | plutonium | promethium | ruthenium | elerium | dilithium
	allElements = []int{cobalt, curium, plutonium, promethium, ruthenium, elerium, dilithium}

	// Test that finished is OK
	st := status{elevator: 3}
	// st.floors[3] = floor{id: 3, chips: []chip{"cobalt", "curium", "ruthenium", "plutonium", "promethium"},
	// 	generators: []generator{"cobalt", "curium", "ruthenium", "plutonium", "promethium"}}
	st.floors[3] = floor{id: 3, chips: cobalt | curium | ruthenium | plutonium | promethium | elerium | dilithium,
		generators: cobalt | curium | ruthenium | plutonium | promethium | elerium | dilithium}

	if s.finished(7) {
		fmt.Println("Initial status unexpectedly marked as finished")
	}

	if !st.finished(7) {
		fmt.Println("Final status unexpectedly not marked as finished")
	}

	// Test that a move gives us what we expect
	a := status{elevator: 1}
	a.floors[0] = floor{id: 0, chips: dilithium | elerium, generators: dilithium | elerium}
	a.floors[1] = floor{id: 1, chips: promethium, generators: cobalt | curium | ruthenium | plutonium | promethium}
	a.floors[2] = floor{id: 2, chips: cobalt | curium | ruthenium | plutonium}
	a.floors[3] = floor{id: 3}
	m := move{destFloor: 1, chips: promethium, generators: promethium}
	aa := s.afterMove(m)
	if !aa.isTheSame(&a) {
		fmt.Printf("After move %v got %v\n", m, aa)
		fmt.Printf("Expected %v", a)
		panic("Wrong")
	}

	// Test version
	// s.floors[0] = floor{id: 0, chips: hydrogen | lithium}
	// s.floors[1] = floor{id: 1, generators: hydrogen}
	// s.floors[2] = floor{id: 2, generators: lithium}
	// s.floors[3] = floor{id: 3}
	// all = hydrogen | lithium
	// allElements = []int{hydrogen, lithium}

	moves := 0
	states := []*status{&s}
	knownStates = make(map[status]int, 1)
	knownStates[s] = 0
	fmt.Printf("Start from %v\n", s)
	for {
		newStates := []*status{}

		for _, state := range states {
			if state.finished(len(allElements)) {
				fmt.Printf("Found a solution with %d moves\n", moves)
				os.Exit(0)
			}

			candidates := state.possibleStates(moves)
			newStates = append(newStates, candidates...)
		}
		states = newStates
		moves++
		fmt.Printf("*** Move: %d New states under consideration %d\n", moves, len(states))
		if len(states) == 0 {
			os.Exit(2)
		}
		if moves > 200 {
			os.Exit(1)
		}
	}
}
