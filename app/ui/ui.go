package ui

import (
	"strconv"

	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/stats"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type UserInterface struct {
	testButton string
}

func NewInterface() *UserInterface {

	u := new(UserInterface)

	return u

}

func (u *UserInterface) Update() {
	showStats()

	rect := graphic.NewRect(50, 50, 50, 50)

	rect.Fill(rl.White, 0)

}

// affiche les statistiques du jeu
func showStats() {
	stats.ShowStats("FPS : ", strconv.FormatInt(int64(rl.GetFPS()), 10), rl.NewVector2(0, 0), graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	stats.ShowStats("Entity amount : ", strconv.FormatInt(int64(global.MyGame.GetEntityAmount()), 10), rl.NewVector2(0, 30), graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)

	size := stats.ShowStats("Camera Y : ", strconv.FormatFloat(float64(global.MyGame.Camera.Target.Y), 'f', 1, 64), rl.NewVector2(0, 0), graphic.ANCHOR_LEFT, graphic.ANCHOR_BOTTOM)
	size = size.Add(stats.ShowStats("Camera X : ", strconv.FormatFloat(float64(global.MyGame.Camera.Target.X), 'f', 1, 64), rl.NewVector2(0, size.Y), graphic.ANCHOR_LEFT, graphic.ANCHOR_BOTTOM))
	stats.ShowStats("Camera Zoom : ", strconv.FormatFloat(float64(global.MyGame.Camera.Zoom), 'f', 1, 64), rl.NewVector2(0, size.Y), graphic.ANCHOR_LEFT, graphic.ANCHOR_BOTTOM)

}
