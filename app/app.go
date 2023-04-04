package app

import (
	game "github.com/RugiSerl/simulisation/app/Game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Game *game.Game
)

func Run() {
	load()

	for !rl.WindowShouldClose() {
		update()

	}
	quit()

}

func load() {

	rl.InitWindow(800, 450, "Simulisation")
	rl.SetTargetFPS(-1)
	Game = game.NewGame() //beaucoup de "game"

}

func update() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Blue)
	rl.DrawText("On n'a pas encore commencé", 190, 200, 20, rl.Red)

	rl.EndDrawing()

}

func quit() {

	rl.CloseWindow()

}
