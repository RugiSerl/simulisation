package ui

import (
	"strconv"

	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/settings"
	"github.com/RugiSerl/simulisation/app/stats"
	"github.com/RugiSerl/simulisation/app/ui/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// largeur du panneau des paramètres
const SETTINGS_WIDTH = 200

// durée en secondes de l'animation lorsque l'utilisateur ouvre la fenêtre de dialogue des paramètres
const ANIMATION_DURATION = 0.1

type UserInterface struct {
	AnimationTime float32

	pauseButton *components.ImageButton
}

func NewInterface() *UserInterface {

	u := new(UserInterface)

	u.pauseButton = components.NewImageButton(graphic.NewVector2(0, 0), rl.LoadTexture("assets/pause.png"), graphic.ANCHOR_HORIZONTAL_MiDDLE, graphic.ANCHOR_BOTTOM)

	return u

}

func (u *UserInterface) Update() {

	if global.SettingsOpen {
		u.UpdateSettings()

	}

	u.pauseButton.Update()

	if u.pauseButton.PressedState {
		settings.GamePaused = !settings.GamePaused
	}

	if rl.IsKeyPressed(rl.KeyS) {
		u.AnimationTime = 0
		global.SettingsOpen = !global.SettingsOpen

	}
	if settings.GameSettings.VisualSettings.DisplayStats {
		u.showStats()

	}

}

// met à jour et affiche la fenêtre des paramètres
func (u *UserInterface) UpdateSettings() {

	size := graphic.NewVector2(SETTINGS_WIDTH, float32(rl.GetScreenHeight()))

	position := graphic.GetRectCoordinatesWithAnchor(graphic.NewVector2(0, 0), graphic.ANCHOR_RIGHT, graphic.ANCHOR_TOP, size, graphic.GetWindowRect())

	//déplace pour l'animation
	if u.AnimationTime < ANIMATION_DURATION {
		position = position.Add(graphic.NewVector2((size.X)*(ANIMATION_DURATION-u.AnimationTime)/ANIMATION_DURATION, 0))

	}

	rect := graphic.NewRectFromVector(position, size)

	rect.Fill(rl.White, 0.1)

	u.AnimationTime += rl.GetFrameTime()

}

// affiche les statistiques du jeu
func (u *UserInterface) showStats() {
	var temp graphic.Vector2

	temp = stats.ShowStats("FPS : ", strconv.FormatInt(int64(rl.GetFPS()), 10), rl.NewVector2(0, 0), graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	temp = temp.Add(stats.ShowStats("FrameTime (ms) : ", strconv.FormatFloat(float64(rl.GetFrameTime()*1000), 'f', 1, 64), rl.NewVector2(0, temp.Y), graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP))

	stats.ShowStats("Entity amount : ", strconv.FormatInt(int64(global.MyGame.GetEntityAmount()), 10), rl.NewVector2(0, temp.Y), graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)

	temp = stats.ShowStats("Camera Y : ", strconv.FormatFloat(float64(global.MyGame.Camera.Target.Y), 'f', 1, 64), rl.NewVector2(0, 0), graphic.ANCHOR_LEFT, graphic.ANCHOR_BOTTOM)
	temp = temp.Add(stats.ShowStats("Camera X : ", strconv.FormatFloat(float64(global.MyGame.Camera.Target.X), 'f', 1, 64), rl.NewVector2(0, temp.Y), graphic.ANCHOR_LEFT, graphic.ANCHOR_BOTTOM))
	stats.ShowStats("Camera Zoom : ", strconv.FormatFloat(float64(global.MyGame.Camera.Zoom), 'f', 1, 64), rl.NewVector2(0, temp.Y), graphic.ANCHOR_LEFT, graphic.ANCHOR_BOTTOM)

}
