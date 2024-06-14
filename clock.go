package rpgtools

type Clock struct {
	UIncrementer
}

// NewClock creates a new clock with the given number of steps.
func NewClock(steps int) Clock {
	c := Clock{NewUIncrementer(steps)}
	c.IncrementBy(1)
	return c
}

// NewClockWithTicks creates a new clock with the given number of steps and ticks.
func NewClockWithTicks(steps, ticks int) Clock {
	c := Clock{NewUIncrementerWithValue(steps, ticks)}
	c.IncrementBy(1)
	return c
}

// NewClockFromJSON creates a new clock from a JSON representation.
func NewClockFromJSON(data []byte) (Clock, error) {
	var c Clock
	err := c.UnmarshalJSON(data)
	return c, err
}
