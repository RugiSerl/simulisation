package components

import (
	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ImageButton struct {
	position     graphic.Vector2
	size         graphic.Vector2
	img          rl.Texture2D
	HoverState   bool
	PressedState bool
	anchorX      int8
	anchorY      int8
}

var (
	ImageButtonPhysicalPosition graphic.Vector2
)

func NewImageButton(position graphic.Vector2, texture rl.Texture2D, horizontalAnchor int8, verticalAnchor int8) *ImageButton {
	b := new(ImageButton)

	b.img = texture

	b.position = position
	b.size = graphic.NewVector2(float32(b.img.Width), float32(b.img.Height))

	b.anchorX = horizontalAnchor
	b.anchorY = verticalAnchor

	b.HoverState = true
	b.PressedState = true

	return b
}

func (b *ImageButton) Update() {
	ImageButtonPhysicalPosition = graphic.GetRectCoordinatesWithAnchor(b.position, b.anchorX, b.anchorY, b.size.Scale(global.InterfaceScale), graphic.GetWindowRect())
	b.handleInput()
	b.render()
}

func (b *ImageButton) handleInput() {
	b.HoverState, b.PressedState = false, false
	if graphic.DetectRectCollision(graphic.GetMouseRect(), graphic.NewRectFromVector(ImageButtonPhysicalPosition, b.size.Scale(global.InterfaceScale))) {
		b.HoverState = true
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			b.PressedState = true
		}
	}

}

func (b *ImageButton) render() {

	if b.HoverState {
		rl.DrawTextureEx(b.img, rl.Vector2(ImageButtonPhysicalPosition), 0, global.InterfaceScale, rl.White)
	} else {
		rl.DrawTextureEx(b.img, rl.Vector2(ImageButtonPhysicalPosition), 0, global.InterfaceScale, rl.NewColor(255, 255, 255, 120))
	}

}
