package app

import (
	game "github.com/RugiSerl/simulisation/app/Game"
	"github.com/RugiSerl/simulisation/app/Game/gameComponents"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	myGame *game.Game
)

func Run() {
	load()

	for !rl.WindowShouldClose() {
		update()

	}
	quit()

}

func load() {
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(800, 450, "Simulisation")
	rl.SetWindowIcon(*rl.LoadImage("assets/person.png"))
	rl.SetTargetFPS(-1)
	myGame = game.NewGame() //beaucoup de "game"
	gameComponents.TextureEntite = rl.LoadTexture("assets/person.png")
	rl.SetTextureFilter(gameComponents.TextureEntite, rl.FilterBilinear)

}

func update() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.White)
	myGame.Update()

	rl.EndDrawing()

}

func quit() {

	rl.CloseWindow()

}
