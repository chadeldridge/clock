package vectors

import "fmt"

type Vector3 struct {
	Vector2
	Z float64
}

// NewVector3 creates a new Vector3 with the given x, y, and z values.
// Vector3 uses a Vector2 object for x and y so Vector2 fuctions can be used on Vector3.
func NewVector3(x, y, z float64) Vector3 {
	return Vector3{Vector2: Vector2{X: x, Y: y}, Z: z}
}

func (v Vector3) Copy() Vector3 {
	return NewVector3(v.X, v.Y, v.Z)
}

//					//
//		  Math			//
//					//

func (v Vector3) Add(s float64) Vector3 {
	return NewVector3(v.X+s, v.Y+s, v.Z+s)
}

func (v Vector3) AddVector(v2 Vector3) Vector3 {
	return NewVector3(v.X+v2.X, v.Y+v2.Y, v.Z+v2.Z)
}

func (v Vector3) Sub(s float64) Vector3 {
	return NewVector3(v.X-s, v.Y-s, v.Z-s)
}

func (v Vector3) SubVector(v2 Vector3) Vector3 {
	return NewVector3(v.X-v2.X, v.Y-v2.Y, v.Z-v2.Z)
}

func (v Vector3) Mul(s float64) Vector3 {
	return NewVector3(v.X*s, v.Y*s, v.Z*s)
}

func (v Vector3) MulVector(v2 Vector3) Vector3 {
	return NewVector3(v.X*v2.X, v.Y*v2.Y, v.Z*v2.Z)
}

func (v Vector3) Div(s float64) Vector3 {
	return NewVector3(v.X/s, v.Y/s, v.Z/s)
}

func (v Vector3) DivVector(v2 Vector3) Vector3 {
	return NewVector3(v.X/v2.X, v.Y/v2.Y, v.Z/v2.Z)
}

func (v Vector3) Neg() Vector3 {
	return NewVector3(-v.X, -v.Y, -v.Z)
}

func (v Vector3) Equal(v2 Vector3) bool {
	return v.X == v2.X && v.Y == v2.Y && v.Z == v2.Z
}

//					//
//		Methods			//
//					//

func (v Vector3) Distance(v2 Vector3) float64 {
	return 0
	// return v.Sub(v2).Magnitude()
}

func (v Vector3) ToVector2() Vector2 {
	return v.Vector2.Copy()
}

func (v Vector3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X, v.Y, v.Z)
}

func (v Vector3) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"x":%f,"y":%f,"z":%f}`, v.X, v.Y, v.Z)), nil
}

func (v *Vector3) UnmarshalJSON(data []byte) error {
	var x, y, z float64
	if _, err := fmt.Sscanf(string(data), `{"x":%f,"y":%f,"z":%f}`, &x, &y, &z); err != nil {
		return err
	}

	v.X, v.Y, v.Z = x, y, z
	return nil
}
