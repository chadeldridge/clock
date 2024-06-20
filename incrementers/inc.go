package incrementers

import (
	"fmt"
)

// Incrementer is a positive incrementer that can be incremented and decremented between 0 and max.
// You can set the amount to increment by to a postive or negative value. Default is 1.
type Incrementer struct {
	inc  int
	val  int
	orig int // Original value when Incrementer was created.
}

// New creates a new Incrementer which will increment a value by 1.
func NewIncrementer() Incrementer { return Incrementer{inc: 1} }

// NewWithValue creates a new Incrementer with initial value of val which will increment value by 1.
func NewIncrementerWithValue(val int) Incrementer {
	return Incrementer{val: val, inc: 1, orig: val}
}

// NewFromJSON creates a new Incrementer from a JSON representation.
func NewIncrementerFromJSON(data []byte) (Incrementer, error) {
	var i Incrementer
	err := i.UnmarshalJSON(data)
	return i, err
}

// Incrementer returns the current incrementer value
func (i Incrementer) Inc() int { return i.inc }

// Value returns the current number of val on the Incrementer.
func (i Incrementer) Value() int { return i.val }

// Original returns the original value of the Incrementer.
func (i Incrementer) Original() int { return i.orig }

// IsEmpty returns true if the Incrementer value is 0.
func (i Incrementer) IsEmpty() bool { return i.val == 0 }

// IsUnchanged returns true if the Incrementer value is the same as the original value.
func (i Incrementer) IsUnchanged() bool { return i.val == i.orig }

// Increment increments the Incrementer by the incrementer value.
func (i *Incrementer) Increment() { i.val += i.inc }

// Decrement decrements the Incrementer by the incrementer value.
func (i *Incrementer) Decrement() { i.val -= i.inc }

// Add increments the Incrementer by the given number of val.
func (i *Incrementer) Add(val int) { i.val += val }

// Remove decrements the Incrementer by the given number of val.
func (i *Incrementer) Remove(val int) { i.val -= val }

// SetIncrementer sets the number the Incrementer will increment by to the given number of inc.
func (i *Incrementer) SetIncrementer(inc int) { i.inc = inc }

// SetValue tries to set the Incrementer value to the given number of val. The value will be clamped.
func (i *Incrementer) SetValue(val int) { i.val = val }

// SetOriginalValue changes the Incrementer's original value to the given number of val.
func (i *Incrementer) SetOriginalValue(val int) { i.orig = val }

// Empty sets the Incrementer value to 0.
func (i *Incrementer) Empty() { i.val = 0 }

// Reset sets the Incrementer to the original value.
func (i *Incrementer) Reset() { i.val = i.orig }

// String returns a string representation of the Incrementer.
func (i Incrementer) String() string { return fmt.Sprintf("%d", i.val) }

// MarshalJSON returns a JSON representation of the Incrementer.
func (i Incrementer) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"inc":%d,"val":%d,"orig":%d}`, i.inc, i.val, i.orig)), nil
}

// UnmarshalJSON parses a JSON representation of the Incrementer.
func (i *Incrementer) UnmarshalJSON(data []byte) error {
	if data == nil {
		return fmt.Errorf("Incrementer.UnmarshalJSON(): data was nil")
	}

	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var inc, val, orig int
	if _, err := fmt.Sscanf(string(data), `{"inc":%d,"val":%d,"orig":%d}`, &inc, &val, &orig); err != nil {
		return err
	}

	i.inc = inc
	i.val = val
	i.orig = orig

	return nil
}
