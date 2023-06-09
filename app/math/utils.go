package math

// Par Raphaël

import (
	"fmt"
	"math"
	"math/rand"
)

//ce fichier contient quelques fonctions mathématiques simples pratiques pour le programme

// renvoie un nombre random de type int selon une intervalle définie
func RandomRange(min int, max int) int {
	if min >= max {
		panic("le minimum est plus grand que le maximum")
	}
	return rand.Intn(max-min) + min
}

// fait un test et retourne true selon la probabilité passée par probability, qui doit se situer entre 0 et 1.
func RandomProbability(probability float64) bool {
	if probability < 0 || probability > 1 {
		fmt.Println("error probability ", probability)
		panic("error out of range :")
	}
	return rand.Float32() < float32(probability)
}

func ArcCos(x float64) float64 {
	return math.Acos(x)
}
