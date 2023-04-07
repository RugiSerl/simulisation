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

	Voisins [NB_VOISINS_MAX]*Entity // Tableau des voisins avec lequels l'entité est en lien
	LiensVoisins [NB_VOISINS_MAX]uint8 // Tableau des forces de liaisons avec les voisins
}

// Initialisation des valeurs de l'entité
func NewEntity() *Entity {

	e := new(Entity)

	e.ValeurMorale = uint8(math.RandomRange(0, 255))
	e.X = uint8(math.RandomRange(0, 39))
	e.Y = uint8(math.RandomRange(0, 39))
	e.renderX = 0
	e.renderY = 0
	e.Voisins = []

	return e
}

// Cette fonction permet de déplacer l'entité
func (e *Entity) Movement(newX uint8,newY uint8) {

	e.X = newX
	e.Y = newY

}

//
func (e *Entity) NewVoisin() {

}


func (e *Entity) EvaluateMoralDiff(otherEntity *Entity) int{
	return math.AbsoluteValue()


}