package gameComponents

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// échelle qui correspond à la taille des entité (1 => 128; 0.5 => 64; ...)
const SCALE = 0.05

var (
	//texture utilisée pour afficher l'entité sur la fenêtre
	TextureEntite    rl.Texture2D
	ShowValeurMorale bool = false
)

// Définition de la classe "Entity"
type Entity struct {
	ValeurMorale uint8 // Valeur aléatoire qui va déterminer le groupe que l'entité rejoindra

	HitBox graphic.Circle
}

// Initialisation d'une instance entité
func NewEntity(position graphic.Vector2) *Entity {

	e := new(Entity)
	e.ValeurMorale = uint8(math.RandomRange(0, 255))
	e.HitBox = graphic.NewCircle(64*SCALE, position.X, position.Y)

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
		sum = sum.Add(entity.HitBox.CenterPosition.Scale(weight))
	}

	average := sum.Scale(1 / weightSum) // division par l'effectif pour faire la moyenne

	e.HitBox.CenterPosition = e.HitBox.CenterPosition.Add(average.Substract(e.HitBox.CenterPosition).Scale(0.01)) // déplacement vers cette position
}

// cette fonction est là pour éviter que les entités se chevauchent, en les "repoussant" assez loin
func (e *Entity) UnCollide(entities []*Entity) {

	for _, entity := range entities {
		if entity.HitBox.DetectCircleCollision(e.HitBox) && entity.HitBox.CenterPosition != e.HitBox.CenterPosition {
			e.HitBox.CenterPosition = entity.HitBox.CenterPosition.Add(e.HitBox.CenterPosition.Substract(entity.HitBox.CenterPosition).ScaleToNorm(entity.HitBox.Radius + e.HitBox.Radius))

		}

	}

}

func (e *Entity) Update(otherEntities []*Entity) {
	e.Move(otherEntities)
	e.UnCollide(otherEntities)
	e.render()

}

// Cette fonction s'occupe d'afficher visuellement l'entité
func (e *Entity) render() {
	rl.DrawTextureEx(TextureEntite, rl.Vector2(e.HitBox.CenterPosition.Substract(graphic.NewVector2(float32(TextureEntite.Width), float32(TextureEntite.Height)).Scale(0.5*SCALE))), 0, SCALE, rl.White)
	if ShowValeurMorale {
		e.HitBox.Fill(rl.NewColor(e.ValeurMorale, e.ValeurMorale, e.ValeurMorale, 255))

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
