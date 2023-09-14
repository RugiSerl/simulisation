package material

import (
	"fmt"
	"math"

	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/settings"
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
	return position.Add(position.Substract(r.rect.GetCenter()).ScaleToNorm(settings.GameSettings.InteractionSpeed).Rotate(r.Angle))

}

func (r *RotateMaterial) Render() {
	r.rect.Fill(rl.Yellow, 0)
}
