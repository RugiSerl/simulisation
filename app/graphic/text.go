package graphic

import (
	"image/color"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	textTexture rl.Texture2D
)

const GRADIENT_WIDTH = 50

func DrawText(font rl.Font, text string, x float32, y float32, size float32, spacing float32, color color.RGBA) {
	//draw standard text

	rl.DrawTextEx(font, text, rl.NewVector2(x, y), size, spacing, color)

}

func DrawShadowedTextLight(font rl.Font, text string, x float32, y float32, size float32, spacing float32, color color.RGBA, shadowOffsetx float32, shadowOffsety float32, shadowSize float32) {

	DrawText(font, text, x+shadowOffsetx, y+shadowOffsety, size, spacing, rl.NewColor(0, 0, 0, 196))

	DrawText(font, text, x, y, size, spacing, color)

}
func DrawShadowedText(font rl.Font, text string, x float32, y float32, size float32, spacing float32, color color.RGBA, shadowOffsetx float32, shadowOffsety float32, shadowSize float32, shadowQuality uint8) float32 {
	//draw text with shadow. not optimised/does not look good for now

	baseX := x

	for _, c := range text {

		letter := string(c)

		var layerSize float32 = size
		var layerAlpha float32

		glyphModel := rl.MeasureTextEx(font, letter, size, spacing)

		centerX := x + glyphModel.X/2 + shadowOffsetx
		centerY := y + glyphModel.Y/2 + shadowOffsety

		for i := 0; uint8(i) < shadowQuality; i++ {

			layerSize = size + float32(i+1)/float32(shadowQuality)*(shadowSize)

			layerAlpha = (float32(shadowQuality) - float32(i)) / float32(shadowQuality) * 255

			glyphModel = rl.MeasureTextEx(font, letter, layerSize, spacing)

			DrawText(font, letter, float32(centerX-glyphModel.X/2), float32(centerY-glyphModel.Y/2), layerSize, spacing, rl.NewColor(0, 0, 0, uint8(layerAlpha)))

		}

		DrawText(font, letter, x, y, size, spacing, color)
		x += glyphModel.X - size/10
	}
	return x - baseX

}

func DrawJustifiedText(font rl.Font, text string, basex float32, xmax float32, basey float32, size float32, spacing float32, color color.RGBA) float32 {
	//draw text and return to line when it exceed a certain x
	words := strings.Split(text, " ")
	x := basex
	y := basey
	GlyphHeight := rl.MeasureTextEx(font, "A", size, spacing).Y

	for i, word := range words {
		var textWidth float32 = rl.MeasureTextEx(font, word+" ", size, spacing).X

		if x+textWidth <= xmax || i == 0 {
			DrawShadowedTextLight(font, word+" ", x, y, size, spacing, color, 2, 2, size)
			x += textWidth

		} else {
			y += GlyphHeight
			x = basex
			DrawShadowedTextLight(font, word+" ", x, y, size, spacing, color, 2, 2, size)
			x += textWidth
		}

		if strings.Contains(word, "\n") { //handle correctly newlines in string
			DrawShadowedTextLight(font, strings.ReplaceAll(word, "\n", "")+" ", x, y, size, spacing, color, 2, 2, size)
			y += GlyphHeight * float32(strings.Count(word, "\n")+1)
		}

	}

	textHeight := y - basey + GlyphHeight

	return textHeight

}

func DrawTextWithAlphaBlending(font rl.Font, text string, textX float32, textY float32, size float32, spacing float32, portionShown float32, color color.RGBA) {
	//extremely unoptimized (raylib might not be optimal for this kind of things)
	//Big memory eater do not use

	rl.UnloadTexture(textTexture) //remove this and your RAM goes brrrr

	textImg := rl.ImageTextEx(font, text, size, spacing, color)
	colorArray := rl.LoadImageColors(textImg)

	var (
		x int32
		y int32
	)

	var alpha uint8

	for x = textImg.Width - 1; x > textImg.Width-GRADIENT_WIDTH && x > 0; x -= 1 {
		for y = 0; y < textImg.Height; y += 1 {
			pixelColor := colorArray[y*textImg.Width+x]

			if pixelColor.R != 0 || pixelColor.G != 0 || pixelColor.B != 0 {
				alpha = uint8(float32(textImg.Width-x) / GRADIENT_WIDTH * 255)
			} else {
				alpha = 0
			}

			rl.ImageDrawPixel(textImg, x, y, rl.NewColor(pixelColor.R, pixelColor.G, pixelColor.B, alpha))
		}
	}

	textTexture = rl.LoadTextureFromImage(textImg)
	rl.UnloadImage(textImg)
	rl.DrawTextureEx(textTexture, rl.NewVector2(textX, textY), 0, 1, rl.White)

}
