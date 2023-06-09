package Entity

// Par Raphaël
import (
	"math"

	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

//--------------------------------------------------
//fonctions de déplacement

func (e *Entity) Move(otherEntities []*Entity) {
	if settings.GameSettings.EntitySettings.GoToClosestNeightbour {
		e.MoveToClosestNeighbour(otherEntities)

	} else {
		e.MoveToWeightedAverage(otherEntities)
	}

}

// Cette fonction permet de déplacer l'entité et de rapprocher l'entité des entités similaires.
// Elle choisit une destination qui est la 'moyenne' des position pondérée à l'aide des 'distances morales'
// Elle ne peut "voir" que les autres entités qui sont dans un certain rayon de cette dernière
func (e *Entity) MoveToClosestNeighbour(otherEntities []*Entity) {

	var min *Entity = nil

	for _, entity := range otherEntities {
		if entity.ID != e.ID {
			// vérification que l'entité est dans le rayon de vision de l'entité
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

// Fonction qui leur permet de tous se regrouper
func (e *Entity) MoveToWeightedAverage(otherEntities []*Entity) {

	var sum graphic.Vector2 = graphic.NewVector2(0, 0)
	var weight float32
	var weightSum float32 = 0

	for _, entity := range otherEntities {
		if entity.ID != e.ID {
			// vérification que l'entité est dans le rayon de vision de l'entité
			if entity.HitBox.CenterPosition.Substract(e.HitBox.CenterPosition).GetNorm() < settings.GameSettings.EntitySettings.RadiusSensivity {
				//weight est le coefficient de la position dans la moyenne
				weight = float32(e.DistanceMorale(entity)) / 255
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

	if e.HitBox.CenterPosition != point {
		e.Facing = float32(float64(point.Substract(e.HitBox.CenterPosition).GetAngle() / (2 * math.Pi) * 360))
	}

	if settings.GameSettings.EntitySettings.LinearMove {
		e.GotoLinear(point)
	} else {
		e.GotoDivide(point)

	}
}

// aller à un point de manière linéaire
func (e *Entity) GotoLinear(point graphic.Vector2) {

	if e.HitBox.CenterPosition.Substract(point).GetNorm() > settings.GameSettings.EntitySettings.Speed*rl.GetFrameTime() {
		e.HitBox.CenterPosition = e.HitBox.CenterPosition.Add(point.Substract(e.HitBox.CenterPosition).ScaleToNorm(settings.GameSettings.EntitySettings.Speed * rl.GetFrameTime()))

	} else {
		e.HitBox.CenterPosition = point
	}

}

// aller à un point en divisant la distance par une certaine valeur
func (e *Entity) GotoDivide(point graphic.Vector2) {
	e.HitBox.CenterPosition = e.HitBox.CenterPosition.Add(point.Substract(e.HitBox.CenterPosition).Scale(.1))
}
