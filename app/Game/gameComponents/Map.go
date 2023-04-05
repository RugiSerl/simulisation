package gameComponents

type Map struct {
}

// taille d'une cellule de la grille
const GRID_CELL_SIZE = 24
const GRID_SIZE = 100

func NewMap() *Map {

	m := new(Map)

	return m

}
