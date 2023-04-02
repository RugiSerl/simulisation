package app

import "github.com/gen2brain/raylib-go/raylib"

func Run() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("On n'a pas encore commencé", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}