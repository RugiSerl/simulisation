package gui

// Par Raphaël

import (
	"github.com/RugiSerl/simulisation/app/assets"
	"github.com/RugiSerl/simulisation/app/global"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/gui/components"
	"github.com/RugiSerl/simulisation/app/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	// largeur et hauteur du panneau des paramètres
	SETTINGS_MENU_WIDTH  = 350
	SETTINGS_MENU_HEIGHT = 1100

	// durée en secondes de l'animation lorsque l'utilisateur ouvre la fenêtre de dialogue des paramètres
	ANIMATION_DURATION = 0.15
	TEXT_SPACING       = 0
	TEXT_SIZE          = 20
)

var (
	font rl.Font

	AnimationTime float32
)

func InitFont() {
	font = rl.LoadFontEx(assets.AssetPath("VarelaRound-Regular.ttf"), TEXT_SIZE, []rune("'\"ÉâéèàabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ:0123456789.- ()"))
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)
}

type UserInterface struct {
	settingsMenuRect graphic.Rect

	rectOffset float32

	settings      []*components.Setting
	saveSettings  *components.ImageButton
	openSettings  *components.ImageButton
	closeSettings *components.ImageButton
}

func NewInterface() *UserInterface {

	InitFont()
	u := new(UserInterface)
	u.rectOffset = 0
	u.openSettings = components.NewImageButton(graphic.NewVector2(20, 20), rl.LoadTexture(assets.AssetPath("menu.png")), graphic.ANCHOR_RIGHT, graphic.ANCHOR_TOP)

	u.InitSettingsPanel()

	return u

}

