package app

// Par Raphaël

import (
	"github.com/RugiSerl/simulisation/app/Game"
	"github.com/RugiSerl/simulisation/app/Game/Entity"
	"github.com/RugiSerl/simulisation/app/gui"
	"github.com/RugiSerl/simulisation/app/gui/components"
	"github.com/RugiSerl/simulisation/app/settings"
	"github.com/RugiSerl/simulisation/app/stats"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	myInterface *gui.UserInterface
	myGame      *Game.Game
	Background  rl.Texture2D
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
	rl.SetTargetFPS(int32(settings.GameSettings.VisualSettings.MaxFps))
	rl.SetExitKey(rl.KeyBackspace)

	myGame = Game.NewGame()
	myInterface = gui.NewInterface()

	Entity.TextureEntite = rl.LoadTexture("assets/person.png")
	rl.SetTextureFilter(Entity.TextureEntite, rl.FilterBilinear)

	stats.InitFont()

	components.InitFont()

}

// fonction appelée à chaque frame
func update() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.DarkGray)

	myGame.Update()

	myInterface.Update()

	if settings.GameSettings.VisualSettings.DisplayStats {
		stats.ShowStats(myGame)

	}

	components.UpdateNotification()

	rl.EndDrawing()

	if rl.IsKeyPressed(rl.KeyF11) {
		rl.SetWindowSize(rl.GetMonitorWidth(rl.GetCurrentMonitor()), rl.GetMonitorHeight(rl.GetCurrentMonitor()))
		rl.ToggleFullscreen()
	}

}

// gère les instructions à la fermeture du jeu
func quit() {

	rl.CloseWindow()

}
