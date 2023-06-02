package material

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Material struct {
	Rect graphic.Rect
}

func NewMaterial(rect graphic.Rect) *Material {
	m := new(Material)
	m.Rect = rect

	return m

}

func (m *Material) Update() {
	m.Rect.Fill(rl.Black, 0)
}
