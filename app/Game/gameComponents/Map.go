package gameComponents

type Map struct {
	fieldTexture rl.Texture2d
	


}

// taille d'une cellule de la grille
const GRID_CELL_SIZE = 24
const GRID_SIZE = 40

func NewMap() *Map {

	m := new(Map)
	m.fieldTexture = rl.LoadTexture("assets/champ.png")

	return m

}
