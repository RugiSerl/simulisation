package Game

import (
	"github.com/RugiSerl/simulisation/app/Game/Entity"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/math"
	"github.com/RugiSerl/simulisation/app/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Classe qui contient le déroulement principal du jeu
type Game struct {
	entities []*Entity.Entity
	Camera   rl.Camera2D
	Paused   bool
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

	g.entities = []*Entity.Entity{}
	g.Camera = rl.NewCamera2D(rl.NewVector2(0, 0), rl.NewVector2(0, 0), 0, 10)
	g.Paused = false

	return g
}

// Cette fonction est appelée à chaque frame et s'occupe de montrer graphiquement l'état du jeu, ainsi que de mettre à jour les entités
func (g *Game) Update() {
	g.UpdateCamera()

	rl.BeginMode2D(g.Camera)

	g.UpdateEntity()

	if rl.IsKeyPressed(rl.KeySpace) {
		g.Paused = !g.Paused
	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsKeyDown(rl.KeyLeftShift) {
		g.SpawnEntity(g.getMouseWorldCoordinates())
	}
	if rl.IsKeyPressed(rl.KeyLeftControl) {
		settings.GameSettings.VisualSettings.GradientEntities = !settings.GameSettings.VisualSettings.GradientEntities
	}

	rl.EndMode2D()

}

func (g *Game) UpdateEntity() {
	//mise à jour des entités
	for _, entity := range g.entities {
		if entity.Dead == false {
			if !g.Paused {
				entity.Update(&g.entities)
			}

			entity.Render()

			if entity.GetPointCollision(g.getMouseWorldCoordinates()) && rl.IsMouseButtonDown(rl.MouseRightButton) {
				entity.Dead = true

			}

		}
	}

	//suppression des entités "mortes"
	var i int = 0
	for {
		if i >= len(g.entities) {
			break
		}
		if g.entities[i].Dead {
			g.entities = remove(g.entities, i)

		}

		i++
	}
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

	e := Entity.NewEntity(position, len(g.entities), uint8(math.RandomRange(0, 255)))

	g.entities = append(g.entities, e)
}

func (g *Game) GetEntityAmount() int {
	return len(g.entities)
}

func remove(s []*Entity.Entity, i int) []*Entity.Entity {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// transformer les coordonnées physiques de la souris dans la fenêtre en coordonnée virtuelle dans le jeu
func (g *Game) getMouseWorldCoordinates() graphic.Vector2 {
	return graphic.Vector2(rl.GetMousePosition()).Scale(1 / g.Camera.Zoom).Add(graphic.Vector2(g.Camera.Target))

}
