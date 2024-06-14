package rpgtools

import "fmt"

const incrementer = 1

// UIncrementer is a positive incrementer that can be incremented and decremented between 0 and max.
// You can set the amount to increment by to a postive or negative value. Default is 1.
type UIncrementer struct {
	max int
	val int
	inc int
}

// New creates a new UIncrementer with the given number of max.
func NewUIncrementer(max int) UIncrementer { return UIncrementer{max: max, inc: 1} }

// NewWithValue creates a new UIncrementer with the given number of max and val.
func NewUIncrementerWithValue(max, val int) UIncrementer {
	return UIncrementer{max: max, val: val, inc: 1}
}

// NewFromJSON creates a new UIncrementer from a JSON representation.
func NewUIncrementerFromJSON(data []byte) (UIncrementer, error) {
	var u UIncrementer
	err := u.UnmarshalJSON(data)
	return u, err
}

func (u UIncrementer) getinc() int {
	if u.inc == 0 {
		return incrementer
	}

	return u.inc
}

// Max returns the size of the UIncrementer.
func (u UIncrementer) Max() int { return u.max }

// Value returns the current number of val on the UIncrementer.
func (u UIncrementer) Value() int { return u.val }

// Incrementer returns the current number of incrementer on the UIncrementer.
func (u UIncrementer) Incrementer() int { return u.getinc() }

// IsFull returns true if the UIncrementer is full.
func (u UIncrementer) IsFull() bool { return u.val == u.max }

// IsEmpty returns true if the UIncrementer is empty.
func (u UIncrementer) IsEmpty() bool { return u.val == 0 }

// IncrementBy sets the size to increment by to the given number of inc.
func (u *UIncrementer) IncrementBy(inc int) {
	u.inc = inc
}

// Advance increments the UIncrementer by the incrementer value. Default is 1.
func (u *UIncrementer) Advance() { u.Increment() }

// Increment increments the UIncrementer by the incrementer value. Default is 1.
func (u *UIncrementer) Increment() {
	u.val += u.getinc()
	if u.val >= u.max {
		u.val = u.max
	}
}

// Decrement decrements the UIncrementer by the incrementer value. Default is 1.
func (u *UIncrementer) Decrement() {
	u.val -= u.getinc()
	if u.val <= 0 {
		u.val = 0
	}
}

// Add increments the UIncrementer by the given number of val.
func (u *UIncrementer) Add(val int) {
	if u.val == u.max {
		return
	}

	u.val += val

	if u.val > u.max {
		u.val = u.max
	}
}

// Remove decrements the UIncrementer by the given number of val.
func (u *UIncrementer) Remove(val int) {
	u.val -= val

	if u.val < 0 {
		u.val = 0
	}
}

// SetValue sets the UIncrementer to the given number of val.
func (u *UIncrementer) SetValue(val int) {
	if val < 0 {
		val = 0
	}

	if val > u.max {
		val = u.max
	}

	u.val = val
}

// SetMax sets the size of the UIncrementer to the given number of max.
func (u *UIncrementer) SetMax(max int) {
	if max < 1 {
		max = 1
	}

	u.max = max

	if u.val > u.max {
		u.val = u.max
	}
}

// AddMax increases the size of the UIncrementer by the given number of max.
func (u *UIncrementer) AddMax(max int) {
	u.max += max
}

// RemoveMax decreases the size of the UIncrementer by the given number of max.
func (u *UIncrementer) RemoveMax(max int) {
	u.max -= max

	if u.max < 1 {
		u.max = 1
	}

	if u.val > u.max {
		u.val = u.max
	}
}

// Fill sets the UIncrementer to full.
func (u *UIncrementer) Fill() {
	u.val = u.max
}

// Empty sets the UIncrementer to 0.
func (u *UIncrementer) Empty() {
	u.val = 0
}

// Reset sets the UIncrementer to 0.
func (u *UIncrementer) Reset() {
	u.val = 0
}

// String returns a string representation of the UIncrementer.
func (u UIncrementer) String() string {
	return fmt.Sprintf("%d/%d", u.val, u.max)
}

// MarshalJSON returns a JSON representation of the UIncrementer.
func (u UIncrementer) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"max":%d,"val":%d,"inc":%d}`, u.max, u.val, u.inc)), nil
}

// UnmarshalJSON parses a JSON representation of the UIncrementer.
func (u *UIncrementer) UnmarshalJSON(data []byte) error {
	var max, val, inc int
	_, err := fmt.Sscanf(string(data), `{"max":%d,"val":%d,"inc":%d}`, &max, &val, &inc)
	if err != nil {
		return err
	}

	if max < 1 {
		return fmt.Errorf("invalid UIncrementer: max must be greater than 0")
	}

	if val < 0 {
		return fmt.Errorf("invalid UIncrementer: val must be greater than or equal to 0")
	}

	if val > max {
		return fmt.Errorf("invalid UIncrementer: val must be less than or equal to max")
	}

	if inc < 0 {
		return fmt.Errorf("invalid UIncrementer: inc must be greater than or equal to 0")
	}

	u.max = max
	u.val = val

	return nil
}
