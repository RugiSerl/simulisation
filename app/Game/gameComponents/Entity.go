package gameComponents

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
}

// Initialisation d'une instance entité
func NewEntity(position graphic.Vector2) *Entity {

	e := new(Entity)
	e.ValeurMorale = uint8(math.RandomRange(0, 255))
	e.Position = position

	return e
}

// Cette fonction permet de déplacer l'entité et de rapprocher l'entité des entités similaires.
// Elle choisit une destination qui est la 'moyenne' des position pondérée à l'aide des 'distances morales'
func (e *Entity) Move(otherEntities []*Entity) {

	var sum graphic.Vector2 = graphic.NewVector2(0, 0)
	var weight float32
	var weightSum float32 = 0

	for _, entity := range otherEntities {
		weight = float32(e.DistanceMorale(entity)) / 255
		weightSum += weight
		sum = sum.Add(entity.Position.Scale(weight))
	}

	average := sum.Scale(1 / weightSum) // division par l'effectif pour faire la moyenne

	e.Position = e.Position.Add(average.Substract(e.Position).Scale(0.01)) // déplacement vers cette position
}

// cette fonction est là pour éviter que les entités se chevauchent, en les "repoussant" assez loin
func (e *Entity) UnCollide() {

}

func (e *Entity) Update(otherEntities []*Entity) {
	e.render()
	e.Move(otherEntities)
}

// Cette fonction s'occupe d'afficher visuellement l'entité
func (e *Entity) render() {
	rl.DrawTextureEx(TextureEntite, rl.Vector2(e.Position.Substract(graphic.NewVector2(float32(TextureEntite.Width), float32(TextureEntite.Height)).Scale(0.25))), 0, SCALE, rl.White)

}

// la valeur morale est "cyclique", ce qui signifie que celle entre 5 et 254 est 6 par exemple
func (e *Entity) DistanceMorale(otherEntity *Entity) uint8 {
	distance := e.ValeurMorale - otherEntity.ValeurMorale
	if distance > 128 {
		return 255 - distance
	}
	return distance

}
