package Game

import (
	"github.com/RugiSerl/simulisation/app/assets"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/gui/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PANEL_HEIGHT = 70.0
	PANEL_WIDTH  = 200.0
)

type SaveLoadPanel struct {
	containingRect graphic.Rect
	SaveButton     *components.ImageButton
	LoadButton     *components.ImageButton
}

func NewSaveLoadPanel() *SaveLoadPanel {
	s := new(SaveLoadPanel)

	s.SaveButton = components.NewImageButton(graphic.NewVector2(0, 0), rl.LoadTexture(assets.AssetPath("save2.png")), graphic.ANCHOR_LEFT, graphic.ANCHOR_BOTTOM)
	s.LoadButton = components.NewImageButton(graphic.NewVector2(0, 0), rl.LoadTexture(assets.AssetPath("load.png")), graphic.ANCHOR_RIGHT, graphic.ANCHOR_BOTTOM)

	return s
}

func (s *SaveLoadPanel) Update() {
	s.updateRect()

	s.containingRect.Fill(rl.White, 0.3)
	s.SaveButton.Update(s.containingRect)
	s.LoadButton.Update(s.containingRect)

}

func (s *SaveLoadPanel) updateRect() {
	size := graphic.NewVector2(PANEL_WIDTH, PANEL_HEIGHT)
	position := graphic.GetRectCoordinatesWithAnchor(graphic.NewVector2(0, 0), graphic.ANCHOR_HORIZONTAL_MiDDLE, graphic.ANCHOR_BOTTOM, size, graphic.GetWindowRect())

	s.containingRect = graphic.NewRectFromVector(position, size)

}
