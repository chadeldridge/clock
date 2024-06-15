package vectors

import (
	"fmt"
	"math"
)

type vector2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// NewVector2 creates a new Vector2 with the given x and y values.
func NewVector2(x, y float64) vector2 {
	return vector2{X: x, Y: y}
}

func (v vector2) Copy() vector2 {
	return NewVector2(v.X, v.Y)
}

//					//
//		Operators		//
//					//

func (v vector2) Add(s float64) vector2 {
	return NewVector2(v.X+s, v.Y+s)
}

func (v vector2) AddVector(v2 vector2) vector2 {
	return NewVector2(v.X+v2.X, v.Y+v2.Y)
}

func (v vector2) Sub(s float64) vector2 {
	return NewVector2(v.X-s, v.Y-s)
}

func (v vector2) SubVector(v2 vector2) vector2 {
	return NewVector2(v.X-v2.X, v.Y-v2.Y)
}

func (v vector2) Mul(s float64) vector2 {
	return NewVector2(v.X*s, v.Y*s)
}

func (v vector2) MulVector(v2 vector2) vector2 {
	return NewVector2(v.X*v2.X, v.Y*v2.Y)
}

func (v vector2) Div(s float64) vector2 {
	return NewVector2(v.X/s, v.Y/s)
}

func (v vector2) DivVector(v2 vector2) vector2 {
	return NewVector2(v.X/v2.X, v.Y/v2.Y)
}

func (v vector2) Neg() vector2 {
	return NewVector2(-v.X, -v.Y)
}

func (v vector2) Equal(v2 vector2) bool {
	return v.X == v2.X && v.Y == v2.Y
}

//					//
//		Methods			//
//					//

func (v vector2) Angle(v2 vector2) float64 {
	return math.Acos(v.Normalize().Dot(v2.Normalize()))
}

func (v vector2) ClampMagnitude(max float64) vector2 {
	if v.Magnitude() > max {
		return v.Normalize().Mul(max)
	}

	return v.Copy()
}

func (v vector2) Cross(v2 vector2) float64 {
	return v.X*v2.Y - v.Y*v2.X
}

func (v vector2) Distance(v2 vector2) float64 {
	return v.SubVector(v2).Magnitude()
}

func (v vector2) Dot(v2 vector2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v vector2) Lerp(v2 vector2, t float64) vector2 {
	return v.AddVector(v2.SubVector(v).Mul(t))
}

func (v vector2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v vector2) Max(v2 vector2) vector2 {
	return NewVector2(math.Max(v.X, v2.X), math.Max(v.Y, v2.Y))
}

func (v vector2) Min(v2 vector2) vector2 {
	return NewVector2(math.Min(v.X, v2.X), math.Min(v.Y, v2.Y))
}

func (v vector2) Normalize() vector2 {
	if m := v.Magnitude(); m > 0 {
		return v.Div(m)
	}

	return v.Copy()
}

func (v vector2) Vector3() vector3 {
	return NewVector3(v.X, v.Y, 0)
}

func (v vector2) String() string {
	return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}
