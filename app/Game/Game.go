package game

import (
	"github.com/RugiSerl/simulisation/app/Game/gameComponents"
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
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

	g.SpawnMultipleEntities(POPULATION_AT_THE_START, graphic.NewVector2(0, 0))

	return g
}

// Cette fonction est appelée à chaque image et s'occupe de montrer graphiquement l'état du jeu, ainsi que de mettre à jour les entités
func (g *Game) Update() {
	for _, entity := range g.entities {
		entity.Update(g.entities)

	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsKeyDown(rl.KeySpace) {
		g.SpawnEntity(graphic.Vector2(rl.GetMousePosition()))
	}

}

// Cette fonction fait apparaître plusieurs entités
func (g *Game) SpawnMultipleEntities(amount int, position graphic.Vector2) {
	for i := 0; i < amount; i++ {
		g.SpawnEntity(position)
	}
}

// Cette fonction est appellée lorsqu'une entité est censée apparaître
func (g *Game) SpawnEntity(position graphic.Vector2) {

	e := gameComponents.NewEntity(position)

	g.entities = append(g.entities, e)
}
