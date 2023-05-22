package settings

// Par Raphaël

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// ce fichier gère les paramètres du jeu, qui sont stockés sous forme de struct

const SETTINGS_FILENAME = "settings.json"

// paramètres principaux
type Settings struct {
	VisualSettings    VisualSettings
	Gamerule          Gamerule
	EntitySettings    EntitySettings
	UserInputSettings UserInputSettings
}

// section sur les paramètres sur l'entrée utilisateur
type UserInputSettings struct {
	// définit si les entités que l'utilisateur fait spawner on une valeur morale random
	SpawnRandomValeurMorale bool

	// valeur morale de la prochaine entité si SpawnRandomValeurMorale est sur false
	EntityValeurMoraleOnSpawn float32

	// rayon dans lequel on supprime les entités lorsque le clic droit de la souris est enfoncé
	DeleteRadius float32
}

// section sur les paramètres graphiques
type VisualSettings struct {
	//mode dans lequel les entités affiche un rond de couleur qui suit leur valeur morale selon le cercle chromatic
	GradientEntities bool

	//définit s'il faut afficher la zone de sensibilité dans laquelle les entités peuvent voir
	DisplaySensibilityZone bool

	//permet d'afficher les statistique du jeu, ex : le nombre d'entités, la fréquence d'images par seconde
	DisplayStats bool

	//définit la cible en images par secondes que doit atteindre le jeu, n'est prise en compte qu'au démarrage de ce dernier
	MaxFps float32

	//définit s'il faut appliquer une couleur sur le jeu pour enlever les résidus de l'image précédente
	ClearBackground bool
}

// section sur les paramètres de l'entité, et des valeurs arbitraires qu'elles utilisent
type EntitySettings struct {
	// écart de différence morale maximum entre une entité et son enfant
	ChildMaximumDifference float32

	// rayon dans lequel une entité "voit" les autres entités
	RadiusSensivity float32

	// durée de vie d'une entité, en s
	MaximumAge float32

	//probabilité qu'une entité se reproduise s'il y a une seule autre entité
	BaseProbabilityReproduction float32

	//probabilité qu'une entité tue une autre entité
	BaseProbabilityKill float32

	//définit la façon dont se déplace l'entité.
	//Sur true, elle se déplace de façon linéaire et constante sur une unité de temps,
	//sur false, elle se déplace en divisant la distance avec la destination, sur un unité de frame
	LinearMove bool

	// vitesse à laquelle se déplacent les entités, si elles se déplacent de façon linéaire
	Speed float32

	//définit la cible que doit suivre l'entité
	//sur true, l'entité suit l'entité la plus "proche" moralement
	//sur false elle suit la moyenne pondérée des entités environnantes en fonction de leur valeur morale
	GoToClosestNeightbour bool

	// définit la méthode selon laquelle les entités se repoussent afin de ne pas se stack.
	// Sur true, l'entité déplace toutes les entités sur son chemin.
	// Sur false, c'est l'entité qui se déplace lorsqu'elle est en collision avec une autre
	UncollideAgressive bool
}

// section sur les paramètres des règles de jeu, qui décide si elles sont appliquées ou non
type Gamerule struct {
	//controle si l'on met à jour l'âge de l'entité et par conséquent si elle meurt
	UpdateAge bool

	//éloigne les entités afin qu'elles ne se stackent pas
	Uncollide bool

	//permet àl'entité de se reproduire
	Reproduce bool

	//permet à l'entité de se déplacer
	Move bool

	//permet à l'entité de tuer les entités environnantes en fonction de leur différence morale
	Kill bool
}

var (
	GameSettings Settings
)

// paramètres par défaut
func getDefaultSettings() Settings {

	return Settings{
		VisualSettings: VisualSettings{
			GradientEntities: true,
			DisplayStats:     false,
			MaxFps:           120,
			ClearBackground:  true,
		},
		EntitySettings: EntitySettings{
			RadiusSensivity:             10,
			ChildMaximumDifference:      5,
			MaximumAge:                  5,
			BaseProbabilityReproduction: 1e-3,
			BaseProbabilityKill:         1e-5,
			LinearMove:                  false,
			Speed:                       20,
			UncollideAgressive:          false,
			GoToClosestNeightbour:       true,
		},
		Gamerule: Gamerule{
			UpdateAge: true,
			Uncollide: true,
			Reproduce: true,
			Move:      true,
			Kill:      true,
		},
		UserInputSettings: UserInputSettings{
			SpawnRandomValeurMorale:   true,
			EntityValeurMoraleOnSpawn: 0,
			DeleteRadius:              2,
		},
	}

}

// permet de charger un fichier de paramètres, et prend les paramètres par défaut si ce dernier n'existe pas
func LoadSettings() {

	file, err := os.ReadFile(SETTINGS_FILENAME)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			GameSettings = getDefaultSettings()

		} else {
			panic(err)
		}
	} else {
		err2 := json.Unmarshal(file, &GameSettings)
		if err2 != nil {
			fmt.Println("parsing settings failed, :", err2)
			fmt.Println("Loading with default settings ...")
			GameSettings = getDefaultSettings()
		}
	}

}

// sauvegarde les paramètres en les écrivant dans un fichier settings
func SaveSettings() {
	json, err := json.MarshalIndent(GameSettings, "", "	")

	if err != nil {
		fmt.Println(err)
	}

	_ = ioutil.WriteFile(SETTINGS_FILENAME, json, 0644)

}
