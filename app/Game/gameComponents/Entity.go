package gameComponents

import (
	"log"

	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Constante du nombre de voisins maximum
const NB_VOISINS_MAX = 6

// échelle qui correspond à la taille des entité (1 => 128; 0.5 => 64; ...)
const SCALE = 0.5

var (
	//texture utilisée pour afficher l'entité sur la fenêtre
	TextureEntite rl.Texture2D
)

// Définition de la classe "Entity"
type Entity struct {
	ValeurMorale uint8 // Valeur aléatoire qui va déterminer le groupe que l'entité rejoindra

	Position graphic.Vector2 // cordonnées de l'entité sur la map

	Voisins      []*Entity // Tableau des voisins avec lequels l'entité est en lien
	LiensVoisins []uint8   // Tableau des forces de liaisons avec les voisins
}

// Initialisation d'une instance entité
func NewEntity() *Entity {

	e := new(Entity)
	e.ValeurMorale = uint8(math.RandomRange(0, 255))

	return e
}

// Cette fonction permet de déplacer l'entité
func (e *Entity) Mouvement(newX float32, newY float32) {

	e.Position.X = newX
	e.Position.Y = newY

}

func (e *Entity) Update() {
	e.render()
}

// Cette fonction s'occupe d'afficher visuellement l'entité
func (e *Entity) render() {
	rl.DrawTextureEx(TextureEntite, rl.Vector2(e.Position), 0, SCALE, rl.White)
	e.renderLiens()

}

func (e *Entity) renderLiens() {
	for _, voisin := range e.Voisins {
		rl.DrawLineEx(voisin.Position.ToRaylibVector2(), e.Position.ToRaylibVector2(), 5, rl.Black)
	}
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
	if len(e.Voisins) < NB_VOISINS_MAX {
		forceDuLien := 128 - e.DistanceMorale(entiteVoisine)

		e.Voisins = append(e.Voisins, entiteVoisine)
		e.LiensVoisins = append(e.LiensVoisins, forceDuLien)

		entiteVoisine.Voisins = append(entiteVoisine.Voisins, e)
		entiteVoisine.LiensVoisins = append(entiteVoisine.LiensVoisins, forceDuLien)

		e.Position = entiteVoisine.Position
		e.Position.Y += 50
	} else {
		log.Fatal("l'entité a déjà atteint le nombre maximal de voisins")
	}

}
