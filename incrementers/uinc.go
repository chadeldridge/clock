package incrementers

import (
	"fmt"
)

type UIncrementer struct {
	Incrementer
}

// NewUIncrementer creates a new counter with a minimum value of 0 and no maximum value.
func NewUIncrementer() UIncrementer {
	return UIncrementer{Incrementer{inc: 1}}
}

// NewUIncrementerWithValue creates a new counter with a starting value of 0 or greater.
func NewUIncrementerWithValue(val int) UIncrementer {
	if val < 0 {
		val = 0
	}

	return UIncrementer{Incrementer{val: val, inc: 1, orig: val}}
}

// NewUIncrementerFromJSON creates a new counter from a JSON representation.
func NewUIncrementerFromJSON(data []byte) (UIncrementer, error) {
	var u UIncrementer
	err := u.UnmarshalJSON(data)
	return u, err
}

// Increment increases the counter by the incrementer value.
func (u *UIncrementer) Increment() { u.val = ClampMin(u.val+u.inc, 0) }

// Decrement decreases the counter by the incrementer value.
func (u *UIncrementer) Decrement() { u.val = ClampMin(u.val-u.inc, 0) }

// Add increases the counter by the given number of val.
func (u *UIncrementer) Add(val int) { u.val = ClampMin(u.val+val, 0) }

// Remove decreases the counter by the given number of val.
func (u *UIncrementer) Remove(val int) { u.val = ClampMin(u.val-val, 0) }

// SetValue sets the counter value to the given number of val with a minimum of 0.
func (u *UIncrementer) SetValue(val int) { u.val = ClampMin(val, 0) }

// SetOrginalValue sets the counter's original value to the given number of val with a minimum of 0.
func (u *UIncrementer) SetOriginalValue(val int) { u.orig = ClampMin(val, 0) }

func (u *UIncrementer) UnmarshalJSON(data []byte) error {
	i, err := NewIncrementerFromJSON(data)
	if err != nil {
		return err
	}

	u.Incrementer = i
	if u.val < 0 {
		return fmt.Errorf("invalid UIncrementer: Incrementer.val must be 0 or greater")
	}

	if u.orig < 0 {
		return fmt.Errorf("invalid UIncrementer: Incrementer.orig must be 0 or greater")
	}

	return nil
}
