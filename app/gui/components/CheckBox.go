package components

import (
	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CheckBox struct {
	Value      bool
	HoverState bool
	position   graphic.Vector2
	anchorX    int8
	anchorY    int8
}

var (
	CheckBoxRect graphic.Rect
)

const (
	SIZE = 120
)

func NewCheckBox(position graphic.Vector2, defaultValue bool, horizontalAnchor int8, verticalAnchor int8) *CheckBox {

	c := new(CheckBox)

	c.anchorX = horizontalAnchor
	c.anchorY = verticalAnchor

	c.Value = defaultValue

	c.position = position

	return c

}

func (c *CheckBox) Update(containingRect graphic.Rect) {
	CheckBoxRect = graphic.NewRectFromVector(graphic.GetRectCoordinatesWithAnchor(c.position, c.anchorX, c.anchorY, graphic.NewVector2(SIZE*global.InterfaceScale, SIZE*global.InterfaceScale), containingRect), graphic.NewVector2(SIZE*global.InterfaceScale, SIZE*global.InterfaceScale))

	c.handleInput()
	c.render()

}

func (c *CheckBox) handleInput() {

	c.HoverState = false

	if graphic.DetectRectCollision(CheckBoxRect, graphic.GetMouseRect()) {
		c.HoverState = true
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			c.Value = !c.Value
		}

	}

}

func (c *CheckBox) render() {
	innerRect := graphic.GetInnerRect(CheckBoxRect, 4)

	CheckBoxRect.Fill(rl.Black, 0)
	innerRect.Fill(rl.White, 0)

	if c.Value {
		innerRectConfirmation := graphic.GetInnerRect(innerRect, 3)
		innerRectConfirmation.Fill(rl.Black, 0)
	}

}
