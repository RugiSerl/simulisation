package graphic

import "math"

//a vector 2 is an object containing two number. It can represent a position on screen, a size, or a movement
type Vector2 struct {
	X float32
	Y float32
}

func NewVector2(x float32, y float32) *Vector2 {
	v := new(Vector2)
	v.X = x
	v.Y = y

	return v
}

func (v *Vector2) GetNorm() float32 {

	return float32(math.Sqrt(math.Pow((float64(v.X)), 2) + math.Pow((float64(v.Y)), 2)))

}

func (v *Vector2) Add(otherVector *Vector2) {
	v.X += otherVector.X
	v.Y += otherVector.Y
}

func (v *Vector2) Substract(otherVector *Vector2) {
	v.X -= otherVector.X
	v.Y -= otherVector.Y
}
