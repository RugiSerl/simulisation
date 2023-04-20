package stats

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	font rl.Font
)

const TEXT_SIZE = 20
const TEXT_SPACING = 0

func InitFont() {
	font = rl.LoadFontEx("assets/VarelaRound-Regular.ttf", TEXT_SIZE, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ:0123456789.- "))
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)
}

func ShowStats(statType string, stat string, position rl.Vector2, horizontalAnchor int8, verticalAnchor int8) graphic.Vector2 {

	sizeRect := rl.MeasureTextEx(font, statType+stat, TEXT_SIZE, TEXT_SPACING)

	position = rl.Vector2(graphic.GetRectCoordinatesWithAnchor(graphic.Vector2(position), horizontalAnchor, verticalAnchor, graphic.Vector2(sizeRect), graphic.GetWindowRect()))

	rl.DrawTextEx(font, statType, position, TEXT_SIZE, TEXT_SPACING, rl.Black)
	position.X += rl.MeasureTextEx(font, statType, TEXT_SIZE, TEXT_SPACING).X

	rl.DrawTextEx(font, stat, position, TEXT_SIZE, TEXT_SPACING, rl.Black)

	return graphic.Vector2(sizeRect)
}
