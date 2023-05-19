package graphic

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//objet représentant un cercle, et quelques fonctions pratiques

type Circle struct {
	CenterPosition Vector2
	Radius         float32
}

func NewCircle(radius float32, x float32, y float32) Circle {
	return Circle{Radius: radius, CenterPosition: NewVector2(x, y)}

}

// Fonction permettant de détecter si deux Cercles se chevauchent
func (c *Circle) DetectCircleCollision(otherCircle Circle) bool {

	return (c.CenterPosition.Substract(otherCircle.CenterPosition).GetNorm() <= c.Radius+otherCircle.Radius)
}

func (c *Circle) DetectMouseCollision() bool {
	return c.DetectPointCollision(Vector2(rl.GetMousePosition()))
}

func (c *Circle) DetectPointCollision(position Vector2) bool {
	return (position.Substract(c.CenterPosition).GetNorm() <= c.Radius)

}

func (c *Circle) Fill(color color.RGBA) {
	rl.DrawCircleV(rl.Vector2(c.CenterPosition), c.Radius, color)

}

func (c *Circle) DrawLines(color color.RGBA) {
	rl.DrawCircleLines(int32(c.CenterPosition.X), int32(c.CenterPosition.Y), c.Radius, color)
	//rl.DrawCircleSectorLines(rl.Vector2(c.CenterPosition), c.Radius, 0, 360, 100, color)
}
