package gameComponents

import (
	"log"

	"github.com/RugiSerl/simulisation/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Constante du nombre de voisins maximum
const NB_VOISINS_MAX = 6

var (
	//texture utilisée pour afficher l'entité sur la fenêtre
	TextureEntite rl.Texture2D
)

// Définition de la classe "Entity"
type Entity struct {
	ValeurMorale uint8 // Valeur aléatoire qui va déterminer le groupe que l'entité rejoindra

	X float32 // cordonnée x de l'entité sur la map
	Y float32 // cordonnée y de l'entité sur la map

	Voisins      []*Entity // Tableau des voisins avec lequels l'entité est en lien
	LiensVoisins []uint8   // Tableau des forces de liaisons avec les voisins
	NbVoisins    uint8     // nombre de voisins de l'entité
}

// Initialisation d'une instance entité
func NewEntity() *Entity {

	e := new(Entity)
	e.ValeurMorale = uint8(math.RandomRange(0, 255))
	e.NbVoisins = 0

	return e
}

// Cette fonction permet de déplacer l'entité
func (e *Entity) Mouvement(newX float32, newY float32) {

	e.X = newX
	e.Y = newY

}

// Cette fonction
func (e *Entity) Update() {
	e.render()
}

// Cette fonction s'occupe d'afficher visuellement l'entité
func (e *Entity) render() {
	rl.DrawTextureEx(TextureEntite, rl.NewVector2(e.X, e.Y), 0, 1, rl.White)

}

// la valeur morale est "cyclique", ce qui signifie que celle entre 5 et 254 est 6 par exemple
func (e *Entity) DistanceMorale(otherEntity *Entity) uint8 {
	distance := e.ValeurMorale - otherEntity.ValeurMorale
	if distance > 128 {
		return 255 - distance
	}

	return distance

}

// création d'un lien avec une autre entité
func (e *Entity) NouveauLien(entiteVoisine *Entity) {
	if len(e.LiensVoisins) < NB_VOISINS_MAX {
		forceDuLien := 128 - e.DistanceMorale(entiteVoisine)

		e.Voisins = append(e.Voisins, entiteVoisine)
		e.LiensVoisins = append(e.LiensVoisins, forceDuLien)

		entiteVoisine.Voisins = append(entiteVoisine.Voisins, e)
		entiteVoisine.LiensVoisins = append(entiteVoisine.LiensVoisins, forceDuLien)

		e.X = entiteVoisine.X
		e.Y = entiteVoisine.Y + 20

	} else {
		log.Fatal("l'entité a déjà atteint le nombre maximal de voisins")
	}

}
