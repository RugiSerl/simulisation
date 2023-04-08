package gameComponents

import rl "github.com/gen2brain/raylib-go/raylib"

type Map struct {
	textures map[string]rl.Texture2D
	// Tableau contenant les informations sur les cellules de la map (champ, maison, ...)
	cells [GRID_SIZE][GRID_SIZE]string
}

// Taille d'une cellule de la grille
const GRID_CELL_SIZE = 24

// Taille de la grille
const GRID_SIZE = 40

// Création d'une instance de la classe Map
func NewMap() *Map {

	m := new(Map)

	m.textures = make(map[string]rl.Texture2D)
	m.textures["field"] = rl.LoadTexture("assets/field.png")
	m.textures["lake"] = rl.LoadTexture("assets/lake.png")
	m.textures["house"] = rl.LoadTexture("assets/house.png")
	m.textures["person"] = rl.LoadTexture("assets/person.png")

	return m

}
