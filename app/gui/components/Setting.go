package components

import (
	"log"

	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// abtraction layer for visual setting
type Setting struct {
	label         *Label
	componentType int

	checkBox *CheckBox
	slider   *Slider

	boolValue  *bool
	floatValue *float32

	size, position   graphic.Vector2
	anchorX, anchorY int8
}

const (
	TYPE_NO_COMPONENT = -1
	TYPE_BOOL         = 0
	TYPE_SLIDER       = 1
	TYPE_DROP_MENU    = 2
)

func NewSetting(name string, componentType int, font rl.Font, fontSize float32, position graphic.Vector2, horizontalAnchor int8, verticalAnchor int8) *Setting {
	s := new(Setting)

	s.position = position
	s.anchorX, s.anchorY = horizontalAnchor, verticalAnchor

	if componentType != TYPE_NO_COMPONENT {
		name += " : "
	}

	s.label = Newlabel(name, font, fontSize, graphic.NewVector2(0, 0), graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	s.size = s.label.GetSize()

	s.componentType = componentType

	switch componentType {
	case TYPE_BOOL:
		s.checkBox = NewCheckBox(graphic.NewVector2(0, 0), graphic.ANCHOR_RIGHT, graphic.ANCHOR_TOP)
		s.size = s.size.Add(graphic.NewVector2(CHECKBOX_SIZE*global.InterfaceScale, 0))

	case TYPE_SLIDER:
		s.slider = NewSlider(graphic.NewVector2(0, 0), graphic.ANCHOR_RIGHT, graphic.ANCHOR_TOP)
		s.size = s.size.Add(graphic.NewVector2(s.slider.GetSize().X, 0))

	case TYPE_NO_COMPONENT:

	default:
		log.Fatal("incorrect component type :", componentType)

	}

	return s
}

func (s *Setting) Update(containingRect graphic.Rect) {
	settingRect := graphic.NewRectFromVector(graphic.GetRectCoordinatesWithAnchor(s.position, s.anchorX, s.anchorY, s.size, containingRect), s.size)

	s.label.Render(settingRect)

	switch s.componentType {
	case TYPE_BOOL:
		s.checkBox.Update(settingRect)
	case TYPE_SLIDER:
		s.slider.Update(settingRect)

	case TYPE_NO_COMPONENT:

	}
}

func (s *Setting) SetBool(value *bool) {
	if s.componentType != TYPE_BOOL {
		incorrectValue()
	}

	s.checkBox.SetValue(value)

}

func incorrectValue() {
	log.Fatal("incorrect value type")
}
