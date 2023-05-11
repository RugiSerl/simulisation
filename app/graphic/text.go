package graphic

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	textTexture rl.Texture2D
)

func DrawText(font rl.Font, text string, x float32, y float32, size float32, spacing float32, color color.RGBA) {
	//draw standard text

	rl.DrawTextEx(font, text, rl.NewVector2(x, y), size, spacing, color)

}

func DrawShadowedTextLight(font rl.Font, text string, position Vector2, size float32, spacing float32, color color.RGBA, shadowOffset Vector2, shadowSize float32) {

	rl.DrawTextEx(font, text, rl.Vector2(position.Add(shadowOffset)), size, spacing, rl.NewColor(0, 0, 0, 196))

	rl.DrawTextEx(font, text, rl.Vector2(position), size, spacing, color)
}
