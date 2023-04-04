package gameComponents

import (
	"github.com/RugiSerl/simulisation/app/math"
)

type Entity struct {
	Age  uint8
	Dead bool

	X float32 // cordonnée x du joueur sur la map
	Y float32 // cordonnée y du joueur sur la map
}

func NewEntity() *Entity {

	e := new(Entity)
	e.Age = 0
	e.Dead = false

	return e
}

// Cette fonction est appelée à chaque tour, et mets à jours les statistiques de l'entité
func (e *Entity) Update() {
	e.Age += 5
	if e.Age == uint8(math.RandomRange(50, 90)) {
		e.Dead = true
	}
}

func (e *Entity) Mouvement() {

}
