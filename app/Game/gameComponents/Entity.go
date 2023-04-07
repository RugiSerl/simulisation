package gameComponents

import (
	"github.com/RugiSerl/simulisation/app/math"
)

// Définition de la classe "Entity"
type Entity struct {

	Age  uint8
	Dead bool

	X uint8 // cordonnée x de l'entité sur la map
	Y uint8 // cordonnée y de l'entité sur la map

	renderX float32 // cordonnée x de l'entité sur la map pour affichage
	renderY float32 // cordonnée y de l'entité sur la map pour affichage
}

// Initialisation des valeurs de l'entité
func NewEntity() *Entity {

	e := new(Entity)

	e.Age = 0
	e.Dead = false
	e.X = uint8(math.RandomRange(0, 39))
	e.Y = uint8(math.RandomRange(0, 39))
	e.renderX = 0
	e.renderY = 0

	return e
}

// Les prochaines fonctions sont appelées à chaque tour de manière à mettre à jour les statistiques de l'entité selon les actions dans la simulation

// Cette fonction détermine si l'entité continue à vivre ou non 
func (e *Entity) Update() {

	e.Age += 5 // augmentation de l'age de l'entité

	if e.Age == uint8(math.RandomRange(50, 90)) {
		e.Dead = true
	}
}

// Cette fonction permet de déplacer l'entité
func (e *Entity) Movement(newX uint8,newY uint8) {

	e.X = newX
	e.Y = newY

}

