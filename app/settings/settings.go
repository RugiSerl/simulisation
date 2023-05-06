package settings

type Settings struct {
	VisualSettings VisualSettings
	Gamerule       Gamerule
}

type VisualSettings struct {
	GradientEntities bool
	DisplayStats bool
}

type EntitySettings struct {
	// écart de différence morale maximum entre une entité et son enfant
	ChildMaximumDifference      uint8

	// rayon dans lequel une entité "voit" les autres entités
	RadiusSensivity             float32

	// durée de vie d'une entité, en s
	MaximumAge                  float32

	//probabilité qu'une entité se reproduise s'il y a une seule autre entité
	BaseProbabilityReproduction float32
}

type Gamerule struct {
	UpdateAge bool
	Uncollide bool
	Reproduce bool
	move bool
}

var (
	GameSettings Settings
)

func GetDefaultSettings() Settings {

	return Settings{
		VisualSettings: VisualSettings{
			GradientEntities: true,
		},
	}

}

func LoadSettings() {
	GameSettings = GetDefaultSettings()
}
