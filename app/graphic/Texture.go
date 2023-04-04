package graphic

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Draw a raylib texture but the x and the y values are the center of the texture
func DrawtextureFromCenter(texture rl.Texture2D, x float32, y float32, rotation float32, scale float32, tint color.RGBA) {
	rl.DrawTextureEx(texture, rl.NewVector2(x-float32(texture.Width)*scale/2, y-float32(texture.Height)*scale/2), rotation, scale, tint)
}

// draw image ajusted to the window
func DrawAdjustedTexture(texture rl.Texture2D, tint color.RGBA) {

	var imagex, imagey, scale, imageW, imageH float32

	winW, winH := float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())
	winRatio := winW / winH

	imgW, imgH := float32(texture.Width), float32(texture.Height)
	imgRatio := imgW / imgH

	if imgRatio > winRatio { //change Height
		imageW = winW
		imageH = imageW / imgRatio

		imagex = 0
		imagey = winH/2 - imageH/2

		scale = imageH / imgH

	} else if imgRatio < winRatio {
		imageH = winH
		imageW = imageH * imgRatio

		imagey = 0
		imagex = winW/2 - imageW/2

		scale = imageW / imgW

	} else {
		imagex = 0
		imagey = 0
		scale = winW / imgW
	}

	rl.DrawTextureEx(texture, rl.NewVector2(imagex, imagey), 0, scale, tint)

}

//draw render texture correctly, else opengl draw it upside down
func DrawTextureRenderer(renderTexture rl.RenderTexture2D, destRect Rect, tint color.RGBA) {

	rl.DrawTexturePro(renderTexture.Texture, rl.NewRectangle(0, 0, float32(renderTexture.Texture.Width), float32(-renderTexture.Texture.Height)), destRect.ToRaylibRect(), rl.NewVector2(0, 0), 0, tint)
}
