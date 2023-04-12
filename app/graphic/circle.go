package graphic

type Circle struct {
	CenterPosition Vector2
	Radius         float32
}

// detect if two circles overlap
func (c *Circle) DetectCircleCollision(otherCircle Circle) bool {

	return (c.CenterPosition.Substract(otherCircle.CenterPosition).GetNorm() <= c.Radius+otherCircle.Radius)

}

func NewCircle(radius float32, x float32, y float32) Circle {
	return Circle{Radius: radius, CenterPosition: NewVector2(x, y)}

}
