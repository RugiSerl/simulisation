package Entity

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

//--------------------------------------------------
//fonctions de déplacement

// Cette fonction permet de déplacer l'entité et de rapprocher l'entité des entités similaires.
// Elle choisit une destination qui est la 'moyenne' des position pondérée à l'aide des 'distances morales'
// Elle ne peut "voir" que les autres entités qui sont dans un certain rayon de cette dernière (RADIUS_SENSIVITY)
func (e *Entity) MoveToClosestNeighbour(otherEntities []*Entity) {

	var min *Entity = nil

	for _, entity := range otherEntities {
		if entity.ID != e.ID {
			if entity.HitBox.CenterPosition.Substract(e.HitBox.CenterPosition).GetNorm() < settings.GameSettings.EntitySettings.RadiusSensivity {
				if min != nil {
					if entity.DistanceMorale(e) < min.DistanceMorale(e) {
						min = entity
					}
				} else {
					min = entity
				}

			}
		}

	}
	if min != nil {
		e.Goto(min.HitBox.CenterPosition) // déplacement vers cette position

	}

}

func (e *Entity) MoveToWeightedAverage(otherEntities []*Entity) {

	var sum graphic.Vector2 = graphic.NewVector2(0, 0)
	var weight float32
	var weightSum float32 = 0

	for _, entity := range otherEntities {
		if entity.ID != e.ID {
			if entity.HitBox.CenterPosition.Substract(e.HitBox.CenterPosition).GetNorm() < settings.GameSettings.EntitySettings.RadiusSensivity {
				weight = float32(e.DistanceMorale(entity)) / 255
				weight = 1
				weightSum += weight
				sum = sum.Add(entity.HitBox.CenterPosition.Scale(weight))
			}
		}

	}
	if weightSum != 0 { // éviter la division par 0, si jamais l'entité n'a aucune entité dans son rayon RADIUS_SENSIVITY
		average := sum.Scale(1 / weightSum) // division par l'effectif pour faire la moyenne

		e.Goto(average) // déplacement vers cette position

	}

}

// aller à un point
func (e *Entity) Goto(point graphic.Vector2) {
	if settings.GameSettings.EntitySettings.LinearMove {
		e.GotoLinear(point)
	} else {
		e.GotoDivide(point)

	}
}

// aller à un point de manière linéaire
func (e *Entity) GotoLinear(point graphic.Vector2) {

	if e.HitBox.CenterPosition.Substract(point).GetNorm() > SPEED*rl.GetFrameTime() {
		e.HitBox.CenterPosition = e.HitBox.CenterPosition.Add(point.Substract(e.HitBox.CenterPosition).ScaleToNorm(SPEED * rl.GetFrameTime()))

	} else {
		e.HitBox.CenterPosition = point
	}

}

// aller à un point en divisant la distance par une certaine valeur
func (e *Entity) GotoDivide(point graphic.Vector2) {
	e.HitBox.CenterPosition = e.HitBox.CenterPosition.Add(point.Substract(e.HitBox.CenterPosition).Scale(.1))
}
