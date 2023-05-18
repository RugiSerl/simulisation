package Game

import (
	"github.com/RugiSerl/simulisation/app/Game/Entity"
	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/gui"
	"github.com/RugiSerl/simulisation/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Classe qui contient le déroulement principal du jeu
type Game struct {
	entities               []*Entity.Entity
	Camera                 rl.Camera2D
	cameraPositionMomentum graphic.Vector2
	cameraZoomMomentum     float32
}

// constante qui définit le nombre d'entités qui apparaîssent lorsque le jeu démarre
const POPULATION_AT_THE_START = 10

// vitesse à laquelle la caméra du jeu se déplace lorsque l'utilisateur appuie sur les flèches directionnelles
const CAMERA_SPEED = 60

// quantité de zoom effectué sur la caméra lorsque l'utilisateur zoom en utilisant la molette de la souris
const CAMERA_ZOOM_AMOUNT = 0.1

var blurShader *graphic.Shader
var textureRender rl.RenderTexture2D

// création d'une instance de Game
func NewGame() *Game {
	g := new(Game)

	g.entities = []*Entity.Entity{}
	g.Camera = rl.NewCamera2D(rl.NewVector2(0, 0), rl.NewVector2(0, 0), 0, 10)
	blurShader = graphic.InitShader("assets/blur.fs")
	var blurAmount float32 = 2
	blurShader.SetValueFromUniformName("size", blurAmount, rl.ShaderUniformFloat)
	textureRender = rl.LoadRenderTexture(1920, 1080)

	return g
}

// Cette fonction est appelée à chaque frame et s'occupe de montrer graphiquement l'état du jeu, ainsi que de mettre à jour les entités
func (g *Game) Update() {

	g.UpdateCamera()
	//tout les éléments du jeu sont rendus sur le renderer
	rl.BeginTextureMode(textureRender)

	rl.ClearBackground(rl.NewColor(194, 187, 186, 255))

	rl.BeginMode2D(g.Camera)

	g.UpdateEntity()

	g.UpdateUserInput()

	rl.EndMode2D()

	rl.EndTextureMode()

	//application du shader de flou, si les paramètres sont ouverts
	if global.SettingsOpen {
		blurShader.Begin()
	}

	rl.DrawTextureRec(textureRender.Texture, rl.NewRectangle(0, 0, float32(textureRender.Texture.Width), float32(-textureRender.Texture.Height)), rl.NewVector2(0, 0), rl.White)

	//fin du shader de flou
	if global.SettingsOpen {
		blurShader.End()
	}
}

// gérer les informations entrées par l'utilisateur
func (g *Game) UpdateUserInput() {
	if (rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsKeyDown(rl.KeyLeftShift)) && (!global.SettingsOpen || rl.GetMousePosition().X < float32(rl.GetScreenWidth())-gui.SETTINGS_WIDTH) {
		g.SpawnEntity(g.getMouseWorldCoordinates())
	}

	if rl.IsMouseButtonDown(rl.MouseMiddleButton) {
		for _, entity := range g.entities {
			entity.Goto(g.getMouseWorldCoordinates())
		}
	}

	if rl.IsKeyPressed(rl.KeyDelete) {
		g.entities = []*Entity.Entity{}
	}

}

// mise à jour des entités
func (g *Game) UpdateEntity() {
	for _, entity := range g.entities {
		if entity.Dead == false {
			if !global.SettingsOpen {
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
	g.cameraPositionMomentum = g.cameraPositionMomentum.Scale(0.7)
	if rl.IsKeyDown(rl.KeyLeft) {
		g.cameraPositionMomentum.X -= CAMERA_SPEED
	}
	if rl.IsKeyDown(rl.KeyRight) {
		g.cameraPositionMomentum.X += CAMERA_SPEED
	}
	if rl.IsKeyDown(rl.KeyUp) {
		g.cameraPositionMomentum.Y -= CAMERA_SPEED
	}
	if rl.IsKeyDown(rl.KeyDown) {
		g.cameraPositionMomentum.Y += CAMERA_SPEED
	}
	// g.cameraMomentum est la vitesse de la caméra, qui augmente lorsque l'utilisateur déplace la caméra, et diminue à chaque frame
	g.Camera.Target = rl.Vector2(graphic.Vector2(g.Camera.Target).Add(g.cameraPositionMomentum.Scale(rl.GetFrameTime())))

	if !global.SettingsOpen || rl.GetMousePosition().X < float32(rl.GetScreenWidth())-gui.SETTINGS_WIDTH {
		//met à jour le zoom de la caméra

		g.cameraZoomMomentum *= 0.8
		g.cameraZoomMomentum += rl.GetMouseWheelMove() * CAMERA_ZOOM_AMOUNT

		g.Camera.Zoom += g.cameraZoomMomentum

		if g.Camera.Zoom < 1 { //1 est le minimum
			g.Camera.Zoom = 1
		}
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

// retourne la quantité d'entités présentes dans le jeu
func (g *Game) GetEntityAmount() int {
	return len(g.entities)
}

// supprime une entité de la liste
func remove(s []*Entity.Entity, i int) []*Entity.Entity {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// transforme les coordonnées physiques de la souris dans la fenêtre en coordonnée virtuelle dans le jeu
func (g *Game) getMouseWorldCoordinates() graphic.Vector2 {
	return graphic.Vector2(rl.GetMousePosition()).Scale(1 / g.Camera.Zoom).Add(graphic.Vector2(g.Camera.Target))

}
