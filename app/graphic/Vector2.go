package graphic

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// a vector 2 is an object containing two number. It can represent a position on screen, a size, or a movement
type Vector2 struct {
	X float32
	Y float32
}

func NewVector2(x float32, y float32) Vector2 {
	return Vector2{X: x, Y: y}
}

func (v Vector2) GetNorm() float32 {

	return float32(math.Sqrt(math.Pow((float64(v.X)), 2) + math.Pow((float64(v.Y)), 2)))

}

func (v Vector2) Add(otherVector Vector2) Vector2 {
	return NewVector2(v.X+otherVector.X, v.Y+otherVector.Y)
}

func (v Vector2) Substract(otherVector Vector2) Vector2 {
	return NewVector2(v.X-otherVector.X, v.Y-otherVector.Y)
}

func (v Vector2) Scale(scale float32) Vector2 {
	return NewVector2(v.X*scale, v.Y*scale)
}

func (v Vector2) ScaleToNorm(norm float32) Vector2 {
	return v.Scale(norm / v.GetNorm())
}

func (v Vector2) ToRaylibVector2() rl.Vector2 {
	return rl.NewVector2(v.X, v.Y)

}
