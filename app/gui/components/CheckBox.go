package components

import (
	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CheckBox struct {
	value      *bool
	HoverState bool
	position   graphic.Vector2
	anchorX    int8
	anchorY    int8
}

var (
	CheckBoxRect graphic.Rect
)

const (
	CHECKBOX_SIZE = 100
)

func NewCheckBox(position graphic.Vector2, horizontalAnchor int8, verticalAnchor int8) *CheckBox {

	c := new(CheckBox)

	c.anchorX = horizontalAnchor
	c.anchorY = verticalAnchor

	c.position = position

	return c

}

func (c *CheckBox) SetValue(value *bool) {
	c.value = value

}

func (c *CheckBox) Update(containingRect graphic.Rect) {
	CheckBoxRect = graphic.NewRectFromVector(graphic.GetRectCoordinatesWithAnchor(c.position, c.anchorX, c.anchorY, graphic.NewVector2(CHECKBOX_SIZE*global.InterfaceScale, CHECKBOX_SIZE*global.InterfaceScale), containingRect), graphic.NewVector2(CHECKBOX_SIZE*global.InterfaceScale, CHECKBOX_SIZE*global.InterfaceScale))

	c.handleInput()
	c.render()

}

func (c *CheckBox) handleInput() {

	c.HoverState = false

	if graphic.DetectRectCollision(CheckBoxRect, graphic.GetMouseRect()) {
		c.HoverState = true
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

			*c.value = !*c.value

		}

	}

}

func (c *CheckBox) render() {
	innerRect := graphic.GetInnerRect(CheckBoxRect, 2)

	CheckBoxRect.Fill(rl.Black, 0)
	innerRect.Fill(rl.White, 0)

	if *c.value {
		innerRectConfirmation := graphic.GetInnerRect(innerRect, 3)
		innerRectConfirmation.Fill(rl.Black, 0)
	}

}
