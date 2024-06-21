package vectors

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// NewVector2 creates a new Vector2 with the given x and y values.
func NewVector2(x, y float64) Vector2 {
	return Vector2{X: x, Y: y}
}

func (v Vector2) Copy() Vector2 {
	return NewVector2(v.X, v.Y)
}

//					//
//		Operators		//
//					//

func (v Vector2) Add(s float64) Vector2 {
	return NewVector2(v.X+s, v.Y+s)
}

func (v Vector2) AddVector(v2 Vector2) Vector2 {
	return NewVector2(v.X+v2.X, v.Y+v2.Y)
}

func (v Vector2) Sub(s float64) Vector2 {
	return NewVector2(v.X-s, v.Y-s)
}

func (v Vector2) SubVector(v2 Vector2) Vector2 {
	return NewVector2(v.X-v2.X, v.Y-v2.Y)
}

func (v Vector2) Mul(s float64) Vector2 {
	return NewVector2(v.X*s, v.Y*s)
}

func (v Vector2) MulVector(v2 Vector2) Vector2 {
	return NewVector2(v.X*v2.X, v.Y*v2.Y)
}

func (v Vector2) Div(s float64) Vector2 {
	return NewVector2(v.X/s, v.Y/s)
}

func (v Vector2) DivVector(v2 Vector2) Vector2 {
	return NewVector2(v.X/v2.X, v.Y/v2.Y)
}

func (v Vector2) Neg() Vector2 {
	return NewVector2(-v.X, -v.Y)
}

func (v Vector2) Equal(v2 Vector2) bool {
	return v.X == v2.X && v.Y == v2.Y
}

//					//
//		Methods			//
//					//

func (v Vector2) Angle(v2 Vector2) float64 {
	return math.Acos(v.Normalize().Dot(v2.Normalize()))
}

func (v Vector2) ClampMagnitude(max float64) Vector2 {
	if v.Magnitude() > max {
		return v.Normalize().Mul(max)
	}

	return v.Copy()
}

func (v Vector2) Cross(v2 Vector2) float64 {
	return v.X*v2.Y - v.Y*v2.X
}

func (v Vector2) Distance(v2 Vector2) float64 {
	return v.SubVector(v2).Magnitude()
}

func (v Vector2) Dot(v2 Vector2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vector2) Lerp(v2 Vector2, t float64) Vector2 {
	return v.AddVector(v2.SubVector(v).Mul(t))
}

func (v Vector2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2) Max(v2 Vector2) Vector2 {
	return NewVector2(math.Max(v.X, v2.X), math.Max(v.Y, v2.Y))
}

func (v Vector2) Min(v2 Vector2) Vector2 {
	return NewVector2(math.Min(v.X, v2.X), math.Min(v.Y, v2.Y))
}

func (v Vector2) Normalize() Vector2 {
	if m := v.Magnitude(); m > 0 {
		return v.Div(m)
	}

	return v.Copy()
}

func (v Vector2) Vector3() Vector3 {
	return NewVector3(v.X, v.Y, 0)
}

func (v Vector2) String() string {
	return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}
