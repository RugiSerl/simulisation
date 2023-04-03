package Entity

import (
	"github.com/RugiSerl/simulisation/app/utils"
)

type Entity struct {
	Age  uint8
	Dead bool
	X    float32
	Y    float32
}

func NewEntity() *Entity {

	e := new(Entity)
	e.Age = 0
	e.Dead = false

	return e
}

func (e *Entity) Update() {
	e.Age += 5
	if e.Age == uint8(utils.RandomRange(50, 90)) {
		e.Dead = true
	}
}

func (e *Entity) Mouvement() {

}
