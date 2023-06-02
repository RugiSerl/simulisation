package Entity

import (
	"github.com/RugiSerl/simulisation/app/Game/material"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/settings"
)

// Par Raphaël
// --------------------------------------------------
// fonctions de détection de collisions
// pour éviter que les entités se regroupe toutes en un même point, elles se repoussent

func (e *Entity) Uncollide(entities []*Entity) {
	if settings.GameSettings.EntitySettings.UncollideAgressive {
		e.UnCollideAgressive(entities)
	} else {
		e.UnCollidePassive(entities)
	}
}

// L'entité balaye toutes les autres entités sur son chemin
func (e *Entity) UnCollideAgressive(entities []*Entity) {

	for _, entity := range entities {
		if entity.ID != e.ID {
			if entity.HitBox.DetectCircleCollision(e.HitBox) && e.HitBox.CenterPosition != entity.HitBox.CenterPosition {
				entity.HitBox.CenterPosition = e.HitBox.CenterPosition.Add(entity.HitBox.CenterPosition.Substract(e.HitBox.CenterPosition).ScaleToNorm(entity.HitBox.Radius + e.HitBox.Radius))
			}
		}

	}
}

// l'entité se déplace lorsqu'elle est en collision avec une autre
func (e *Entity) UnCollidePassive(entities []*Entity) {
	for _, entity := range entities {
		if entity.ID != e.ID {
			if entity.HitBox.DetectCircleCollision(e.HitBox) && e.HitBox.CenterPosition != entity.HitBox.CenterPosition {
				e.HitBox.CenterPosition = entity.HitBox.CenterPosition.Add(e.HitBox.CenterPosition.Substract(entity.HitBox.CenterPosition).ScaleToNorm(entity.HitBox.Radius + e.HitBox.Radius))
			}
		}

	}
}

func (e *Entity) UncollideMaterial(materials []*material.Material) {
	for _, material := range materials {
		if e.HitBox.DetectRectCollision(material.Rect) {
			e.HitBox.CenterPosition = e.HitBox.CenterPosition.Add(graphic.NewVector2(1, 0))
		}
	}

}
