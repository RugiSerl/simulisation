package app

import (
	"github.com/RugiSerl/simulisation/app/Game"
	"github.com/RugiSerl/simulisation/app/Game/components"
	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/settings"
	"github.com/RugiSerl/simulisation/app/stats"
	"github.com/RugiSerl/simulisation/app/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	myInterface *ui.UserInterface
)

// fonction principale
func Run() {
	load()

	for !rl.WindowShouldClose() {
		update()

	}
	quit()

}

// charge les ressources du jeu
func load() {
	settings.LoadSettings()
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(960, 560, "Simulisation")
	rl.SetWindowIcon(*rl.LoadImage("assets/person.png"))
	rl.SetTargetFPS(60)

	global.MyGame = Game.NewGame() //beaucoup de "game"
	myInterface = ui.NewInterface()

	components.TextureEntite = rl.LoadTexture("assets/person.png")
	rl.SetTextureFilter(components.TextureEntite, rl.FilterBilinear)

	stats.InitFont()

}

// fonction appelée à chaque frame
func update() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.DarkGray)
	global.MyGame.Update()
	myInterface.Update()

	rl.EndDrawing()

}

// gère les instructions à la fermeture du jeu
func quit() {

	rl.CloseWindow()

}
