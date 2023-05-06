package settings

type Settings struct {
	VisualSettings VisualSettings
	Gamerule       Gamerule
	EntitySettings EntitySettings
}

type VisualSettings struct {
	//mode dans lequel les entités affiche un rond de couleur qui suit leur valeur morale selon le cercle chromatic
	GradientEntities bool

	//permet d'afficher les statistique du jeu, ex : le nombre d'entités, la fréquence d'images par seconde
	DisplayStats bool
}

type EntitySettings struct {
	// écart de différence morale maximum entre une entité et son enfant
	ChildMaximumDifference uint8

	// rayon dans lequel une entité "voit" les autres entités
	RadiusSensivity float32

	// durée de vie d'une entité, en s
	MaximumAge float32

	//probabilité qu'une entité se reproduise s'il y a une seule autre entité
	BaseProbabilityReproduction float32
}

type Gamerule struct {
	//controle si l'on met à jour l'âge de l'entité et par conséquent si elle meurt
	UpdateAge bool

	//éloigne les entités afin qu'elles ne se stackent pas
	Uncollide bool

	//permet àl'entité de se reproduire
	Reproduce bool

	//permet à l'entité de se déplacer
	Move bool
}

var (
	GameSettings Settings
)

func GetDefaultSettings() Settings {

	return Settings{
		VisualSettings: VisualSettings{
			GradientEntities: true,
			DisplayStats:     true,
		},
		EntitySettings: EntitySettings{
			RadiusSensivity:             0.1 * 100,
			ChildMaximumDifference:      5,
			MaximumAge:                  10,
			BaseProbabilityReproduction: 1e-3,
		},
		Gamerule: Gamerule{
			UpdateAge: true,
			Uncollide: true,
			Reproduce: true,
			Move:      true,
		},
	}

}

func LoadSettings() {
	GameSettings = GetDefaultSettings()
}
