package graphic

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

func (v *Vector2) GetDistance() {

}
