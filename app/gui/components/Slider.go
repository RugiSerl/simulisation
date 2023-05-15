package components

import (
	"fmt"

	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Slider struct {
	value *float32
	min   float32
	max   float32

	HoverState bool
	position   graphic.Vector2
	anchorX    int8
	anchorY    int8

	size graphic.Vector2
}

var (
	SliderRect graphic.Rect
)

func NewSlider(position graphic.Vector2, horizontalAnchor int8, verticalAnchor int8) *Slider {

	s := new(Slider)

	s.anchorX = horizontalAnchor
	s.anchorY = verticalAnchor

	s.position = position

	s.size = graphic.NewVector2(70, CHECKBOX_SIZE*global.InterfaceScale)

	return s

}

func (s *Slider) SetValue(value *float32, min float32, max float32) {
	s.value, s.min, s.max = value, min, max

}

func (s *Slider) GetSize() graphic.Vector2 {
	return s.size
}

func (s *Slider) Update(containingRect graphic.Rect) {
	SliderRect = graphic.NewRectFromVector(graphic.GetRectCoordinatesWithAnchor(s.position, s.anchorX, s.anchorY, s.size, containingRect), s.size)

	s.handleInput()
	s.render()

}

func (s *Slider) handleInput() {

}

func (s *Slider) render() {
	//SliderRect.Fill(rl.Green, 0)

	bar := graphic.GetInnerHorizontalrect(SliderRect, SliderRect.Height/3)

	bar.Fill(rl.Black, 0)

	fmt.Println(bar)
}
