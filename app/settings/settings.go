package settings

type Settings struct {
	GradientEntities bool
}

var (
	GameSettings Settings
)

func GetDefaultSettings() Settings {

	return Settings{GradientEntities: false}

}

func LoadSettings() {
	GameSettings = GetDefaultSettings()
}
