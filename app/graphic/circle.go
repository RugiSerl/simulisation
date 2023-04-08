package graphic

type Circle struct {
	CenterPosition *Vector2
	Radius         float32
}

// detect if two circles overlap
func (c *Circle) DetectCircleCollision(otherCircle *Circle) bool {
	distance := c.CenterPosition
	distance.Substract(otherCircle.CenterPosition)

	return (distance.GetNorm() <= c.Radius+otherCircle.Radius)

}
