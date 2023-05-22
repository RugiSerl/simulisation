package components

// Par Raphaël

import (
	"github.com/RugiSerl/simulisation/app/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// le slider, ou "curseur" est une barre verticale qui permet d'ajuster une valeur en déplaçant le rond qui est devant
type Slider struct {
	value *float32
	min   float32
	max   float32

	HoverState bool
	position   graphic.Vector2
	anchorX    int8
	anchorY    int8

	size graphic.Vector2
}

var (
	SliderRect graphic.Rect
)

// Initialisation du slider
func NewSlider(position graphic.Vector2, horizontalAnchor int8, verticalAnchor int8) *Slider {

	s := new(Slider)

	s.anchorX = horizontalAnchor
	s.anchorY = verticalAnchor

	s.position = position

	s.size = graphic.NewVector2(70, 20)

	return s

}

func (s *Slider) SetValue(value *float32, min float32, max float32) {
	s.value, s.min, s.max = value, min, max

}

func (s *Slider) GetSize() graphic.Vector2 {
	return s.size
}

// Fonction de mise à jour du slider
func (s *Slider) Update(containingRect graphic.Rect) {
	SliderRect = graphic.NewRectFromVector(graphic.GetRectCoordinatesWithAnchor(s.position, s.anchorX, s.anchorY, s.size, containingRect), s.size)

	s.handleInput()
	s.render()

}

// Fonction permettant de gérer les inputs du slider
func (s *Slider) handleInput() {
	s.HoverState = false
	if graphic.DetectRectCollision(SliderRect, graphic.GetMouseRect()) {
		s.HoverState = true
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			*s.value = s.min + (s.max-s.min)*(rl.GetMousePosition().X-SliderRect.X)/SliderRect.Width
		}

	}

}

// Fonction d'affichage du slider
func (s *Slider) render() {

	bar := graphic.GetInnerHorizontalrect(SliderRect, SliderRect.Height/3)

	ballXPosition := (*s.value-s.min)/(s.max-s.min)*SliderRect.Width + SliderRect.X
	ball := graphic.NewCircle(5, ballXPosition, SliderRect.Y+SliderRect.Height/2)
	ball.Fill(rl.Black)

	bar.Fill(rl.Black, 0)

}
