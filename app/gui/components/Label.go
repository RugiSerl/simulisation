package components

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Label struct {
	position graphic.Vector2
	anchorX  int8
	anchorY  int8
	texture  rl.RenderTexture2D
}

func Newlabel(text string, font rl.Font, position graphic.Vector2, horizontalAnchor int8, verticalAnchor int8) {
	l := new(Label)

	l.position = position
	l.anchorX = horizontalAnchor
	l.anchorY = verticalAnchor

}
