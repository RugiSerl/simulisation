package app

import (
	"strconv"

	game "github.com/RugiSerl/simulisation/app/Game"
	"github.com/RugiSerl/simulisation/app/Game/gameComponents"
	"github.com/RugiSerl/simulisation/app/Game/stats"
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
	stats.InitFont()

}

func update() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.White)
	myGame.Update()

	stats.ShowStats("FPS :", strconv.FormatInt(int64(rl.GetFPS()), 10), rl.NewVector2(0, 0))
	stats.ShowStats("entity amount :", strconv.FormatInt(int64(myGame.GetEntityAmount()), 10), rl.NewVector2(0, 30))

	rl.EndDrawing()

}

func quit() {

	rl.CloseWindow()

}
