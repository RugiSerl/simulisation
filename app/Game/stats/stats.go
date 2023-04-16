package stats

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	font rl.Font
)

const TEXT_SIZE = 30
const TEXT_SPACING = 0

func InitFont() {
	font = rl.LoadFontEx("assets/VarelaRound-Regular.ttf", TEXT_SIZE, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ:0123456789 "))
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)
}

func ShowStats(statType string, stat string, position rl.Vector2) {
	rl.DrawTextEx(font, statType, position, TEXT_SIZE, TEXT_SPACING, rl.Black)
	position.X += rl.MeasureTextEx(font, statType, TEXT_SIZE, TEXT_SPACING).X
	rl.DrawTextEx(font, stat, position, TEXT_SIZE, TEXT_SPACING, rl.Black)

}
