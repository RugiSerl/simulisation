package game

import (
	"github.com/RugiSerl/simulisation/app/Game/gameComponents"
)

// Classe qui contient le déroulement principal du jeu
type Game struct {
	gameMap  *gameComponents.Map
	entities []*gameComponents.Entity
}

func newGame() *Game {
	g := new(Game)

	g.entities = []*gameComponents.Entity{}
	g.gameMap = gameComponents.NewMap()

	return g
}

func (g *Game) SpawnEntities(amount int) {
	for i := 0; i < amount; i++ {
		g.entities = append(g.entities, gameComponents.NewEntity())

	}

}
