package graphic

// Par Raphaël

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// objet contenant un vecteur qui a deux dimensions, qui peut représenter une position 2d, un déplacement 2d, et quelques méthodes pratiques
type Vector2 struct {
	X float32
	Y float32
}

func NewVector2(x float32, y float32) Vector2 {
	return Vector2{X: x, Y: y}
}
func NewVectorFromAngle(angle float64) Vector2 {
	cos := float32(math.Cos(angle))
	sin := float32(math.Sin(angle))
	return NewVector2(cos, sin)

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

func (v Vector2) GetAngle() float64 {
	if v.Y > 0 {
		return math.Acos(float64(v.ScaleToNorm(1).X))
	} else {
		return -math.Acos(float64(v.ScaleToNorm(1).X))
	}
}

func (v Vector2) Rotate(angle float64) Vector2 {
	return NewVectorFromAngle(v.GetAngle() + angle).ScaleToNorm(v.GetNorm())
}

func (v Vector2) FlattenToLine(lineAngle float64) Vector2 {
	return NewVector2(v.Rotate(lineAngle).X, 0).Rotate(-lineAngle)
}

func (v Vector2) ToRaylibVector2() rl.Vector2 {
	return rl.NewVector2(v.X, v.Y)

}

/*
func init() {
	v := NewVector2(-1, -1)
	fmt.Println(v.GetAngle())
	fmt.Println(NewVector2(-1, 1).GetAngle())
	v = v.Rotate(math.Pi / 2)
	fmt.Println(v)

	fmt.Println(v.GetAngle())

	log.Fatal("test")
}*/
