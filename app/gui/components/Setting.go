package components

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// abtraction layer for visual setting
type Setting struct {
	label         *Label
	componentType int
	component     interface{}
}

const (
	TYPE_BOOL      = 0
	TYPE_SLIDER    = 1
	TYPE_DROP_MENU = 2
)

func NewSetting(name string, componentType int, font rl.Font, fontSize float32, position graphic.Vector2, horizontalAnchor int8, verticalAnchor int8) *Setting {
	s := new(Setting)
	s.label = Newlabel(name, font, fontSize, position, horizontalAnchor, verticalAnchor)

	labelSize := s.label.GetSize()

	switch componentType {
	case TYPE_BOOL:
		s.component = NewCheckBox(position.Add(graphic.NewVector2(labelSize.X, 0)), false, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	}

	return s
}
