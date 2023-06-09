package Entity

// Par Raphaël et Gaël
import (
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/math"
	"github.com/RugiSerl/simulisation/app/settings"
)

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

	var probability float64 = float64(entityClose) * float64(settings.GameSettings.EntitySettings.BaseProbabilityReproduction)
	if math.RandomProbability(probability) {
		*othersEntities = append(*othersEntities, NewEntity(e.HitBox.CenterPosition.Add(graphic.NewVector2(1, 1)), len(*othersEntities), generateCloseValue(int(e.ValeurMorale), int(settings.GameSettings.EntitySettings.ChildMaximumDifference))))
	}

}

// Fonction qui permet, grâce à une probabilité, de retourner True ou False de manière aléatoire
func generateCloseValue(value int, gap int) uint8 {
	return uint8(math.RandomRange(value-gap, (value + gap)))
}

// fonction permettant à chaque entités d'avoir une chance de tuer une autre entité qui aurait une valeur morale différente
// plus la différence entre les valeurs morales est grande, et plus l'entité a de chance de tuer l'entité concernée
// Par Gaël
func (e *Entity) Kill(othersEntities *[]*Entity) {

	for _, entity := range *othersEntities {
		if entity.HitBox.CenterPosition.Substract(e.HitBox.CenterPosition).GetNorm() < settings.GameSettings.EntitySettings.RadiusSensivity && entity.ID != e.ID {
			var probability float64 = float64(e.DistanceMorale(entity)) * float64(settings.GameSettings.EntitySettings.BaseProbabilityKill)
			if math.RandomProbability(probability) {
				entity.Dead = true
			}
		}
	}

}
