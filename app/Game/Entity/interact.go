package Entity

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/math"
)

// --------------------------------------------------
// fonction pour faire se reproduire les entités
// les nouvelles cellules sont proches "moralement" de celles qui les ont engendré
func (e *Entity) Reproduce(othersEntities *[]*Entity) {
	var entityClose int32 = 0
	for _, entity := range *othersEntities {
		if entity.HitBox.CenterPosition.Substract(e.HitBox.CenterPosition).GetNorm() < RADIUS_SENSIVITY && entity.ID != e.ID {
			entityClose += 1
		}
	}

	if entityClose > 5 {
		entityClose = 5
	}

	var probability float64 = float64(entityClose) / 1000

	if math.RandomProbability(probability) {
		*othersEntities = append(*othersEntities, NewEntity(e.HitBox.CenterPosition.Add(graphic.NewVector2(0, 1)), len(*othersEntities), uint8(math.RandomRange(int(e.ValeurMorale)-CHILD_MAXIMUM_DIFFERENCE, (int(e.ValeurMorale)+CHILD_MAXIMUM_DIFFERENCE)))))
	}

}
