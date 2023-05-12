package components

import (
	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Label struct {
	position graphic.Vector2
	size     graphic.Vector2
	anchorX  int8
	anchorY  int8
	texture  rl.RenderTexture2D
}

func Newlabel(text string, font rl.Font, fontSize float32, position graphic.Vector2, horizontalAnchor int8, verticalAnchor int8) *Label {
	l := new(Label)

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

func (l *Label) Render(surfaceRect graphic.Rect) {

	physicPosition := graphic.GetRectCoordinatesWithAnchor(l.position, l.anchorX, l.anchorY, l.size.Scale(global.InterfaceScale), surfaceRect)
	rl.DrawTextureRec(l.texture.Texture, rl.NewRectangle(0, 0, l.size.X*global.InterfaceScale, -l.size.Y*global.InterfaceScale), rl.Vector2(physicPosition), rl.White)

}

func (l *Label) GetSize() graphic.Vector2 {
	return l.size
}
