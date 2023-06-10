package material

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MaterialType int8

const (
	PUSH_MATERIAL MaterialType = iota
	PUSH_MATERIAL_DIRECTION
	ROTATE_MATERIAL
)

var MATERIAL_TYPE_NAME = []string{"matériau repoussant", "matériau repoussant directionnel", "matériau circulaire"}

type IMaterial interface {
	Update()
	Interact(graphic.Vector2) graphic.Vector2
	GetRect() graphic.Rect
	DrawLines()
}

type Material struct {
	rect graphic.Rect
	Type MaterialType
}

func NewMaterial(rect graphic.Rect, Type MaterialType) IMaterial {
	m := new(Material)

	switch Type {
	case PUSH_MATERIAL:
		return NewPushMaterial(m, rect)
	case PUSH_MATERIAL_DIRECTION:
		return NewDirectionnalPushMaterial(m, rect)
	case ROTATE_MATERIAL:
		return NewRotateMaterial(m, rect)
	default:
		panic("incorrect material type")

	}

}

func (m *Material) Update() {
	m.rect.Fill(rl.Black, 0)
}

func (m *Material) DrawLines() {
	m.rect.DrawLines(rl.Red, 0, 0.33)
}

func (m *Material) GetRect() graphic.Rect {
	return m.rect
}

func (m *Material) Interact(position graphic.Vector2) graphic.Vector2 {
	return graphic.NewVector2(1, 0)
}
