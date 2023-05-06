package Entity

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/math"
	"github.com/RugiSerl/simulisation/app/settings"
)

const BASE_PROBABILITY_REPRODUCE = 1e-3

// fonction pour faire se reproduire les entités.
// les nouvelles cellules sont proches "moralement" de celles qui les ont engendré
func (e *Entity) Reproduce(othersEntities *[]*Entity) {
	var entityClose int32 = 0
	for _, entity := range *othersEntities {
		if entity.HitBox.CenterPosition.Substract(e.HitBox.CenterPosition).GetNorm() < settings.GameSettings.EntitySettings.RadiusSensivity && entity.ID != e.ID {
			entityClose += 1
		}
	}

	if entityClose > 5 {
		entityClose = 5
	}

	var probability float64 = float64(entityClose) * BASE_PROBABILITY_REPRODUCE
	if math.RandomProbability(probability) {
		*othersEntities = append(*othersEntities, NewEntity(e.HitBox.CenterPosition.Add(graphic.NewVector2(1, 0)), len(*othersEntities), generateCloseValue(int(e.ValeurMorale), int(settings.GameSettings.EntitySettings.ChildMaximumDifference))))
	}

}

func generateCloseValue(value int, gap int) uint8 {
	return uint8(math.RandomRange(value-gap, (value + gap)))
}
