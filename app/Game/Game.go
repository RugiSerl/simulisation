package game

import (
	"fmt"

	"github.com/RugiSerl/simulisation/app/Game/gameComponents"
)

// Classe qui contient le déroulement principal du jeu
type Game struct {
	gameMap  *gameComponents.Map
	entities []*gameComponents.Entity
}

const POPULATION_AT_THE_START = 10

func NewGame() *Game {
	g := new(Game)

	g.entities = []*gameComponents.Entity{}
	g.gameMap = gameComponents.NewMap()

	g.SpawnEntities()
	g.SpawnEntities()

	return g
}

func (g *Game) SpawnMultipleEntities(amount int) {
	for i := 0; i < amount; i++ {
		g.SpawnEntities()
	}
}

// Cette fonction est appellée lorsqu'une entité est censée apparaître
func (g *Game) SpawnEntities() {

	var minEcartMoral uint8 = 128
	var minEcartMoralIndex int = -1

	e := gameComponents.NewEntity()

	for i, entity := range g.entities {
		if e.DistanceMorale(entity) < minEcartMoral && entity.NbVoisins < gameComponents.NB_VOISINS_MAX {
			minEcartMoralIndex = i
		}
		fmt.Println(i)

	}

	if minEcartMoralIndex != -1 { // l'entité peut aussi naître sans aucun voisin, ex: la toute première entité

	}

	g.entities = append(g.entities, e)
	fmt.Println(minEcartMoralIndex)
}
