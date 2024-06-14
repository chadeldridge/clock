package clock

import "fmt"

type Clock struct {
	Steps int
	ticks int
}

// New creates a new clock with the given number of steps.
func New(steps int) Clock { return Clock{Steps: steps} }

// NewWithTicks creates a new clock with the given number of steps and ticks.
func NewWithTicks(steps, ticks int) Clock { return Clock{Steps: steps, ticks: ticks} }

// NewFromJSON creates a new clock from a JSON representation.
func NewFromJSON(data []byte) (Clock, error) {
	var c Clock
	err := c.UnmarshalJSON(data)
	return c, err
}

// Ticks returns the current number of ticks on the clock.
func (c Clock) Ticks() int { return c.ticks }

// IsFull returns true if the clock is full.
func (c Clock) IsFull() bool { return c.ticks == c.Steps }

// IsEmpty returns true if the clock is empty.
func (c Clock) IsEmpty() bool { return c.ticks == 0 }

// Tick increments the clock by one tick.
func (c *Clock) Tick() {
	if c.ticks == c.Steps {
		return
	}

	c.ticks++
}

// Add increments the clock by the given number of ticks.
func (c *Clock) Add(ticks int) {
	if c.ticks == c.Steps {
		return
	}

	c.ticks += ticks

	if c.ticks > c.Steps {
		c.ticks = c.Steps
	}
}

// Remove decrements the clock by the given number of ticks.
func (c *Clock) Remove(ticks int) {
	c.ticks -= ticks

	if c.ticks < 0 {
		c.ticks = 0
	}
}

// SetSteps sets the size of the clock to the given number of steps.
func (c *Clock) SetSteps(steps int) {
	if steps < 1 {
		steps = 1
	}

	c.Steps = steps

	if c.ticks > c.Steps {
		c.ticks = c.Steps
	}
}

// AddSteps increases the size of the clock by the given number of steps.
func (c *Clock) AddSteps(steps int) {
	c.Steps += steps
}

// RemoveSteps decreases the size of the clock by the given number of steps.
func (c *Clock) RemoveSteps(steps int) {
	c.Steps -= steps

	if c.Steps < 1 {
		c.Steps = 1
	}

	if c.ticks > c.Steps {
		c.ticks = c.Steps
	}
}

// Fill sets the clock to full.
func (c *Clock) Fill() {
	c.ticks = c.Steps
}

// Empty sets the clock to 0.
func (c *Clock) Empty() {
	c.ticks = 0
}

// Reset sets the clock to 0.
func (c *Clock) Reset() {
	c.ticks = 0
}

// String returns a string representation of the clock.
func (c Clock) String() string {
	return fmt.Sprintf("%d/%d", c.ticks, c.Steps)
}

// MarshalJSON returns a JSON representation of the clock.
func (c Clock) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"steps":%d,"ticks":%d}`, c.Steps, c.ticks)), nil
}

// UnmarshalJSON parses a JSON representation of the clock.
func (c *Clock) UnmarshalJSON(data []byte) error {
	var steps, ticks int
	_, err := fmt.Sscanf(string(data), `{"steps":%d,"ticks":%d}`, &steps, &ticks)
	if err != nil {
		return err
	}

	if steps < 1 {
		return fmt.Errorf("invalid clock: steps must be greater than 0")
	}

	if ticks < 0 {
		return fmt.Errorf("invalid clock: ticks must be greater than or equal to 0")
	}

	if ticks > steps {
		return fmt.Errorf("invalid clock: ticks must be less than or equal to steps")
	}

	c.Steps = steps
	c.ticks = ticks

	return nil
}
