package game

import (
	"fmt"

	"github.com/RugiSerl/simulisation/app/Game/gameComponents"
	"github.com/RugiSerl/simulisation/app/graphic"
)

// Classe qui contient le déroulement principal du jeu
type Game struct {
	gameMap  *gameComponents.Map
	entities []*gameComponents.Entity
}

// constante qui définit le nombre d'entités qui apparaîssent lorsque le jeu démarre
const POPULATION_AT_THE_START = 10

// création d'une instance de Game
func NewGame() *Game {
	g := new(Game)

	g.entities = []*gameComponents.Entity{}
	g.gameMap = gameComponents.NewMap()

	g.SpawnMultipleEntities(POPULATION_AT_THE_START)

	return g
}

// Cette fonction est appelée à chaque image et s'occupe de montrer graphiquement l'état du jeu, ainsi que de mettre à jour les entités
func (g *Game) Update() {
	for _, entity := range g.entities {
		entity.Update()

	}

}

// Cette fonction fait apparaître plusieurs entités
func (g *Game) SpawnMultipleEntities(amount int) {
	for i := 0; i < amount; i++ {
		g.SpawnEntity()
		fmt.Println(len(g.entities[i].Voisins))
	}
}

// Cette fonction est appellée lorsqu'une entité est censée apparaître
func (g *Game) SpawnEntity() {

	var minEcartMoral uint8 = 128
	var minEcartMoralIndex int = -1

	e := gameComponents.NewEntity()

	for i, entity := range g.entities {
		if e.DistanceMorale(entity) < minEcartMoral && len(entity.Voisins) < gameComponents.NB_VOISINS_MAX {
			minEcartMoralIndex = i
			minEcartMoral = e.DistanceMorale(entity)
		}

	}

	if minEcartMoralIndex != -1 { // rattachement de l'entité créee à l'entité voisine la plus "similaire"
		e.NouveauLien(g.entities[minEcartMoralIndex])

	} else { // l'entité peut aussi naître sans aucun voisin, ex: la toute première entité, et il faut donc lui rajouter une position dans le jeu
		e.Position = graphic.NewVector2(0, 0)
	}

	g.entities = append(g.entities, e)
}
