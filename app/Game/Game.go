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

// Cette fonction est appellée lorsqu'une entité est censée apparaître
func (g *Game) SpawnEntities() {

	var minEcartMoral uint8 = 128
	var minEcartMoralIndex int = 0

	e := gameComponents.NewEntity()

	for i, entity := range g.entities {
		if e.DistanceMorale(entity) < minEcartMoral {
			minEcartMoralIndex = i
		}
		fmt.Println(i)

	}

	g.entities = append(g.entities, e)
	fmt.Println(minEcartMoralIndex)
}
