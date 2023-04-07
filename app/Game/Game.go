package game

import (
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

	g.SpawnEntities(POPULATION_AT_THE_START)

	return g
}

// Cette fonction est appellée lorsqu'une entité est censée apparaître
func (g *Game) SpawnEntities() {
	minEcartMoral := 0
	for i, entity := range entities {
		
		
	}
}
