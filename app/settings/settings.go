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
	ChildMaximumDifference      uint8
	RadiusSensivity             float32
	MaximumAge                  float32
	BaseProbabilityReproduction float32
}

type Gamerule struct {
}

var (
	GameSettings Settings
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
