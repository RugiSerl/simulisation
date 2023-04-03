package app

import rl "github.com/gen2brain/raylib-go/raylib"

func Run() {
	load()

	for !rl.WindowShouldClose() {
		rl.DrawText("ceci est un test", 0, 0, 20, rl.Black)

	}
	quit()

}

func load() {

	rl.InitWindow(800, 450, "Simulisation")
	rl.SetTargetFPS(-1)

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
