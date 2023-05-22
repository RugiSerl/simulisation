package stats

// Par Raphaël

import (
	"strconv"

	"github.com/RugiSerl/simulisation/app/Game"
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	font rl.Font
)

const TEXT_SIZE = 20
const TEXT_SPACING = 0

func InitFont() {
	font = rl.LoadFontEx("assets/VarelaRound-Regular.ttf", TEXT_SIZE, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ:0123456789.- ()"))
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)
}

// affiche une statistique du jeu
func ShowStat(statType string, stat string, position rl.Vector2, horizontalAnchor int8, verticalAnchor int8) graphic.Vector2 {

	sizeRect := rl.MeasureTextEx(font, statType+stat, TEXT_SIZE, TEXT_SPACING)

	position = rl.Vector2(graphic.GetRectCoordinatesWithAnchor(graphic.Vector2(position), horizontalAnchor, verticalAnchor, graphic.Vector2(sizeRect), graphic.GetWindowRect()))

	rl.DrawTextEx(font, statType, position, TEXT_SIZE, TEXT_SPACING, rl.Black)
	position.X += rl.MeasureTextEx(font, statType, TEXT_SIZE, TEXT_SPACING).X

	rl.DrawTextEx(font, stat, position, TEXT_SIZE, TEXT_SPACING, rl.Black)

	return graphic.Vector2(sizeRect)
}

// affiche les statistiques du jeu
func ShowStats(game *Game.Game) {
	var temp graphic.Vector2

	temp = ShowStat("FPS : ", strconv.FormatInt(int64(rl.GetFPS()), 10), rl.NewVector2(0, 0), graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	temp = temp.Add(ShowStat("FrameTime (ms) : ", strconv.FormatFloat(float64(rl.GetFrameTime()*1000), 'f', 1, 64), rl.NewVector2(0, temp.Y), graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP))

	ShowStat("Entity amount : ", strconv.FormatInt(int64(game.GetEntityAmount()), 10), rl.NewVector2(0, temp.Y), graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)

	temp = ShowStat("Camera Y : ", strconv.FormatFloat(float64(game.Camera.Target.Y), 'f', 1, 64), rl.NewVector2(0, 0), graphic.ANCHOR_LEFT, graphic.ANCHOR_BOTTOM)
	temp = temp.Add(ShowStat("Camera X : ", strconv.FormatFloat(float64(game.Camera.Target.X), 'f', 1, 64), rl.NewVector2(0, temp.Y), graphic.ANCHOR_LEFT, graphic.ANCHOR_BOTTOM))
	ShowStat("Camera Zoom : ", strconv.FormatFloat(float64(game.Camera.Zoom), 'f', 1, 64), rl.NewVector2(0, temp.Y), graphic.ANCHOR_LEFT, graphic.ANCHOR_BOTTOM)

}
