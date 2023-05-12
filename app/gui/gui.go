package gui

import (
	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/gui/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	// largeur du panneau des paramètres
	SETTINGS_WIDTH = 300
	// durée en secondes de l'animation lorsque l'utilisateur ouvre la fenêtre de dialogue des paramètres
	ANIMATION_DURATION = 0.15
	TEXT_SPACING       = 0
	TEXT_SIZE          = 24
)

var (
	font rl.Font
)

func InitFont() {
	font = rl.LoadFontEx("assets/VarelaRound-Regular.ttf", TEXT_SIZE, []rune("èabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ:0123456789.- ()"))
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)
}

type UserInterface struct {
	AnimationTime float32

	menuRect     graphic.Rect
	testLabel    *components.Label
	testCheckBox *components.CheckBox
}

func NewInterface() *UserInterface {

	InitFont()
	u := new(UserInterface)
	u.testLabel = components.Newlabel("Paramètres", font, TEXT_SIZE, graphic.NewVector2(0, 15), graphic.ANCHOR_HORIZONTAL_MiDDLE, graphic.ANCHOR_TOP)
	u.testCheckBox = components.NewCheckBox(graphic.NewVector2(100, 100), false, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)

	return u

}

func (u *UserInterface) Update() {

	if global.SettingsOpen {
		u.UpdateSettings()

	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		u.AnimationTime = 0
		global.SettingsOpen = !global.SettingsOpen

	}

}

// met à jour et affiche la fenêtre des paramètres
func (u *UserInterface) UpdateSettings() {

	u.DrawRectangle()

	u.testLabel.Render(u.menuRect)
	u.testCheckBox.Update(u.menuRect)

}

func (u *UserInterface) DrawRectangle() {
	size := graphic.NewVector2(SETTINGS_WIDTH, float32(rl.GetScreenHeight()))

	position := graphic.GetRectCoordinatesWithAnchor(graphic.NewVector2(0, 0), graphic.ANCHOR_RIGHT, graphic.ANCHOR_TOP, size, graphic.GetWindowRect())

	//déplace pour l'animation
	if u.AnimationTime < ANIMATION_DURATION {
		position = position.Add(graphic.NewVector2((size.X)*(ANIMATION_DURATION-u.AnimationTime)/ANIMATION_DURATION, 0))

	}

	u.menuRect = graphic.NewRectFromVector(position, size)

	u.menuRect.Fill(rl.White, 0.1)

	u.AnimationTime += rl.GetFrameTime()

}
