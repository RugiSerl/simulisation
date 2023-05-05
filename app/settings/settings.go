package settings

type Settings struct {
	VisualSettings VisualSettings
	Gamerule       Gamerule
}

type VisualSettings struct {
	GradientEntities bool
}

type Gamerule struct {
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