// initialise les paramètres à afficher et leur propriétés
func (u *UserInterface) InitSettingsPanel() {

	u.closeSettings = components.NewImageButton(graphic.NewVector2(20, 20), rl.LoadTexture(assets.AssetPath("close.png")), graphic.ANCHOR_RIGHT, graphic.ANCHOR_TOP)

	position := graphic.NewVector2(0, 15)
	parameteres := components.NewSetting("Paramètres", components.TYPE_NO_COMPONENT, font, TEXT_SIZE, position, graphic.ANCHOR_HORIZONTAL_MiDDLE, graphic.ANCHOR_TOP)
	position = position.Add(graphic.NewVector2(10, 45))

	gamerule := components.NewSetting("Règles du jeu", components.TYPE_NO_COMPONENT, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	position = position.Add(graphic.NewVector2(0, 32))

	UpdateAge := components.NewSetting("Mise à jour de l'âge", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	UpdateAge.SetBool(&settings.GameSettings.Gamerule.UpdateAge)
	position = position.Add(graphic.NewVector2(0, 30))

	ReactMaterial := components.NewSetting("Réaction aux matériaux", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	ReactMaterial.SetBool(&settings.GameSettings.Gamerule.ReactMaterial)
	position = position.Add(graphic.NewVector2(0, 30))

	Uncollide := components.NewSetting("Les entités se repoussent", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	Uncollide.SetBool(&settings.GameSettings.Gamerule.Uncollide)
	position = position.Add(graphic.NewVector2(0, 30))

	Reproduce := components.NewSetting("Reproduction", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	Reproduce.SetBool(&settings.GameSettings.Gamerule.Reproduce)
	position = position.Add(graphic.NewVector2(0, 30))

	Move := components.NewSetting("Déplacement", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	Move.SetBool(&settings.GameSettings.Gamerule.Move)
	position = position.Add(graphic.NewVector2(0, 30))

	Kill := components.NewSetting("Les entités s'entretuent", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	Kill.SetBool(&settings.GameSettings.Gamerule.Kill)
	position = position.Add(graphic.NewVector2(0, 45))

	visualSettings := components.NewSetting("Paramètres graphiques", components.TYPE_NO_COMPONENT, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	position = position.Add(graphic.NewVector2(0, 32))

	GradientEntities := components.NewSetting("Version couleur", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	GradientEntities.SetBool(&settings.GameSettings.VisualSettings.GradientEntities)
	position = position.Add(graphic.NewVector2(0, 30))

	DisplaySensibilityZone := components.NewSetting("Afficher les zones de vision", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	DisplaySensibilityZone.SetBool(&settings.GameSettings.VisualSettings.DisplaySensibilityZone)
	position = position.Add(graphic.NewVector2(0, 30))

	DisplayStats := components.NewSetting("Afficher les statistiques", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	DisplayStats.SetBool(&settings.GameSettings.VisualSettings.DisplayStats)
	position = position.Add(graphic.NewVector2(0, 30))

	Mode3d := components.NewSetting("Mode 3d", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	Mode3d.SetBool(&settings.GameSettings.VisualSettings.Mode3d)
	position = position.Add(graphic.NewVector2(0, 30))

	MaxFps := components.NewSetting("Fps cible (1-240)", components.TYPE_SLIDER, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	MaxFps.SetSliderValue(&settings.GameSettings.VisualSettings.MaxFps, 1, 240)
	position = position.Add(graphic.NewVector2(0, 30))

	ClearBackground := components.NewSetting("nettoyer les résidus", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	ClearBackground.SetBool(&settings.GameSettings.VisualSettings.ClearBackground)
	position = position.Add(graphic.NewVector2(0, 30))

	Background := components.NewSetting("fond d'écran", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	Background.SetBool(&settings.GameSettings.VisualSettings.Background)
	position = position.Add(graphic.NewVector2(0, 45))

	entitySettings := components.NewSetting("Paramètres de l'entité", components.TYPE_NO_COMPONENT, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	position = position.Add(graphic.NewVector2(0, 32))

	linearMove := components.NewSetting("Déplacement linéaire", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	linearMove.SetBool(&settings.GameSettings.EntitySettings.LinearMove)
	position = position.Add(graphic.NewVector2(0, 30))

	UnCollideAgressive := components.NewSetting("Éloignement agressif", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	UnCollideAgressive.SetBool(&settings.GameSettings.EntitySettings.UncollideAgressive)
	position = position.Add(graphic.NewVector2(0, 30))

	GoToClosestNeightbour := components.NewSetting("suit le plus proche \"voisin moral\"", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	GoToClosestNeightbour.SetBool(&settings.GameSettings.EntitySettings.GoToClosestNeightbour)
	position = position.Add(graphic.NewVector2(0, 30))

	radiusSensivity := components.NewSetting("portée des entités", components.TYPE_SLIDER, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	radiusSensivity.SetSliderValue(&settings.GameSettings.EntitySettings.RadiusSensivity, 0, 20)
	position = position.Add(graphic.NewVector2(0, 30))

	Speed := components.NewSetting("vitesse des entités", components.TYPE_SLIDER, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	Speed.SetSliderValue(&settings.GameSettings.EntitySettings.Speed, 1, 100)
	position = position.Add(graphic.NewVector2(0, 30))

	ChildMaximumDifference := components.NewSetting("différence morale avec l'enfant", components.TYPE_SLIDER, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	ChildMaximumDifference.SetSliderValue(&settings.GameSettings.EntitySettings.ChildMaximumDifference, 1, 300)
	position = position.Add(graphic.NewVector2(0, 30))

	MaximumAge := components.NewSetting("age maximal (0-20s)", components.TYPE_SLIDER, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	MaximumAge.SetSliderValue(&settings.GameSettings.EntitySettings.MaximumAge, 0, 20)
	position = position.Add(graphic.NewVector2(0, 30))

	BaseProbabilityReproduction := components.NewSetting("probabilité de reproduction", components.TYPE_SLIDER, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	BaseProbabilityReproduction.SetSliderValue(&settings.GameSettings.EntitySettings.BaseProbabilityReproduction, 1e-5, 3e-3)
	position = position.Add(graphic.NewVector2(0, 30))

	BaseProbabilityKill := components.NewSetting("probabilité de meurtre", components.TYPE_SLIDER, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	BaseProbabilityKill.SetSliderValue(&settings.GameSettings.EntitySettings.BaseProbabilityKill, 1e-7, 5e-5)
	position = position.Add(graphic.NewVector2(0, 45))

	materialSettings := components.NewSetting("Paramètres du matériau", components.TYPE_NO_COMPONENT, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	position = position.Add(graphic.NewVector2(0, 32))

	InteractionSpeed := components.NewSetting("force des interactions", components.TYPE_SLIDER, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	InteractionSpeed.SetSliderValue(&settings.GameSettings.InteractionSpeed, 0.01, 10)
	position = position.Add(graphic.NewVector2(0, 45))

	UserInputSettings := components.NewSetting("Paramètres d'entrée utilisateur", components.TYPE_NO_COMPONENT, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	position = position.Add(graphic.NewVector2(0, 32))

	SpawnRandomValeurMorale := components.NewSetting("valeurs morales aléatoire", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	SpawnRandomValeurMorale.SetBool(&settings.GameSettings.UserInputSettings.SpawnRandomValeurMorale)
	position = position.Add(graphic.NewVector2(0, 30))

	EntityValeurMoraleOnSpawn := components.NewSetting("valeur morale des entités", components.TYPE_SLIDER, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	EntityValeurMoraleOnSpawn.SetSliderValue(&settings.GameSettings.UserInputSettings.EntityValeurMoraleOnSpawn, 0, 255)
	position = position.Add(graphic.NewVector2(0, 30))

	DeleteRadius := components.NewSetting("Rayon de suppression", components.TYPE_SLIDER, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	DeleteRadius.SetSliderValue(&settings.GameSettings.UserInputSettings.DeleteRadius, 0.5, 40)
	position = position.Add(graphic.NewVector2(0, 30))

	SpawnMaterial := components.NewSetting("Insertion de matériau (alt)", components.TYPE_BOOL, font, TEXT_SIZE, position, graphic.ANCHOR_LEFT, graphic.ANCHOR_TOP)
	SpawnMaterial.SetBool(&settings.GameSettings.UserInputSettings.SpawnMaterial)
	position = position.Add(graphic.NewVector2(0, 45))

	u.settings = []*components.Setting{parameteres, gamerule, gamerule, UpdateAge, Uncollide, ReactMaterial, Reproduce, Move, Kill, visualSettings, GradientEntities, Mode3d, DisplaySensibilityZone, MaxFps, DisplayStats, ClearBackground, Background, entitySettings, linearMove, GoToClosestNeightbour, UnCollideAgressive, radiusSensivity, Speed, ChildMaximumDifference, MaximumAge, BaseProbabilityReproduction, BaseProbabilityKill, materialSettings, InteractionSpeed, UserInputSettings, SpawnRandomValeurMorale, EntityValeurMoraleOnSpawn, DeleteRadius, SpawnMaterial}

	u.saveSettings = components.NewImageButton(position, rl.LoadTexture(assets.AssetPath("save.png")), graphic.ANCHOR_HORIZONTAL_MiDDLE, graphic.ANCHOR_TOP)

}

// fonction principale de mise à jour
func (u *UserInterface) Update() {

	if global.SettingsOpen {
		u.UpdateSettings()

	} else {
		u.openSettings.Update(graphic.GetWindowRect())
		if u.openSettings.PressedState {
			global.SettingsOpen = true
		}
	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		AnimationTime = 0
		global.SettingsOpen = !global.SettingsOpen

		// disable/enable cursor when the game is in 3d while not in the settings
		// unhandled exceptions (I'm lazy) when the settings for the first time or if the user closes the settings with the cross button
		if settings.GameSettings.Mode3d {
			if global.SettingsOpen {
				rl.EnableCursor()
			} else {
				rl.DisableCursor()

			}
		}

	}

}

// met à jour et affiche la fenêtre des paramètres
func (u *UserInterface) UpdateSettings() {

	//mettre à jour et tracer le rectangle qui sert de base aux paramètres
	u.DrawRectangle()

	for _, setting := range u.settings {
		setting.Update(u.settingsMenuRect)
	}

	u.saveSettings.Update(u.settingsMenuRect)
	if u.saveSettings.PressedState {
		settings.SaveSettings()
	}

	u.closeSettings.Update(u.settingsMenuRect)
	if u.closeSettings.PressedState {
		global.SettingsOpen = false
		AnimationTime = 0
	}

	//gérer le scroll pour faire descendre le panneau des paramètres
	if rl.GetMousePosition().X > float32(rl.GetScreenWidth())-SETTINGS_MENU_WIDTH {
		u.rectOffset += rl.GetMouseWheelMove() * 20

		if u.rectOffset > 0 {
			u.rectOffset = 0
		}
	}

}

// affiche le rectangle blanc qui sert de base pour afficher les paramètres
func (u *UserInterface) DrawRectangle() {
	size := graphic.NewVector2(SETTINGS_MENU_WIDTH, SETTINGS_MENU_HEIGHT)

	position := graphic.GetRectCoordinatesWithAnchor(graphic.NewVector2(0, u.rectOffset), graphic.ANCHOR_RIGHT, graphic.ANCHOR_TOP, size, graphic.GetWindowRect())

	//déplace ce rectangle pour l'animation
	if AnimationTime < ANIMATION_DURATION {
		offset := (ANIMATION_DURATION - AnimationTime) / ANIMATION_DURATION
		position = position.Add(graphic.NewVector2(offset*offset*size.X, 0))

	}

	u.settingsMenuRect = graphic.NewRectFromVector(position, size)
	u.settingsMenuRect.Fill(rl.White, 0.1)

	AnimationTime += rl.GetFrameTime()

}
