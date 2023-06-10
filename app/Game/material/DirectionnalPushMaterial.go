package material

import (
	"math"

	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type DirectionnalPushMaterial struct {
	*Material
	Angle float64
}

func NewDirectionnalPushMaterial(materialBase *Material, rect graphic.Rect) *DirectionnalPushMaterial {
	p := new(DirectionnalPushMaterial)
	p.Material = materialBase
	p.rect = rect
	if p.rect.Width/p.rect.Height > 1.0 {
		p.Angle = math.Pi / 2
	} else {
		p.Angle = 0
	}

	return p

}

func (p *DirectionnalPushMaterial) Interact(position graphic.Vector2) graphic.Vector2 {
	return position.Add(position.Substract(p.rect.GetCenter()).FlattenToLine(p.Angle).ScaleToNorm(settings.GameSettings.InteractionSpeed))

}

func (p *DirectionnalPushMaterial) Update() {
	p.rect.Fill(rl.Blue, 0)
	rl.DrawLineV(p.rect.GetCenter().ToRaylibVector2(), p.rect.GetCenter().Add(graphic.NewVectorFromAngle(p.Angle).Scale(5)).ToRaylibVector2(), rl.Red)
}
