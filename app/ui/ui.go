package ui

import (
	"strconv"

	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/stats"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// largeur du panneau des paramètres
const SETTINGS_WIDTH = 200

type UserInterface struct {
	testButton string
}

func NewInterface() *UserInterface {

	u := new(UserInterface)

	return u

}

func (u *UserInterface) Update() {
	u.UpdateSettings()
	u.showStats()

}

func (u *UserInterface) UpdateSettings() {
	size := graphic.NewVector2(SETTINGS_WIDTH, float32(rl.GetScreenHeight()))

	position := graphic.GetRectCoordinatesWithAnchor(graphic.NewVector2(0, 0), graphic.ANCHOR_RIGHT, graphic.ANCHOR_TOP, size, graphic.GetWindowRect())

	rect := graphic.NewRectFromVector(position, size)

	rect.Fill(rl.White, 0.2)

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
