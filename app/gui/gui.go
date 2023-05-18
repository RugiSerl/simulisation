package gui

import (
	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/gui/components"
	"github.com/RugiSerl/simulisation/app/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	// largeur du panneau des paramètres
	SETTINGS_WIDTH = 350

	SETTINGS_HEIGHT = 700
	// durée en secondes de l'animation lorsque l'utilisateur ouvre la fenêtre de dialogue des paramètres
	ANIMATION_DURATION = 0.15
	TEXT_SPACING       = 0
	TEXT_SIZE          = 20
)

var (
	font rl.Font
)

func InitFont() {
	font = rl.LoadFontEx("assets/VarelaRound-Regular.ttf", TEXT_SIZE, []rune("'\"éèabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ:0123456789.- ()"))
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)
}

type UserInterface struct {
	AnimationTime float32

	menuRect   graphic.Rect
	rectOffset float32

	settings     []*components.Setting
	saveSettings *components.ImageButton
}

func NewInterface() *UserInterface {

	InitFont()
	u := new(UserInterface)
	u.rectOffset = 0

	u.InitSettingsPanel()

	return u

}

// initialise les paramètres à afficher et leur propriétés
func (u *UserInterface) InitSettingsPanel() {

	position := graphic.NewVector2(0, 15)
	parameteres := components.NewSetting("Paramètres", components.TYPE_NO_COMPONENT, font, TEXT_SIZE, position, graphic.ANCHOR_HORIZONTAL_MiDDLE, graphic.ANCHOR_TOP)
	position = position.Add(graphic.NewVector2(10, 45))

	gamerule := components.NewSetting("Gamerules", components.TYPE_NO_COMPONENT, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	position = position.Add(graphic.NewVector2(0, 32))

	UpdateAge := components.NewSetting("Update age", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	UpdateAge.SetBool(&settings.GameSettings.Gamerule.UpdateAge)
	position = position.Add(graphic.NewVector2(0, 30))

	Uncollide := components.NewSetting("Uncollide", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	Uncollide.SetBool(&settings.GameSettings.Gamerule.Uncollide)
	position = position.Add(graphic.NewVector2(0, 30))

	Reproduce := components.NewSetting("Reproduce", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	Reproduce.SetBool(&settings.GameSettings.Gamerule.Reproduce)
	position = position.Add(graphic.NewVector2(0, 30))

	Move := components.NewSetting("Move", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	Move.SetBool(&settings.GameSettings.Gamerule.Move)
	position = position.Add(graphic.NewVector2(0, 30))

	Kill := components.NewSetting("Kill", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	Kill.SetBool(&settings.GameSettings.Gamerule.Kill)
	position = position.Add(graphic.NewVector2(0, 45))

	visualSettings := components.NewSetting("Paramètres graphiques", components.TYPE_NO_COMPONENT, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	position = position.Add(graphic.NewVector2(0, 32))

	GradientEntities := components.NewSetting("Version couleur", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	GradientEntities.SetBool(&settings.GameSettings.VisualSettings.GradientEntities)
	position = position.Add(graphic.NewVector2(0, 30))

	DisplayStats := components.NewSetting("Afficher les statistiques", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	DisplayStats.SetBool(&settings.GameSettings.VisualSettings.DisplayStats)
	position = position.Add(graphic.NewVector2(0, 45))

	entitySettings := components.NewSetting("Paramètres de l'entité", components.TYPE_NO_COMPONENT, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	position = position.Add(graphic.NewVector2(0, 32))

	linearMove := components.NewSetting("Déplacement linéaire", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	linearMove.SetBool(&settings.GameSettings.EntitySettings.LinearMove)
	position = position.Add(graphic.NewVector2(0, 30))

	GoToClosestNeightbour := components.NewSetting("suit le plus proche \"voisin moral\"", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	GoToClosestNeightbour.SetBool(&settings.GameSettings.EntitySettings.GoToClosestNeightbour)
	position = position.Add(graphic.NewVector2(0, 30))

	radiusSensivity := components.NewSetting("portée des entités", components.TYPE_SLIDER, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	radiusSensivity.SetSliderValue(&settings.GameSettings.EntitySettings.RadiusSensivity, 0, 20)
	position = position.Add(graphic.NewVector2(0, 30))

	ChildMaximumDifference := components.NewSetting("différence morale avec l'enfant", components.TYPE_SLIDER, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	ChildMaximumDifference.SetSliderValue(&settings.GameSettings.EntitySettings.ChildMaximumDifference, 0, 300)
	position = position.Add(graphic.NewVector2(0, 30))

	MaximumAge := components.NewSetting("age maximal (0-20s)", components.TYPE_SLIDER, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	MaximumAge.SetSliderValue(&settings.GameSettings.EntitySettings.MaximumAge, 0, 20)
	position = position.Add(graphic.NewVector2(0, 30))

	BaseProbabilityReproduction := components.NewSetting("probabilité de reproduction", components.TYPE_SLIDER, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	BaseProbabilityReproduction.SetSliderValue(&settings.GameSettings.EntitySettings.BaseProbabilityReproduction, 0, 3e-3)
	position = position.Add(graphic.NewVector2(0, 30))

	u.settings = []*components.Setting{parameteres, gamerule, gamerule, UpdateAge, Uncollide, Reproduce, Move, Kill, visualSettings, GradientEntities, DisplayStats, entitySettings, linearMove, GoToClosestNeightbour, radiusSensivity, ChildMaximumDifference, MaximumAge, BaseProbabilityReproduction}

	u.saveSettings = components.NewImageButton(position.Add(graphic.NewVector2(0, 30)), rl.LoadTexture("assets/save.png"), graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)

}

// fonction principale de mise à jour
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

	for _, setting := range u.settings {
		setting.Update(u.menuRect)
	}
	u.saveSettings.Update(u.menuRect)

	if u.saveSettings.PressedState {
		settings.SaveSettings()
	}
	if rl.GetMousePosition().X > float32(rl.GetScreenWidth())-SETTINGS_WIDTH {
		u.rectOffset += rl.GetMouseWheelMove() * 20

		if u.rectOffset > 0 {
			u.rectOffset = 0
		}
	}

}

// affiche le rectangle blanc qui sert de base pour afficher les paramètres
func (u *UserInterface) DrawRectangle() {
	size := graphic.NewVector2(SETTINGS_WIDTH, SETTINGS_HEIGHT)

	position := graphic.GetRectCoordinatesWithAnchor(graphic.NewVector2(0, u.rectOffset), graphic.ANCHOR_RIGHT, graphic.ANCHOR_TOP, size, graphic.GetWindowRect())

	//déplace pour l'animation
	if u.AnimationTime < ANIMATION_DURATION {
		position = position.Add(graphic.NewVector2((size.X)*(ANIMATION_DURATION-u.AnimationTime)/ANIMATION_DURATION, 0))

	}

	u.menuRect = graphic.NewRectFromVector(position, size)

	u.menuRect.Fill(rl.White, 0.1)

	u.AnimationTime += rl.GetFrameTime()

}
