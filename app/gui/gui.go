package gui

import (
	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/gui/components"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// largeur du panneau des paramètres
const SETTINGS_WIDTH = 200

// durée en secondes de l'animation lorsque l'utilisateur ouvre la fenêtre de dialogue des paramètres
const ANIMATION_DURATION = 0.15

const TEXT_SIZE = 20
const TEXT_SPACING = 0

var (
	font rl.Font
)

func InitFont() {
	font = rl.LoadFontEx("assets/VarelaRound-Regular.ttf", TEXT_SIZE, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ:0123456789.- ()"))
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)
}

type UserInterface struct {
	AnimationTime float32

	menuRect    graphic.Rect
	pauseButton *components.ImageButton
}

func NewInterface() *UserInterface {

	u := new(UserInterface)
	u.pauseButton = components.NewImageButton(graphic.NewVector2(30, 30), rl.LoadTexture("assets/pause.png"), graphic.ANCHOR_LEFT, graphic.ANCHOR_BOTTOM)

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

	u.pauseButton.Update(u.menuRect)

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
