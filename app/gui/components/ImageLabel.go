package components

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// le TextLabel est un objet désignant un texte, qui est statique, et qui est destiné à être affiché.
type ImageLabel struct {
	position graphic.Vector2
	size     graphic.Vector2
	anchorX  int8
	anchorY  int8
	texture  rl.Texture2D
}

// Initialisation du texte, qui est rendu sur une texture qui fait renderer
func NewImageLabel(texture rl.Texture2D, position graphic.Vector2, horizontalAnchor int8, verticalAnchor int8) *ImageLabel {
	i := new(ImageLabel)

	i.position = position
	i.anchorX = horizontalAnchor
	i.anchorY = verticalAnchor

	i.size = graphic.NewVector2(float32(texture.Width), float32(texture.Height))

	i.texture = texture

	return i

}

// Fonctions permettant d'afficher le texte
func (i *ImageLabel) Render(surfaceRect graphic.Rect) {

	physicPosition := graphic.GetRectCoordinatesWithAnchor(i.position, i.anchorX, i.anchorY, i.size, surfaceRect)
	rl.DrawTextureEx(i.texture, rl.Vector2(physicPosition), 0, 1, rl.White)

}

// retourne la taille du label
func (i *ImageLabel) GetSize() graphic.Vector2 {
	return i.size
}
