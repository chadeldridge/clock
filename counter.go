package rpgtools

import "github.com/chadeldridge/rpgtools/incrementers"

type Counter struct {
	incrementers.Incrementer
}

// NewCounter creates a new counter with a minimum and maximum value.
func NewCounter(min, max int) Counter {
	c := Counter{incrementers.NewIncrementer(min, max)}
	c.IncrementBy(1)
	return c
}

// NewCounterWithTicks creates a new counter with a minimum, maximum, and a starting value.
func NewCounterWithTicks(min, max, ticks int) Counter {
	c := Counter{incrementers.NewIncrementerWithValue(min, max, ticks)}
	c.IncrementBy(1)
	return c
}

// NewCounterFromJSON creates a new counter from a JSON representation.
func NewCounterFromJSON(data []byte) (Counter, error) {
	var c Counter
	err := c.UnmarshalJSON(data)
	return c, err
}
