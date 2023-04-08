package graphic

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Circle struct {
	CenterPosition Vector2
	Radius         float32
}

// Detect circle collision with mouse
func DetectCircleMouseCollision(circleCenterX float64, circleCenterY float64, circleRadius float64) bool {

	x, y := float64(rl.GetMouseX()), float64(rl.GetMouseY())
	distance := math.Sqrt(math.Pow((x-circleCenterX), 2) + math.Pow((y-circleCenterY), 2))
	rl.NewVector2(0, 0)

	return (distance <= circleRadius)

}
