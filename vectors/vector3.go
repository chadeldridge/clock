package vectors

import "fmt"

type vector3 struct {
	vector2
	Z float64
}

// NewVector3 creates a new Vector3 with the given x, y, and z values.
// Vector3 uses a Vector2 object for x and y so Vector2 fuctions can be used on Vector3.
func NewVector3(x, y, z float64) vector3 {
	return vector3{vector2: vector2{X: x, Y: y}, Z: z}
}

func (v vector3) Copy() vector3 {
	return NewVector3(v.X, v.Y, v.Z)
}

//					//
//		  Math			//
//					//

func (v vector3) Add(s float64) vector3 {
	return NewVector3(v.X+s, v.Y+s, v.Z+s)
}

func (v vector3) AddVector(v2 vector3) vector3 {
	return NewVector3(v.X+v2.X, v.Y+v2.Y, v.Z+v2.Z)
}

func (v vector3) Sub(s float64) vector3 {
	return NewVector3(v.X-s, v.Y-s, v.Z-s)
}

func (v vector3) SubVector(v2 vector3) vector3 {
	return NewVector3(v.X-v2.X, v.Y-v2.Y, v.Z-v2.Z)
}

func (v vector3) Mul(s float64) vector3 {
	return NewVector3(v.X*s, v.Y*s, v.Z*s)
}

func (v vector3) MulVector(v2 vector3) vector3 {
	return NewVector3(v.X*v2.X, v.Y*v2.Y, v.Z*v2.Z)
}

func (v vector3) Div(s float64) vector3 {
	return NewVector3(v.X/s, v.Y/s, v.Z/s)
}

func (v vector3) DivVector(v2 vector3) vector3 {
	return NewVector3(v.X/v2.X, v.Y/v2.Y, v.Z/v2.Z)
}

func (v vector3) Neg() vector3 {
	return NewVector3(-v.X, -v.Y, -v.Z)
}

func (v vector3) Equal(v2 vector3) bool {
	return v.X == v2.X && v.Y == v2.Y && v.Z == v2.Z
}

//					//
//		Methods			//
//					//

func (v vector3) Distance(v2 vector3) float64 {
	return 0
	// return v.Sub(v2).Magnitude()
}

func (v vector3) Vector2() vector2 {
	return v.vector2
}

func (v vector3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X, v.Y, v.Z)
}

func (v vector3) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"x":%f,"y":%f,"z":%f}`, v.X, v.Y, v.Z)), nil
}

func (v *vector3) UnmarshalJSON(data []byte) error {
	var x, y, z float64
	if _, err := fmt.Sscanf(string(data), `{"x":%f,"y":%f,"z":%f}`, &x, &y, &z); err != nil {
		return err
	}

	v.X, v.Y, v.Z = x, y, z
	return nil
}
