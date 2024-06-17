package rpgtools

import (
	"fmt"

	"github.com/chadeldridge/rpgtools/incrementers"
)

type Counter interface {
	Value() int
	Inc() int
	Original() int

	IsEmpty() bool

	Advance()
	Increment()
	Decrement()
	Add(int)
	Remove(int)

	SetValue(int)
	SetIncrementer(int)

	Empty()
	Reset()

	String() string
	MarshalJSON() ([]byte, error)
}

// NewCounter creates a new counter with a minimum value of 0 and no maximum value.
func NewCounter() Counter {
	i := incrementers.NewIncrementerClamped(0, 100)
	i.SetNoMax(true)
	return &i
}

// NewCounterWithValue creates a new counter with a minimum, maximum, and a starting value.
func NewCounterWithValue(val int) Counter {
	if val < 0 {
		return nil
	}

	i := incrementers.NewIncrementerClampedWithValue(0, 100, val)
	i.SetNoMax(true)
	return &i
}

// NewCounterFromJSON creates a new counter from a JSON representation.
func NewCounterFromJSON(data []byte) (Counter, error) {
	var i incrementers.Incrementer
	if err := i.UnmarshalJSON(data); err != nil {
		return &i, err
	}

	if i.Min() != 0 {
		return &i, fmt.Errorf("invalid Counter: min must be 0")
	}

	i.SetNoMin(false)
	i.SetNoMax(true)

	return &i, nil
}
