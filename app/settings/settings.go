package settings

type Settings struct {
	VisualSettings VisualSettings
}

type VisualSettings struct {
	GradientEntities bool
	

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
