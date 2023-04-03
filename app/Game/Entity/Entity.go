package Entity
import "math/rand"

type Entity struct {
	Age uint8
}

func NewEntity() *Entity {
	
	e := new(Entity)
	e.Age = 0 
}

func (e *Entity) Update() {
	e.Age += 5
	if e.Age
}