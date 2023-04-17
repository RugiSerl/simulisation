package math

import (
	"log"
	"math"
	"math/rand"
)

// renvoie un nombre random de type int selon une intervalle définie
func RandomRange(min int, max int) int {
	if min > max {
		log.Fatal("le minimum est plus grand que le maximum")
	}
	return rand.Intn(max-min) + min
}

func RandomProbability(probability float64) bool {
	return rand.Float32() < float32(probability)
}

func Exp(x float64) float64 {
	return math.Exp(x)
}
