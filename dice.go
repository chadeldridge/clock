package rpgtools

import (
	"math/rand/v2"
)

type Die int

const (
	D2   Die = 2
	D4   Die = 4
	D6   Die = 6
	D8   Die = 8
	D10  Die = 10
	D12  Die = 12
	D20  Die = 20
	D100 Die = 100
)

type DicePool int

type DieResults struct {
	Die
	Highest int
	Lowest  int
	Sum     int
	All     []int
}

// NewDieResults creates a new DieResults struct.
func NewDieResults(die Die) DieResults {
	r := DieResults{Die: die}
	r.Highest = 0
	r.Lowest = int(die)
	r.All = []int{}

	return r
}

// Roll the given die.
func Roll(d Die) int { return rand.IntN(int(d)) + 1 }

// Roll the given die the given number of times.
func (d Die) Roll(pool int) []int {
	var r []int
	for range pool {
		r = append(r, Roll(d))
	}

	return r
}

// NewDicePool creates a new DicePool with the given number of dice.
func NewDicePool(n int) DicePool { return DicePool(n) }

// Roll the given die pool.
func (p DicePool) Roll(die Die) DieResults {
	var r DieResults
	r.All = die.Roll(int(p))
	r.Lowest = int(die)

	for _, i := range r.All {
		if i > r.Highest {
			r.Highest = i
		}

		if i < r.Lowest {
			r.Lowest = i
		}

		r.Sum += i
	}

	return r
}
