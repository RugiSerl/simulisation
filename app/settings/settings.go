package settings

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const SETTINGS_FILENAME = "settings.json"

type Settings struct {
	VisualSettings VisualSettings
	Gamerule       Gamerule
	EntitySettings EntitySettings
}

type VisualSettings struct {
	//mode dans lequel les entités affiche un rond de couleur qui suit leur valeur morale selon le cercle chromatic
	GradientEntities bool

	//définit s'il faut afficher la zone de sensibilité dans laquelle les entités peuvent voir
	DisplaySensibilityZone bool

	//permet d'afficher les statistique du jeu, ex : le nombre d'entités, la fréquence d'images par seconde
	DisplayStats bool
}

type EntitySettings struct {
	// écart de différence morale maximum entre une entité et son enfant
	ChildMaximumDifference float32

	// rayon dans lequel une entité "voit" les autres entités
	RadiusSensivity float32

	// durée de vie d'une entité, en s
	MaximumAge float32

	//probabilité qu'une entité se reproduise s'il y a une seule autre entité
	BaseProbabilityReproduction float32

	//définit la façon dont se déplace l'entité.
	//Sur true, elle se déplace de façon linéaire et constante sur une unité de temps,
	//sur false, elle se déplace en divisant la distance avec la destination, sur un unité de frame
	LinearMove bool

	//définit la cible que doit suivre l'entité
	//sur true, l'entité suit l'entité la plus "proche" moralement
	//sur false elle suit la moyenne pondérée des entités environnantes en fonction de leur valeur morale
	GoToClosestNeightbour bool

	// définit la méthode selon laquelle les entités se repoussent afin de ne pas se stack.
	// Sur true, l'entité déplace toutes les entités sur son chemin.
	// Sur false, c'est l'entité qui se déplace lorsqu'elle est en collision avec une autre
	UncollideAgressive bool
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

	//permet à l'entité de tuer les entités environnantes en fonction de leur différence morale
	Kill bool
}

var (
	GameSettings Settings
)

func getDefaultSettings() Settings {

	return Settings{
		VisualSettings: VisualSettings{
			GradientEntities: true,
			DisplayStats:     false,
		},
		EntitySettings: EntitySettings{
			RadiusSensivity:             0.1 * 100,
			ChildMaximumDifference:      5,
			MaximumAge:                  5,
			BaseProbabilityReproduction: 1e-3,
			LinearMove:                  false,
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
	}

}

// permet de charger un fichier de paramètres, et prend les paramètres par défaut si ce dernier n'existe pas
func LoadSettings() {

	file, err := os.ReadFile(SETTINGS_FILENAME)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			GameSettings = getDefaultSettings()

		} else {
			log.Fatal(err)
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
