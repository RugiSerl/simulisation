package material

import (
	"fmt"
	"math"

	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RotateMaterial struct {
	*Material
	Angle float64
}

func NewRotateMaterial(materialBase *Material, rect graphic.Rect) *RotateMaterial {
	r := new(RotateMaterial)
	r.Material = materialBase
	r.rect = rect

	r.Angle = -21 * math.Pi / 40

	return r

}

func (r *RotateMaterial) Interact(position graphic.Vector2) graphic.Vector2 {
	fmt.Println(r.rect.GetCenter(), position)
	return position.Add(position.Substract(r.rect.GetCenter()).ScaleToNorm(1).Rotate(r.Angle))

}

func (r *RotateMaterial) Update() {
	r.rect.Fill(rl.Yellow, 0)
	rl.DrawLineV(r.rect.GetCenter().ToRaylibVector2(), r.rect.GetCenter().Add(graphic.NewVectorFromAngle(r.Angle).Scale(5)).ToRaylibVector2(), rl.Red)
}
