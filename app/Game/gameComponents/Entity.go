package gameComponents

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// échelle qui correspond à la taille des entité (1 => 128px; 0.5 => 64px; ...)
const SCALE = 0.01

// rayon dans lequel une entité "voit" les autres entités
const RADIUS_SENSIVITY = 0.1 * 30 //px

var (
	//texture utilisée pour afficher l'entité sur la fenêtre
	TextureEntite rl.Texture2D

	//bool utilisé pour savoir si l'on affiche une représentation graphique de la valeur morale de l'entité
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

func (e *Entity) Update(otherEntities []*Entity) {
	e.Move(otherEntities)             //on déplace l'entité
	e.UnCollidePassive(otherEntities) //On évite que les entités se stack
	e.render()                        //on affiche l'entité

}

//--------------------------------------------------
//fonctions de déplacement

// Cette fonction permet de déplacer l'entité et de rapprocher l'entité des entités similaires.
// Elle choisit une destination qui est la 'moyenne' des position pondérée à l'aide des 'distances morales'
// Elle ne peut "voir" que les autres entités qui sont dans un certain rayon de cette dernière (RADIUS_SENSIVITY)
func (e *Entity) Move(otherEntities []*Entity) {

	var sum graphic.Vector2 = graphic.NewVector2(0, 0)
	var weight float32
	var weightSum float32 = 0

	for _, entity := range otherEntities {
		if entity.HitBox.CenterPosition.Substract(e.HitBox.CenterPosition).GetNorm() < RADIUS_SENSIVITY {
			weight = float32(e.DistanceMorale(entity)) / 255
			weight = float32(math.Exp(float64(weight)))
			weightSum += weight
			sum = sum.Add(entity.HitBox.CenterPosition.Scale(weight))
		}

	}
	if weightSum != 0 { // éviter la division par 0, si jamais l'entité n'a aucune entité dans son rayon RADIUS_SENSIVITY
		average := sum.Scale(1 / weightSum) // division par l'effectif pour faire la moyenne

		e.HitBox.CenterPosition = e.HitBox.CenterPosition.Add(average.Substract(e.HitBox.CenterPosition).Scale(0.01)) // déplacement vers cette position

	}

}

// --------------------------------------------------
// fonctions de détection de collisions

// L'entité balaye toutes les autres entités sur son chemin
func (e *Entity) UnCollideAgressive(entities []*Entity) {

	for _, entity := range entities {

		if entity.HitBox.DetectCircleCollision(e.HitBox) && e.HitBox.CenterPosition != entity.HitBox.CenterPosition {
			entity.HitBox.CenterPosition = e.HitBox.CenterPosition.Add(entity.HitBox.CenterPosition.Substract(e.HitBox.CenterPosition).ScaleToNorm(entity.HitBox.Radius + e.HitBox.Radius))
		}
	}
}

// l'entité se déplace lorsqu'elle est en collision avec une autre
func (e *Entity) UnCollidePassive(entities []*Entity) {
	for _, entity := range entities {

		if entity.HitBox.DetectCircleCollision(e.HitBox) && e.HitBox.CenterPosition != entity.HitBox.CenterPosition {
			e.HitBox.CenterPosition = entity.HitBox.CenterPosition.Add(e.HitBox.CenterPosition.Substract(entity.HitBox.CenterPosition).ScaleToNorm(entity.HitBox.Radius + e.HitBox.Radius))
		}
	}
}

//--------------------------------------------------
//fonction d'affichage

// Cette fonction s'occupe d'afficher visuellement l'entité
func (e *Entity) render() {
	rl.DrawTextureEx(TextureEntite, rl.Vector2(e.HitBox.CenterPosition.Substract(graphic.NewVector2(float32(TextureEntite.Width), float32(TextureEntite.Height)).Scale(0.5*SCALE))), 0, SCALE, rl.White)
	if ShowValeurMorale {
		e.HitBox.Fill(graphic.NewColorFromGradient(float64(e.ValeurMorale) / 255.0 * 360.0))

	}
}

//--------------------------------------------------
//autre

// la valeur morale est "cyclique", ce qui signifie que celle entre 5 et 254 est 6 par exemple
func (e *Entity) DistanceMorale(otherEntity *Entity) uint8 {
	distance := e.ValeurMorale - otherEntity.ValeurMorale
	if distance > 128 {
		return 255 - distance
	}
	return distance

}
