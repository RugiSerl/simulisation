package components

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const DURATION = 1

var (
	time     float32 = -1
	text     string
	textSize graphic.Vector2
	font     rl.Font
	fontSize float32 = 20
)

func InitFont() {
	font = rl.LoadFontEx("assets/VarelaRound-Regular.ttf", int32(fontSize), []rune("é!èabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ:0123456789.- ()"))
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)
}

func NewNotificationText(notifText string) {
	text = notifText
	time = 0
	textSize = graphic.Vector2(rl.MeasureTextEx(font, text, fontSize, 0))

}

func UpdateNotification() {
	time += rl.GetFrameTime()
	if time > 0 && time < DURATION {
		rect := graphic.NewRect(0, 0, textSize.X, textSize.Y)
		rect.Fill(rl.Black, 0.3)
		rl.DrawTextEx(font, text, rl.NewVector2(0, 0), fontSize, 0, rl.White)
	}

}
