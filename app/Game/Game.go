package Game

// Par Raphaël
import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/RugiSerl/simulisation/app/Game/Entity"
	"github.com/RugiSerl/simulisation/app/Game/material"
	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/gui"
	"github.com/RugiSerl/simulisation/app/gui/components"
	"github.com/RugiSerl/simulisation/app/math"
	"github.com/RugiSerl/simulisation/app/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Classe qui contient le déroulement principal du jeu
type Game struct {
	entities               []*Entity.Entity
	materials              []*material.Material
	Camera                 rl.Camera2D
	cameraPositionMomentum graphic.Vector2
	cameraZoomMomentum     float32
	saveLoadPanel          *SaveLoadPanel

	materialSpawnPreview       bool
	materialSpawnPreviewOrigin graphic.Vector2
}

// constante qui définit le nombre d'entités qui apparaîssent lorsque le jeu démarre
const POPULATION_AT_THE_START = 10

// vitesse à laquelle la caméra du jeu se déplace lorsque l'utilisateur appuie sur les flèches directionnelles
const CAMERA_SPEED = 200

// quantité de zoom effectué sur la caméra lorsque l'utilisateur zoom en utilisant la molette de la souris
const CAMERA_ZOOM_AMOUNT = 5e-2

const SAVE_FILENAME = "save.txt"

var blurShader *graphic.Shader
var textureRender rl.RenderTexture2D

var (
	blurAmount float32

	Background rl.Texture2D
)

// création d'une instance de Game
func NewGame() *Game {
	g := new(Game)

	g.entities = []*Entity.Entity{}
	g.Camera = rl.NewCamera2D(rl.NewVector2(0, 0), rl.NewVector2(0, 0), 0, 10)
	blurShader = graphic.InitShader("assets/blur.fs")
	blurAmount = 4
	blurShader.SetValueFromUniformName("size", blurAmount, rl.ShaderUniformFloat)
	textureRender = rl.LoadRenderTexture(1920, 1080)

	Background = rl.LoadTexture("assets/background.png")
	rl.SetTextureFilter(Background, rl.FilterBilinear)

	g.materials = []*material.Material{}

	g.saveLoadPanel = NewSaveLoadPanel()
	g.materialSpawnPreview = false

	return g
}

// Cette fonction est appelée à chaque frame et s'occupe de montrer graphiquement l'état du jeu, ainsi que de mettre à jour les entités
func (g *Game) Update() {

	g.UpdateCamera()
	//tout les éléments du jeu sont rendus sur le renderer
	rl.BeginTextureMode(textureRender)

	if settings.GameSettings.VisualSettings.ClearBackground || rl.IsKeyPressed(rl.KeyBackspace) {
		rl.ClearBackground(rl.NewColor(0, 0, 0, 0))
		if settings.GameSettings.VisualSettings.Background {
			g.drawBackground()

		}
	}

	rl.BeginMode2D(g.Camera)

	g.UpdateMaterials()

	g.UpdateEntity()

	g.UpdateUserInput()

	rl.EndMode2D()

	rl.EndTextureMode()

	//application du shader de flou, si les paramètres sont ouverts
	if global.SettingsOpen {
		blurShader.SetValueFromUniformName("size", rl.Clamp(blurAmount*gui.AnimationTime/gui.ANIMATION_DURATION, 0, blurAmount), rl.ShaderUniformFloat)
		blurShader.Begin()
	}

	rl.DrawTextureRec(textureRender.Texture, rl.NewRectangle(0, 0, float32(textureRender.Texture.Width), float32(-textureRender.Texture.Height)), rl.NewVector2(0, 0), rl.White)

	//fin du shader de flou
	if global.SettingsOpen {
		blurShader.End()
	}
	g.saveLoadPanel.Update()

}

// gérer les informations entrées par l'utilisateur
func (g *Game) UpdateUserInput() {
	if (rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsKeyDown(rl.KeyLeftShift)) && (!global.SettingsOpen || rl.GetMousePosition().X < float32(rl.GetScreenWidth())-gui.SETTINGS_MENU_WIDTH) && !graphic.DetectRectCollision(g.saveLoadPanel.containingRect, graphic.GetMouseRect()) {

		if settings.GameSettings.UserInputSettings.SpawnMaterial {
			g.materialSpawnPreview = true
			g.materialSpawnPreviewOrigin = g.getMouseWorldCoordinates()

		} else {
			g.SpawnEntity(g.getMouseWorldCoordinates())

		}
	}
	//end preview mode of material spawning
	if g.materialSpawnPreview && rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		g.materialSpawnPreview = false
		// spawn the said material as well
		g.SpawnMaterial()

	}

	if g.materialSpawnPreview {
		g.PreviewMaterialSpawn()
	}

	if rl.IsMouseButtonDown(rl.MouseMiddleButton) {
		for _, entity := range g.entities {
			entity.Goto(g.getMouseWorldCoordinates())
		}
	}

	if rl.IsKeyPressed(rl.KeyDelete) {
		g.entities = []*Entity.Entity{}

	}

	if g.saveLoadPanel.SaveButton.PressedState {
		g.Save()
	}

	if g.saveLoadPanel.LoadButton.PressedState {
		g.Load()
	}

	if rl.IsMouseButtonDown(rl.MouseRightButton) {
		circle := graphic.NewCircle(settings.GameSettings.UserInputSettings.DeleteRadius, g.getMouseWorldCoordinates().X, g.getMouseWorldCoordinates().Y)
		circle.DrawCross(rl.Red)
	}
}

