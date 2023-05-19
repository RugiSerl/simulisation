package components

import (
	"log"

	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// l'objet Setting est un objet qui représente un paramètre à afficher graphiquement.
// Il joue le rôle d'interface et permet ainsi de simplifer la création de l'interface faite manuellement
type Setting struct {
	label         *TextLabel
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

// Initialisation d'un paramètre
func NewSetting(name string, componentType int, font rl.Font, fontSize float32, position graphic.Vector2, horizontalAnchor int8, verticalAnchor int8) *Setting {
	s := new(Setting)

	s.position = position
	s.anchorX, s.anchorY = horizontalAnchor, verticalAnchor

	if componentType != TYPE_NO_COMPONENT {
		name += " : "
	} else {
		fontSize += 4
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

// Fonction qui met à jour le paramètre
func (s *Setting) Update(containingRect graphic.Rect) {
	settingRect := graphic.NewRectFromVector(graphic.GetRectCoordinatesWithAnchor(s.position, s.anchorX, s.anchorY, s.size, containingRect), s.size)

	if graphic.DetectRectCollision(settingRect, graphic.GetMouseRect()) {
		settingRect.Fill(rl.NewColor(128, 128, 128, 32), 0)
	}

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

func (s *Setting) SetSliderValue(value *float32, min float32, max float32) {
	if s.componentType != TYPE_SLIDER {
		incorrectValue()
	}

	s.slider.SetValue(value, min, max)

}

func incorrectValue() {
	log.Fatal("incorrect value type")
}
