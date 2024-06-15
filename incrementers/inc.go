package incrementers

import (
	"fmt"
)

// Incrementer is a positive incrementer that can be incremented and decremented between 0 and max.
// You can set the amount to increment by to a postive or negative value. Default is 1.
type Incrementer struct {
	min  int
	max  int
	val  int
	inc  int
	orig int // Original value when Incrementer was created.
}

// New creates a new Incrementer with the given number of max.
func NewIncrementer(min, max int) Incrementer { return Incrementer{min: min, max: max, inc: 1} }

// NewWithValue creates a new Incrementer with the given number of max and val.
func NewIncrementerWithValue(min, max, val int) Incrementer {
	return Incrementer{min: min, max: max, val: val, inc: 1, orig: val}
}

// NewFromJSON creates a new Incrementer from a JSON representation.
func NewIncrementerFromJSON(data []byte) (Incrementer, error) {
	var i Incrementer
	err := i.UnmarshalJSON(data)
	return i, err
}

// Min returns the min size of the Incrementer.
func (i Incrementer) Min() int { return i.min }

// Max returns the max size of the Incrementer.
func (i Incrementer) Max() int { return i.max }

// Value returns the current number of val on the Incrementer.
func (i Incrementer) Value() int { return i.val }

// Incrementer returns the current incrementer value
func (i Incrementer) Inc() int { return i.inc }

// Original returns the original value of the Incrementer.
func (i Incrementer) Original() int { return i.orig }

// IsFull returns true if the Incrementer is full.
func (i Incrementer) IsFull() bool { return i.val == i.max }

// IsEmpty returns true if the Incrementer is empty.
func (i Incrementer) IsEmpty() bool { return i.val == 0 }

// IsMin returns true if the Incrementer is at the minimum value.
func (i Incrementer) IsMin() bool { return i.val == i.min }

// IncrementBy sets the size to increment by to the given number of inc.
func (i *Incrementer) IncrementBy(inc int) {
	i.inc = inc
}

// Advance increments the Incrementer by the incrementer value.
func (i *Incrementer) Advance() { i.Increment() }

// Increment increments the Incrementer by the incrementer value.
func (i *Incrementer) Increment() { i.val = Clamp(i.val+i.inc, i.min, i.max) }

// Decrement decrements the Incrementer by the incrementer value.
func (i *Incrementer) Decrement() { i.val = Clamp(i.val-i.inc, i.min, i.max) }

// Add increments the Incrementer by the given number of val.
func (i *Incrementer) Add(val int) { i.val = Clamp(i.val+val, i.min, i.max) }

// Remove decrements the Incrementer by the given number of val.
func (i *Incrementer) Remove(val int) { i.val = Clamp(i.val-val, i.min, i.max) }

// SetMin sets the minimum size of the Incrementer to the given number of min.
func (i *Incrementer) SetMin(min int) {
	if min >= i.max {
		min = i.max - 1
	}

	i.min = min
	i.val = ClampMin(i.val, i.min)
}

// SetMax sets the maximum size of the Incrementer to the given number of max.
func (i *Incrementer) SetMax(max int) {
	if max <= i.min {
		max = i.min + 1
	}

	i.max = max
	i.val = ClampMax(i.val, i.max)
}

// SetValue tries to set the Incrementer value to the given number of val. The value will be clamped.
func (i *Incrementer) SetValue(val int) { i.val = Clamp(val, i.min, i.max) }

// SetIncrementer sets the number the Incrementer will increment by to the given number of inc.
func (i *Incrementer) SetIncrementer(inc int) { i.inc = inc }

// SetOriginalValue changes the Incrementer's original value to the given number of val.
// The value will be clamped.
func (i *Incrementer) SetOriginalValue(val int) { i.orig = Clamp(val, i.min, i.max) }

// Fill sets the Incrementer value to the maximum size.
func (i *Incrementer) Fill() { i.val = i.max }

// Empty tries to sets the Incrementer value to 0. The value will be clamped.
func (i *Incrementer) Empty() { i.val = Clamp(0, i.min, i.max) }

// Floor sets the Incrementer to the minimum size.
func (i *Incrementer) Floor() { i.val = i.min }

// Reset sets the Incrementer to 0.
func (i *Incrementer) Reset() {
	i.val = i.orig
	i.val = Clamp(i.val, i.min, i.max)
}

// String returns a string representation of the Incrementer.
func (i Incrementer) String() string {
	return fmt.Sprintf("%d/%d", i.val, i.max)
}

// MarshalJSON returns a JSON representation of the Incrementer.
func (i Incrementer) MarshalJSON() ([]byte, error) {
	return []byte(
			fmt.Sprintf(
				`{"min":%d,"max":%d,"val":%d,"inc":%d,"orig":%d}`,
				i.min, i.max, i.val, i.inc, i.orig,
			),
		),
		nil
}

// UnmarshalJSON parses a JSON representation of the Incrementer.
func (i *Incrementer) UnmarshalJSON(data []byte) error {
	var min, max, val, inc, o int
	if _, err := fmt.Sscanf(
		string(data),
		`{"min":%d,"max":%d,"val":%d,"inc":%d,"orig":%d}`,
		&min, &max, &val, &inc, &o,
	); err != nil {
		return err
	}

	if max <= min {
		return fmt.Errorf("invalid Incrementer: max must be greater than min")
	}

	if val < min {
		return fmt.Errorf("invalid Incrementer: val must be greater than or equal to min")
	}

	if val > max {
		return fmt.Errorf("invalid Incrementer: val must be less than or equal to max")
	}

	i.min = min
	i.max = max
	i.val = val
	i.inc = inc
	i.orig = o

	return nil
}
