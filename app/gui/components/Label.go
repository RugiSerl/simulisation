package components

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// le TextLabel est un objet désignant un texte, qui est statique, et qui est destiné à être affiché.
type TextLabel struct {
	position graphic.Vector2
	size     graphic.Vector2
	anchorX  int8
	anchorY  int8
	texture  rl.RenderTexture2D
}

// Initialisation du texte, qui est rendu sur une texture qui fait renderer
func Newlabel(text string, font rl.Font, fontSize float32, position graphic.Vector2, horizontalAnchor int8, verticalAnchor int8) *TextLabel {
	l := new(TextLabel)

	l.position = position
	l.anchorX = horizontalAnchor
	l.anchorY = verticalAnchor

	l.size = graphic.Vector2(rl.MeasureTextEx(font, text, fontSize, 0))

	l.texture = rl.LoadRenderTexture(int32(l.size.X), int32(l.size.Y))
	rl.SetTextureFilter(l.texture.Texture, rl.FilterBilinear)

	rl.BeginTextureMode(l.texture)

	rl.DrawTextEx(font, text, rl.NewVector2(0, 0), fontSize, 0, rl.Black)

	rl.EndTextureMode()

	return l

}

// Fonctions permettant d'afficher le texte
func (l *TextLabel) Render(surfaceRect graphic.Rect) {

	physicPosition := graphic.GetRectCoordinatesWithAnchor(l.position, l.anchorX, l.anchorY, l.size, surfaceRect)
	rl.DrawTextureRec(l.texture.Texture, rl.NewRectangle(0, 0, l.size.X, -l.size.Y), rl.Vector2(physicPosition), rl.White)

}

// retourne la taille du label
func (l *TextLabel) GetSize() graphic.Vector2 {
	return l.size
}
