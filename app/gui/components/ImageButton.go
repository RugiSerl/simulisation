package components

// Par Raphaël

import (
	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// l'objet ImageButton désigne un bouton qui a pour hitbox et pour visuel une texture
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

// Initialisation du bouton
func NewImageButton(position graphic.Vector2, texture rl.Texture2D, horizontalAnchor int8, verticalAnchor int8) *ImageButton {
	b := new(ImageButton)

	b.img = texture

	b.position = position
	b.size = graphic.NewVector2(float32(b.img.Width), float32(b.img.Height))

	b.anchorX = horizontalAnchor
	b.anchorY = verticalAnchor

	b.HoverState = false
	b.PressedState = false

	return b
}

// Fonction de mise à jour du bouton
func (b *ImageButton) Update(containingRect graphic.Rect) {
	ImageButtonPhysicalPosition = graphic.GetRectCoordinatesWithAnchor(b.position, b.anchorX, b.anchorY, b.size.Scale(global.InterfaceScale), containingRect)
	b.handleInput()
	b.render()
}

// Fonction qui permet de gérer les inputs du bouton
func (b *ImageButton) handleInput() {
	b.HoverState, b.PressedState = false, false
	if graphic.DetectRectCollision(graphic.GetMouseRect(), graphic.NewRectFromVector(ImageButtonPhysicalPosition, b.size.Scale(global.InterfaceScale))) {
		b.HoverState = true
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			b.PressedState = true
		}
	}

}

// Fonction d'affichage du bouton
func (b *ImageButton) render() {

	if b.HoverState {
		rl.DrawTextureEx(b.img, rl.Vector2(ImageButtonPhysicalPosition), 0, global.InterfaceScale, rl.White)
	} else {
		rl.DrawTextureEx(b.img, rl.Vector2(ImageButtonPhysicalPosition), 0, global.InterfaceScale, rl.NewColor(255, 255, 255, 120))
	}

}
