package game

import (
	"github.com/RugiSerl/simulisation/app/Game/gameComponents"
)

// Classe qui contient le déroulement principal du jeu
type Game struct {
	gameMap  *gameComponents.Map
	entities []*gameComponents.Entity
}

const POPULATION_AT_THE_STAART = 10

func NewGame() *Game {
	g := new(Game)

	g.entities = []*gameComponents.Entity{}
	g.gameMap = gameComponents.NewMap()

	g.SpawnEntities(POPULATION_AT_THE_STAART)

	return g
}

func (g *Game) SpawnEntities(amount int) {
	for i := 0; i < amount; i++ {
		g.entities = append(g.entities, gameComponents.NewEntity())

	}

}