// mise à jour des entités
func (g *Game) UpdateEntity() {

	//affiche les zones de "vision" des entités
	if settings.GameSettings.VisualSettings.DisplaySensibilityZone {
		for _, entity := range g.entities {
			if !entity.Dead {
				entity.RenderSensibilityZone()
			}
		}
	}

	//mise à jour des positions et affichage des entités
	for _, entity := range g.entities {
		if entity.Dead == false {
			if !global.SettingsOpen && !(settings.GameSettings.UserInputSettings.SpawnMaterial && rl.IsMouseButtonDown(rl.MouseLeftButton)) {
				entity.Update(&g.entities, g.materials)
			}

			entity.Render()

			if entity.HitBox.CenterPosition.Substract(g.getMouseWorldCoordinates()).GetNorm() < settings.GameSettings.UserInputSettings.DeleteRadius && rl.IsMouseButtonDown(rl.MouseRightButton) {
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

func (g *Game) UpdateMaterials() {
	for _, material := range g.materials {
		material.Update()
	}
}

// met à jour la caméra pour visualiser le jeu et appliquer les transformations de cette dernière
func (g *Game) UpdateCamera() {

	//déplacement éventuel de la caméra
	g.cameraPositionMomentum = g.cameraPositionMomentum.Scale(0.7)
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		g.cameraPositionMomentum.X -= CAMERA_SPEED
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		g.cameraPositionMomentum.X += CAMERA_SPEED
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		g.cameraPositionMomentum.Y -= CAMERA_SPEED
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		g.cameraPositionMomentum.Y += CAMERA_SPEED
	}
	// g.cameraMomentum est la vitesse de la caméra, qui augmente lorsque l'utilisateur déplace la caméra, et diminue à chaque frame
	g.Camera.Target = rl.Vector2(graphic.Vector2(g.Camera.Target).Add(g.cameraPositionMomentum.Scale(rl.GetFrameTime() / g.Camera.Zoom)))

	//décalage de la caméra, pour que la cible, c'est-à-dire les coordonnées de la caméra, se trouve au milieu de l'écran
	g.Camera.Offset = rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2)

	if !global.SettingsOpen || rl.GetMousePosition().X < float32(rl.GetScreenWidth())-gui.SETTINGS_MENU_WIDTH {
		//met à jour le zoom de la caméra

		g.cameraZoomMomentum *= 0.8
		g.cameraZoomMomentum += rl.GetMouseWheelMove() * CAMERA_ZOOM_AMOUNT * g.Camera.Zoom

		g.Camera.Zoom += g.cameraZoomMomentum

		if g.Camera.Zoom < 1 { //1 est le minimum
			g.Camera.Zoom = 1
		}
	}

}

// affiche le fond d'écran du jeu
func (g *Game) drawBackground() {
	rl.DrawTexturePro(Background, rl.NewRectangle(0, 0, float32(Background.Width), float32(Background.Height)), rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())), rl.NewVector2(0, 0), 0, rl.White)
}

// Cette fonction est appellée lorsqu'une entité est censée apparaître
func (g *Game) SpawnEntity(position graphic.Vector2) {

	var e *Entity.Entity

	if settings.GameSettings.UserInputSettings.SpawnRandomValeurMorale {
		e = Entity.NewEntity(position, len(g.entities), uint8(math.RandomRange(0, 255)))

	} else {
		e = Entity.NewEntity(position, len(g.entities), uint8(settings.GameSettings.UserInputSettings.EntityValeurMoraleOnSpawn))
	}

	g.entities = append(g.entities, e)
}

func (g *Game) SpawnMaterial() {
	g.materials = append(g.materials, material.NewMaterial(g.getMaterialRect()))
	fmt.Println("a")

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
	return graphic.Vector2(rl.GetMousePosition()).Substract(graphic.Vector2(g.Camera.Offset)).Scale(1 / g.Camera.Zoom).Add(graphic.Vector2(g.Camera.Target))
}

func (g *Game) PreviewMaterialSpawn() {
	rect := g.getMaterialRect()
	rect.Fill(rl.DarkBlue, 0)

}
func (g *Game) getMaterialRect() graphic.Rect {
	return graphic.NewRectFromVector(g.materialSpawnPreviewOrigin, g.getMouseWorldCoordinates().Substract(g.materialSpawnPreviewOrigin))
}

func (g *Game) Save() {

	entitiesJson, err := json.MarshalIndent(g.entities, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(SAVE_FILENAME, entitiesJson, 0644)

	if err != nil {
		panic(err)

	} else {
		fmt.Println("jeu sauvegardé avec succès !")
		components.NewNotificationText("jeu sauvegardé avec succès !")
	}

}

func (g *Game) Load() {
	file, err := os.ReadFile(SAVE_FILENAME)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("aucune sauvegarde à charger")

		} else {
			panic(err)
		}

	} else {
		err2 := json.Unmarshal(file, &g.entities)
		if err2 != nil {
			fmt.Println("impossible de charger la sauvegarde", err2)

		} else {
			fmt.Println("jeu chargé avec succès !")
			components.NewNotificationText("jeu chargé avec succès !")
		}
	}
}
