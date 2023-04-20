package components

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// échelle qui correspond à la taille des entité (1 => 128px; 0.5 => 64px; ...)
const SCALE = 0.01

// vitesse à laquelle se déplacent les entités
const SPEED = 20

const MAXIMUM_AGE_ENTITY = 60

// écart de différence morale maximum entre une entité et son enfant
const CHILD_MAXIMUM_DIFFERENCE = 5

// rayon dans lequel une entité "voit" les autres entités
const RADIUS_SENSIVITY = 0.1 * 100 //px

var (
	//texture utilisée pour afficher l'entité sur la fenêtre
	TextureEntite rl.Texture2D

	//bool utilisé pour savoir si l'on affiche une représentation graphique de la valeur morale de l'entité
	ShowValeurMorale       bool = false
	ShowEntityRadiusVision      = false
)

// Définition de la classe "Entity"
type Entity struct {
	ValeurMorale uint8 // Valeur aléatoire qui va déterminer le groupe que l'entité rejoindra

	HitBox graphic.Circle
	ID     int
}

// Initialisation d'une instance entité
func NewEntity(position graphic.Vector2, id int, valeurMorale uint8) *Entity {

	e := new(Entity)
	e.ValeurMorale = valeurMorale
	e.HitBox = graphic.NewCircle(64*SCALE, position.X, position.Y)
	e.ID = id

	return e
}

func (e *Entity) Update(otherEntities *[]*Entity) {
	e.MoveToWeightedAverage(*otherEntities) //on déplace l'entité

	e.UnCollidePassive(*otherEntities) //On évite que les entités se stackent
	e.Reproduce(otherEntities)
	e.render() //on affiche l'entité

}

//--------------------------------------------------
//fonctions de déplacement

// Cette fonction permet de déplacer l'entité et de rapprocher l'entité des entités similaires.
// Elle choisit une destination qui est la 'moyenne' des position pondérée à l'aide des 'distances morales'
// Elle ne peut "voir" que les autres entités qui sont dans un certain rayon de cette dernière (RADIUS_SENSIVITY)
func (e *Entity) MoveToWeightedAverage(otherEntities []*Entity) {

	var sum graphic.Vector2 = graphic.NewVector2(0, 0)
	var weight float32
	var weightSum float32 = 0

	for _, entity := range otherEntities {
		if entity.ID != e.ID {
			if entity.HitBox.CenterPosition.Substract(e.HitBox.CenterPosition).GetNorm() < RADIUS_SENSIVITY {
				//weight = float32(e.DistanceMorale(entity)) / 255
				weight = 1
				weightSum += weight
				sum = sum.Add(entity.HitBox.CenterPosition.Scale(weight))
			}
		}

	}
	if weightSum != 0 { // éviter la division par 0, si jamais l'entité n'a aucune entité dans son rayon RADIUS_SENSIVITY
		average := sum.Scale(1 / weightSum) // division par l'effectif pour faire la moyenne

		e.GotoDivide(average) // déplacement vers cette position

	}

}

// aller à un point de manière linéaire
func (e *Entity) GotoLinear(point graphic.Vector2) {

	if e.HitBox.CenterPosition.Substract(point).GetNorm() > SPEED*rl.GetFrameTime() {
		e.HitBox.CenterPosition = e.HitBox.CenterPosition.Add(point.Substract(e.HitBox.CenterPosition).ScaleToNorm(SPEED * rl.GetFrameTime()))

	} else {
		e.HitBox.CenterPosition = point
	}

}

// aller à un point en divisant la distance par une certaine valeur
func (e *Entity) GotoDivide(point graphic.Vector2) {
	e.HitBox.CenterPosition = e.HitBox.CenterPosition.Add(point.Substract(e.HitBox.CenterPosition).Scale(.1))
}

// --------------------------------------------------
// fonctions de détection de collisions

// L'entité balaye toutes les autres entités sur son chemin
func (e *Entity) UnCollideAgressive(entities []*Entity) {

	for _, entity := range entities {
		if entity.ID != e.ID {
			if entity.HitBox.DetectCircleCollision(e.HitBox) && e.HitBox.CenterPosition != entity.HitBox.CenterPosition {
				entity.HitBox.CenterPosition = e.HitBox.CenterPosition.Add(entity.HitBox.CenterPosition.Substract(e.HitBox.CenterPosition).ScaleToNorm(entity.HitBox.Radius + e.HitBox.Radius))
			}
		}

	}
}

// l'entité se déplace lorsqu'elle est en collision avec une autre
func (e *Entity) UnCollidePassive(entities []*Entity) {
	for _, entity := range entities {
		if entity.ID != e.ID {
			if entity.HitBox.DetectCircleCollision(e.HitBox) && e.HitBox.CenterPosition != entity.HitBox.CenterPosition {
				e.HitBox.CenterPosition = entity.HitBox.CenterPosition.Add(e.HitBox.CenterPosition.Substract(entity.HitBox.CenterPosition).ScaleToNorm(entity.HitBox.Radius + e.HitBox.Radius))
			}
		}

	}
}

// --------------------------------------------------
// fonction pour faire se reproduire les entités
// les nouvelles cellules sont proches "moralement" de celles qui les ont engendré
func (e *Entity) Reproduce(othersEntities *[]*Entity) {
	var entityClose int32 = 0
	for _, entity := range *othersEntities {
		if entity.HitBox.CenterPosition.Substract(e.HitBox.CenterPosition).GetNorm() < RADIUS_SENSIVITY && entity.ID != e.ID {
			entityClose += 1
		}
	}

	if entityClose > 5 {
		entityClose = 5
	}

	var probability float64 = float64(entityClose) / 10000

	if math.RandomProbability(probability) {
		*othersEntities = append(*othersEntities, NewEntity(e.HitBox.CenterPosition.Add(graphic.NewVector2(0, 1)), len(*othersEntities), uint8(math.RandomRange(int(e.ValeurMorale)-CHILD_MAXIMUM_DIFFERENCE, (int(e.ValeurMorale)+CHILD_MAXIMUM_DIFFERENCE)))))
	}

}

//--------------------------------------------------
//fonction d'affichage

// Cette fonction s'occupe d'afficher visuellement l'entité
func (e *Entity) render() {

	if ShowEntityRadiusVision {
		rl.DrawCircleV(rl.Vector2(e.HitBox.CenterPosition), RADIUS_SENSIVITY, rl.NewColor(0, 0, 0, 100))
	}
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
