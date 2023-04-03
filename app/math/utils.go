package math

import (
	"log"
	"math/rand"
)

func RandomRange(min int, max int) int {
	if min > max {
		log.Fatal("le minimum est plus grand que le maximum")
	}
	return rand.Intn(max-min) + min
}
