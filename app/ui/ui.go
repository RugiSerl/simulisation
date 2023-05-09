package ui

import (
	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// largeur du panneau des paramètres
const SETTINGS_WIDTH = 200

// durée en secondes de l'animation lorsque l'utilisateur ouvre la fenêtre de dialogue des paramètres
const ANIMATION_DURATION = 0.1

type UserInterface struct {
	AnimationTime float32
}

func NewInterface() *UserInterface {

	u := new(UserInterface)

	return u

}

func (u *UserInterface) Update() {

	if global.SettingsOpen {
		u.UpdateSettings()

	}

	if rl.IsKeyPressed(rl.KeyS) {
		u.AnimationTime = 0
		global.SettingsOpen = !global.SettingsOpen

	}

}

// met à jour et affiche la fenêtre des paramètres
func (u *UserInterface) UpdateSettings() {

	size := graphic.NewVector2(SETTINGS_WIDTH, float32(rl.GetScreenHeight()))

	position := graphic.GetRectCoordinatesWithAnchor(graphic.NewVector2(0, 0), graphic.ANCHOR_RIGHT, graphic.ANCHOR_TOP, size, graphic.GetWindowRect())

	//déplace pour l'animation
	if u.AnimationTime < ANIMATION_DURATION {
		position = position.Add(graphic.NewVector2((size.X)*(ANIMATION_DURATION-u.AnimationTime)/ANIMATION_DURATION, 0))

	}

	rect := graphic.NewRectFromVector(position, size)

	rect.Fill(rl.White, 0.1)

	u.AnimationTime += rl.GetFrameTime()

}
