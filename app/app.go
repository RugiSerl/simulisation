package app

import (
	"strconv"

	"github.com/RugiSerl/simulisation/app/Game"
	"github.com/RugiSerl/simulisation/app/Game/components"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/stats"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	myGame *Game.Game
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
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.InitWindow(800, 450, "Simulisation")
	rl.SetWindowIcon(*rl.LoadImage("assets/person.png"))
	rl.SetTargetFPS(60)
	myGame = Game.NewGame() //beaucoup de "game"

	components.TextureEntite = rl.LoadTexture("assets/person.png")
	rl.SetTextureFilter(components.TextureEntite, rl.FilterBilinear)

	stats.InitFont()

}

// fonction appelée à chaque frame
func update() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.LightGray)
	myGame.Update()

	showStats()

	rl.EndDrawing()

}

// affiche les statistiques du jeu
func showStats() {
	stats.ShowStats("FPS : ", strconv.FormatInt(int64(rl.GetFPS()), 10), rl.NewVector2(0, 0), graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	stats.ShowStats("Entity amount : ", strconv.FormatInt(int64(myGame.GetEntityAmount()), 10), rl.NewVector2(0, 30), graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)

	size := stats.ShowStats("Camera Y : ", strconv.FormatFloat(float64(myGame.Camera.Target.Y), 'f', 1, 64), rl.NewVector2(0, 0), graphic.ANCHOR_LEFT, graphic.ANCHOR_BOTTOM)
	size = size.Add(stats.ShowStats("Camera X : ", strconv.FormatFloat(float64(myGame.Camera.Target.X), 'f', 1, 64), rl.NewVector2(0, size.Y), graphic.ANCHOR_LEFT, graphic.ANCHOR_BOTTOM))
	stats.ShowStats("Camera Zoom : ", strconv.FormatFloat(float64(myGame.Camera.Zoom), 'f', 1, 64), rl.NewVector2(0, size.Y), graphic.ANCHOR_LEFT, graphic.ANCHOR_BOTTOM)

}

// gère les instructions à la fermeture du jeu
func quit() {

	rl.CloseWindow()

}
