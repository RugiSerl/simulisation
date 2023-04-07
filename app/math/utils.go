package math

import (
	"log"
	"math/rand"
	"math"
)

// renvoie un nombre random de type int selon une intervalle définie
func RandomRange(min int, max int) int {
	if min > max {
		log.Fatal("le minimum est plus grand que le maximum")
	}
	return rand.Intn(max-min) + min
}

func Distance(value1 int, value2 int) int{
	return math.Abs(value1-value2)
}
