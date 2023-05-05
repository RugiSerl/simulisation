package Entity

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// échelle qui correspond à la taille des entité (1 => 128px; 0.5 => 64px; ...)
const SCALE = 0.01

// vitesse à laquelle se déplacent les entités
const SPEED = 20

// écart de différence morale maximum entre une entité et son enfant
const CHILD_MAXIMUM_DIFFERENCE = 5

// rayon dans lequel une entité "voit" les autres entités
const RADIUS_SENSIVITY = 0.1 * 100 //px

// durée de vie d'une entité, en s
const MAXIMUM_AGE = 10

var (
	//texture utilisée pour afficher l'entité sur la fenêtre
	TextureEntite rl.Texture2D

	//bool utilisé pour savoir si l'on affiche une représentation graphique de la valeur morale de l'entité
	ShowEntityRadiusVision = false
)

// Définition de la classe "Entity"
type Entity struct {
	ValeurMorale uint8 // Valeur aléatoire qui va déterminer le groupe que l'entité rejoindra

	HitBox    graphic.Circle
	ID        int
	Dead      bool
	TimeAlive float32
}

// Initialisation d'une instance entité
func NewEntity(position graphic.Vector2, id int, valeurMorale uint8) *Entity {

	e := new(Entity)
	e.ValeurMorale = valeurMorale
	e.HitBox = graphic.NewCircle(64*SCALE, position.X, position.Y)
	e.ID = id
	e.TimeAlive = 0
	e.Dead = false

	return e
}

func (e *Entity) Update(otherEntities *[]*Entity) {
	e.MoveToClosestNeighbour(*otherEntities) //on déplace l'entité

	e.UnCollideAgressive(*otherEntities) //On évite que les entités se stackent
	e.Reproduce(otherEntities)
	e.render() //on affiche l'entité
	e.UpdateAge()

}

//--------------------------------------------------
//fonction d'affichage

// Cette fonction s'occupe d'afficher visuellement l'entité
func (e *Entity) render() {

	if ShowEntityRadiusVision {
		rl.DrawCircleV(rl.Vector2(e.HitBox.CenterPosition), RADIUS_SENSIVITY, rl.NewColor(0, 0, 0, 100))
	}
	rl.DrawTextureEx(TextureEntite, rl.Vector2(e.HitBox.CenterPosition.Substract(graphic.NewVector2(float32(TextureEntite.Width), float32(TextureEntite.Height)).Scale(0.5*SCALE))), 0, SCALE, rl.White)
	if settings.GameSettings.VisualSettings.GradientEntities {
		e.HitBox.Fill(graphic.NewColorFromGradient(float64(e.ValeurMorale)/256.0*360.0, (MAXIMUM_AGE-float64(e.TimeAlive))/MAXIMUM_AGE/2))

	}

}

//--------------------------------------------------
// la valeur morale est "cyclique", ce qui signifie que celle entre 5 et 254 est 6 par exemple

func (e *Entity) DistanceMorale(otherEntity *Entity) uint8 {
	distance := e.ValeurMorale - otherEntity.ValeurMorale
	if distance > 128 {
		return 255 - distance
	}
	return distance

}

//--------------------------------------------------
// fonction qui élimine l'entité au bout d'un moment donné

func (e *Entity) UpdateAge() {
	e.TimeAlive += rl.GetFrameTime()
	if e.TimeAlive > MAXIMUM_AGE {
		e.Dead = true
	}

}
