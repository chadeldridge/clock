package rpgtools

import (
	"fmt"

	"github.com/chadeldridge/rpgtools/incrementers"
)

type Clock interface {
	Min() int
	Max() int
	Value() int
	Inc() int
	Original() int

	IsFull() bool
	IsEmpty() bool
	IsMin() bool

	Advance()
	Increment()
	Decrement()
	Add(int)
	Remove(int)

	SetMax(int)
	SetValue(int)

	Fill()
	Empty()
	Reset()

	String() string
	MarshalJSON() ([]byte, error)
}

// NewClock creates a new clock with from 0 to the maximum value of steps.
func NewClock(steps int) Clock {
	i := incrementers.NewIncrementer(0, steps)
	i.IncrementBy(1)
	return &i
}

// NewClockWithTicks creates a new clock from 0 to the maximum steps, with a starting value of ticks.
func NewClockWithTicks(steps, ticks int) Clock {
	i := incrementers.NewIncrementerWithValue(0, steps, ticks)
	i.IncrementBy(1)
	return &i
}

// NewClockFromJSON creates a new clock from a JSON representation.
func NewClockFromJSON(data []byte) (Clock, error) {
	var i incrementers.Incrementer
	if err := i.UnmarshalJSON(data); err != nil {
		return nil, err
	}

	if i.Min() != 0 {
		return nil, fmt.Errorf("invalid Clock: min must be 0")
	}

	if i.Inc() != 1 {
		return nil, fmt.Errorf("invalid Clock: inc must be 1")
	}

	return &i, nil
}
