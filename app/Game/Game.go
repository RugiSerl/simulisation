package Game

import (
	"github.com/RugiSerl/simulisation/app/Game/components"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Classe qui contient le déroulement principal du jeu
type Game struct {
	entities []*components.Entity
	Camera   rl.Camera2D
}

// constante qui définit le nombre d'entités qui apparaîssent lorsque le jeu démarre
const POPULATION_AT_THE_START = 10

// vitesse à laquelle la caméra du jeu se déplace lorsque l'utilisateur appuie sur les flèches directionnelles
const CAMERA_SPEED = 100

// quantité de zoom effectué sur la caméra lorsque l'utilisateur zoom en utilisant la molette de la souris
const CAMERA_ZOOM_AMOUNT = 0.2

// création d'une instance de Game
func NewGame() *Game {
	g := new(Game)

	g.entities = []*components.Entity{}
	g.Camera = rl.NewCamera2D(rl.NewVector2(0, 0), rl.NewVector2(0, 0), 0, 10)

	return g
}

// Cette fonction est appelée à chaque frame et s'occupe de montrer graphiquement l'état du jeu, ainsi que de mettre à jour les entités
func (g *Game) Update() {
	g.UpdateCamera()

	rl.BeginMode2D(g.Camera)

	for _, entity := range g.entities {
		if entity.Dead == false {
			entity.Update(&g.entities)
		}
	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsKeyDown(rl.KeySpace) {
		g.SpawnEntity(graphic.Vector2(rl.GetMousePosition()).Scale(1 / g.Camera.Zoom).Add(graphic.Vector2(g.Camera.Target)))
	}
	if rl.IsKeyPressed(rl.KeyLeftControl) {
		components.ShowValeurMorale = !components.ShowValeurMorale
	}

	rl.EndMode2D()

}

// met à jour la caméra pour visualiser le jeu et appliquer les transformations de cette dernière
func (g *Game) UpdateCamera() {

	//déplacement éventuel de la caméra
	if rl.IsKeyDown(rl.KeyLeft) {
		g.Camera.Target.X -= CAMERA_SPEED * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyRight) {
		g.Camera.Target.X += CAMERA_SPEED * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyUp) {
		g.Camera.Target.Y -= CAMERA_SPEED * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyDown) {
		g.Camera.Target.Y += CAMERA_SPEED * rl.GetFrameTime()
	}

	//met à jour le zoom de la caméra
	g.Camera.Zoom += rl.GetMouseWheelMove() * CAMERA_ZOOM_AMOUNT
	if g.Camera.Zoom < 1 { //1 est le minimum
		g.Camera.Zoom = 1
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

	e := components.NewEntity(position, len(g.entities), uint8(math.RandomRange(0, 255)))

	g.entities = append(g.entities, e)
}

func (g *Game) GetEntityAmount() int {
	return len(g.entities)
}
