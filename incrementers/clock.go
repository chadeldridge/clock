package incrementers

import "fmt"

type Clock interface {
	Max() int
	Value() int
	Original() int

	IsFull() bool
	IsEmpty() bool

	Increment()
	Decrement()
	Add(int)
	Remove(int)

	SetMax(int)
	SetValue(int)
	SetOriginalValue(int)

	Fill()
	Empty()
	Reset()

	String() string
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

// NewClock creates a new clock with from 0 to the maximum value of steps.
func NewClock(steps int) Clock {
	return &ClampedIncrementer{min: 0, max: steps, Incrementer: Incrementer{inc: 1}}
}

// NewClockWithTicks creates a new clock from 0 to the maximum steps, with a starting value of ticks.
func NewClockWithTicks(steps, ticks int) Clock {
	return &ClampedIncrementer{min: 0, max: steps, Incrementer: Incrementer{inc: 1, val: ticks, orig: ticks}}
}

// NewClockFromJSON creates a new clock from a JSON representation.
func NewClockFromJSON(data []byte) (Clock, error) {
	var i ClampedIncrementer
	if err := i.UnmarshalJSON(data); err != nil {
		return nil, err
	}

	if i.min != 0 {
		return nil, fmt.Errorf("invalid Clock: min must be 0")
	}

	if i.max < 1 {
		return nil, fmt.Errorf("invalid Clock: max must be greater than 0")
	}

	if i.inc != 1 {
		return nil, fmt.Errorf("invalid Clock: inc must be 1")
	}

	return &i, nil
}
