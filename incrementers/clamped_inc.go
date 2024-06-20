package incrementers

import (
	"fmt"
)

type ClampedIncrementer struct {
	min int
	max int
	Incrementer
}

// NewClampedIncrementer creates a new counter with a minimum value of 0 and no maximum value.
func NewClampedIncrementer(min, max int) ClampedIncrementer {
	return ClampedIncrementer{min: min, max: max, Incrementer: Incrementer{inc: 1}}
}

// NewClampedIncrementerWithValue creates a new counter with a starting value of 0 or greater.
func NewClampedIncrementerWithValue(min, max, val int) ClampedIncrementer {
	if max == 0 {
		val = ClampMin(val, min)
	} else {
		val = Clamp(val, min, max)
	}

	return ClampedIncrementer{min: min, max: max, Incrementer: Incrementer{val: val, inc: 1, orig: val}}
}

// NewClampedIncrementerFromJSON creates a new counter from a JSON representation.
func NewClampedIncrementerFromJSON(data []byte) (ClampedIncrementer, error) {
	var c ClampedIncrementer
	err := c.UnmarshalJSON(data)
	return c, err
}

// IsFull returns true if the counter is at the maximum value.
func (c ClampedIncrementer) IsFull() bool { return c.val == c.max }

// Min returns the minimum value of the incrementer.
func (c ClampedIncrementer) Min() int { return c.min }

// Max returns the maximum value of the incrementer.
func (c ClampedIncrementer) Max() int { return c.max }

// Increment increases the counter by the incrementer value clamped.
func (c *ClampedIncrementer) Increment() { c.val = c.val + c.inc; c.Clamp() }

// Decrement decreases the counter by the incrementer value clamped.
func (c *ClampedIncrementer) Decrement() { c.val = c.val - c.inc; c.Clamp() }

// Add increases the counter by the given number of val clamped.
func (c *ClampedIncrementer) Add(val int) { c.val = c.val + val; c.Clamp() }

// Remove decreases the counter by the given number of val clamped.
func (c *ClampedIncrementer) Remove(val int) { c.val = c.val - val; c.Clamp() }

// SetMin sets the minimum value of the incrementer.
func (c *ClampedIncrementer) SetMin(min int) { c.min = min; c.Clamp() }

// SetMax sets the maximmum value of the incrementer.
func (c *ClampedIncrementer) SetMax(max int) { c.max = max; c.Clamp() }

// SetValue sets the counter value to the given number of val clamped.
func (c *ClampedIncrementer) SetValue(val int) { c.val = val; c.Clamp() }

// SetOrginalValue sets the counter's original value to the given number of val clamped.
func (c *ClampedIncrementer) SetOriginalValue(val int) { c.orig = val; c.ClampOriginalValue() }

// Clamp sets the value to the min max range. If max is 0 then the value will be clamped to the minimum.
func (c *ClampedIncrementer) Clamp() {
	if c.max == 0 {
		c.val = ClampMin(c.val, c.min)
		return
	}

	c.val = Clamp(c.val, c.min, c.max)
}

// ClampOriginalValue sets the original value to the min max range. If max is 0 then the value will
// be clamped to the minimum.
func (c *ClampedIncrementer) ClampOriginalValue() {
	if c.max == 0 {
		c.orig = ClampMin(c.orig, c.min)
		return
	}

	c.orig = Clamp(c.orig, c.min, c.max)
}

// Fill sets the counter to the maximum value.
func (c *ClampedIncrementer) Fill() { c.val = c.max }

// Floor sets the counter to the minimum value.
func (c *ClampedIncrementer) Floor() { c.val = c.min }

// String returns a string representation of the Incrementer.
func (c ClampedIncrementer) String() string { return fmt.Sprintf("%d/%d", c.val, c.max) }

// MarshalJSON returns a JSON representation of the counter.
func (c ClampedIncrementer) MarshalJSON() ([]byte, error) {
	j, _ := c.Incrementer.MarshalJSON()
	return []byte(fmt.Sprintf(`{"min":%d,"max":%d,"incrementer":%s}`, c.min, c.max, j)), nil
}

func (c *ClampedIncrementer) UnmarshalJSON(data []byte) error {
	var iData []byte

	if data == nil {
		return fmt.Errorf("Incrementer.UnmarshalJSON(): data was nil")
	}

	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	data = data[1 : len(data)-1]
	if _, err := fmt.Sscanf(string(data), `"min":%d,"max":%d,"incrementer":%s`, &c.min, &c.max, &iData); err != nil {
		return err
	}

	if c.max != 0 && c.min >= c.max {
		return fmt.Errorf("invalid ClampedIncrementer: min must be less than max")
	}

	i, err := NewIncrementerFromJSON(iData)
	if err != nil {
		return err
	}

	c.Incrementer = i
	if c.val < c.min {
		return fmt.Errorf("invalid ClampedIncrementer: Incrementer.val must be min or greater")
	}

	if c.orig < c.min {
		return fmt.Errorf("invalid ClampedIncrementer: Incrementer.orig must be min or greater")
	}

	if c.max != 0 && c.val > c.max {
		return fmt.Errorf("invalid ClampedIncrementer: Incrementer.val must min <= val <= max")
	}

	if c.max != 0 && c.orig > c.max {
		return fmt.Errorf("invalid ClampedIncrementer: Incrementer.orig must min <= orig <= max")
	}

	return nil
}
