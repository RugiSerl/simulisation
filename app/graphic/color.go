package graphic

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewColorAbsvalue(offset uint8) color.RGBA {
	var r float64 = -math.Abs(3*float64(offset)) + 255
	var g float64 = -math.Abs(3*float64(offset)-255) + 255
	var b float64 = -math.Abs(3*float64(offset)-255*2) + 255

	return rl.NewColor(uint8(r), uint8(g), uint8(b), 255)
}
