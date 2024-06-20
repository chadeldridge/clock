package incrementers

import "fmt"

type Counter interface {
	Inc() int
	Value() int
	Original() int

	IsFull() bool
	IsEmpty() bool

	Increment()
	Decrement()
	Add(int)
	Remove(int)

	SetMax(int)
	SetIncrementer(int)
	SetValue(int)
	SetOriginalValue(int)

	Empty()
	Reset()

	String() string
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

// NewCounter creates a new counter with a minimum value of 0 and no maximum value.
func NewCounter() Counter {
	return &ClampedIncrementer{min: 0, max: 0, Incrementer: Incrementer{inc: 1}}
}

// NewCounterWithValue creates a new counter with a starting value of 0 or greater.
func NewCounterWithValue(val int) Counter {
	if val < 0 {
		val = 0
	}

	return &ClampedIncrementer{min: 0, max: 0, Incrementer: Incrementer{val: val, inc: 1, orig: val}}
}

func NewCounterFromJSON(data []byte) (Counter, error) {
	var i ClampedIncrementer
	if err := i.UnmarshalJSON(data); err != nil {
		return nil, err
	}

	if i.min != 0 {
		return nil, fmt.Errorf("invalid Counter: min must be 0")
	}

	if i.max != 0 {
		return nil, fmt.Errorf("invalid Counter: max must be 0")
	}

	return &i, nil
}
