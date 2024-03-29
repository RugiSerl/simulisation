package Entity

// Par Raphaël et Gaël
import (
	"image/color"

	"github.com/RugiSerl/simulisation/app/Game/material"
	"github.com/RugiSerl/simulisation/app/graphic"
	"github.com/RugiSerl/simulisation/app/settings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// échelle qui correspond à la taille des entité (1 => 128px; 0.5 => 64px; ...)
const SCALE = 0.01

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
	Facing    float32
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

// Fonction qui met à jour les variables de l'entité
func (e *Entity) Update(otherEntities *[]*Entity, materials []material.IMaterial) {

	if settings.GameSettings.Gamerule.Move {
		e.Move(*otherEntities) //on déplace l'entité
	}

	if settings.GameSettings.Gamerule.ReactMaterial {
		e.UncollideMaterial(materials)
	}

	if settings.GameSettings.Gamerule.Uncollide {
		e.UnCollideAgressive(*otherEntities) //On évite que les entités se stackent
	}

	if settings.GameSettings.Gamerule.Reproduce {
		e.Reproduce(otherEntities) // On leur permet de se reproduire
	}

	if settings.GameSettings.Gamerule.UpdateAge {
		e.UpdateAge() // On met à jour l'age
	}

	if settings.GameSettings.Gamerule.Kill {
		e.Kill(otherEntities) // On leur permet de s'entretuer

	}
}

// Fonction qui définit la Hitbox de l'entité
func (e *Entity) GetPointCollision(point graphic.Vector2) bool {
	return e.HitBox.DetectPointCollision(point)

}

//--------------------------------------------------
//fonction d'affichage

// Cette fonction s'occupe d'afficher visuellement l'entité
func (e *Entity) Render() {

	if e.TimeAlive < settings.GameSettings.EntitySettings.MaximumAge {

		color := e.getColor()
		if settings.GameSettings.VisualSettings.GradientEntities {
			e.drawCircle(color)
		} else {
			rl.DrawTextureEx(TextureEntite, rl.Vector2(e.HitBox.CenterPosition.Substract(graphic.NewVector2(float32(TextureEntite.Width), float32(TextureEntite.Height)).Scale(0.5*SCALE))), 0, SCALE, rl.White)
			color.A = 128
			e.drawCircle(color)
		}

	}
}

func (e *Entity) drawCircle(color color.RGBA) {
	if settings.GameSettings.Mode3d {
		rl.DrawCylinderEx(rl.NewVector3(e.HitBox.CenterPosition.X, e.HitBox.CenterPosition.Y, 0), rl.NewVector3(e.HitBox.CenterPosition.X, e.HitBox.CenterPosition.Y, -1), e.HitBox.Radius, e.HitBox.Radius, 40, color)
	} else {
		e.HitBox.Fill(color)
	}
}

// retourne la couleur de l'entité en fonction de sa valeur morale
func (e *Entity) getColor() color.RGBA {
	return graphic.NewColorFromGradient(float64(e.ValeurMorale)/256.0*360.0, (float64(settings.GameSettings.EntitySettings.MaximumAge)-float64(e.TimeAlive))/float64(settings.GameSettings.EntitySettings.MaximumAge)/2)

}

// Cette fonction permet d'afficher le périmètre dans lequel l'entité peut voir les autres entités
func (e *Entity) RenderSensibilityZone() {
	circle := graphic.NewCircle(settings.GameSettings.EntitySettings.RadiusSensivity, e.HitBox.CenterPosition.X, e.HitBox.CenterPosition.Y)

	circle.DrawLines(e.getColor(), e.Facing)

}

//--------------------------------------------------
// la valeur morale est "cyclique", ce qui signifie que celle entre 4 et 254 est 5 par exemple

func (e *Entity) DistanceMorale(otherEntity *Entity) uint8 {
	distance := e.ValeurMorale - otherEntity.ValeurMorale
	if distance > 128 {
		return 255 - distance
	}
	return distance

}

// --------------------------------------------------
// Fonction qui élimine l'entité au bout d'un moment donné
func (e *Entity) UpdateAge() {
	e.TimeAlive += rl.GetFrameTime()
	if e.TimeAlive > settings.GameSettings.EntitySettings.MaximumAge {
		e.Dead = true
	}

}
