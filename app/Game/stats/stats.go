package stats

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	font rl.Font
)

func InitFont() {
	font = rl.LoadFont("assets/VarelaRound-Regular.ttf")
}

func ShowStats(statType string, stat string, position rl.Vector2) {
	rl.DrawTextEx(font, statType, position, 30, 0, rl.Black)

}
