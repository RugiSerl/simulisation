package gameComponents

import (
	"github.com/RugiSerl/simulisation/app/math"
)

// Constante du nombre de voisins maximum
const NB_VOISINS_MAX = 6

// Définition de la classe "Entity"
type Entity struct {
	ValeurMorale uint8 // Valeur aléatoire qui va déterminer le groupe que l'entité rejoindra

	X uint8 // cordonnée x de l'entité sur la map
	Y uint8 // cordonnée y de l'entité sur la map

	renderX float32 // cordonnée x de l'entité sur la map pour affichage
	renderY float32 // cordonnée y de l'entité sur la map pour affichage

	Voisins      [NB_VOISINS_MAX]*Entity // Tableau des voisins avec lequels l'entité est en lien
	LiensVoisins [NB_VOISINS_MAX]uint8   // Tableau des forces de liaisons avec les voisins
	NbVoisins    uint8                   // nombre de voisins de l'entité
}

// Initialisation d'une instance entité
func NewEntity() *Entity {

	e := new(Entity)

	e.ValeurMorale = uint8(math.RandomRange(0, 255))
	e.X = uint8(math.RandomRange(0, 39))
	e.Y = uint8(math.RandomRange(0, 39))
	e.renderX = 0
	e.renderY = 0
	e.NbVoisins = 0

	return e
}

// Cette fonction permet de déplacer l'entité
func (e *Entity) Mouvement(newX uint8, newY uint8) {

	e.X = newX
	e.Y = newY

}

// la valeur morale est "cyclique", ce qui signifie que celle entre 5 et 254 est 6 par exemple
func (e *Entity) DistanceMorale(otherEntity *Entity) uint8 {
	distance := e.ValeurMorale - otherEntity.ValeurMorale
	if distance > 128 {
		return 255 - distance
	}

	return distance

}

func (e *Entity) NouveauLien(entiteVoisine *Entity) {
	e.
	e.LiensVoisins[e.NbVoisins] = 
	e.NbVoisins++
	
}
