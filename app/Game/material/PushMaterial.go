package material

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/settings"
)

type PushMaterial struct {
	*Material
}

func NewPushMaterial(materialBase *Material, rect graphic.Rect) *PushMaterial {
	p := new(PushMaterial)
	p.Material = materialBase
	p.rect = rect

	return p

}

func (p *PushMaterial) Interact(position graphic.Vector2) graphic.Vector2 {
	return position.Add(position.Substract(p.rect.GetCenter()).ScaleToNorm(settings.GameSettings.InteractionSpeed))

}
