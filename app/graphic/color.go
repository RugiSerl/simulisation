package graphic

import (
	"errors"
	"fmt"
	"image/color"
	"log"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	ErrOutOfRange = errors.New("colorconv: inputs out of range")
)

func NewColorFromGradient(offset float64, brightness float64) color.RGBA {
	r, g, b, err := HSLToRGB(offset, 1, brightness)

	if err != nil {
		fmt.Println(offset, 1, brightness)
		log.Fatal(err)

	}

	return rl.NewColor(uint8(r), uint8(g), uint8(b), 255)
}

func HSLToRGB(h, s, l float64) (r, g, b uint8, err error) {
	if h < 0 || h >= 360 ||
		s < 0 || s > 1 ||
		l < 0 || l > 1 {
		return 0, 0, 0, ErrOutOfRange
	}
	// When 0 ≤ h < 360, 0 ≤ s ≤ 1 and 0 ≤ l ≤ 1:
	C := (1 - math.Abs((2*l)-1)) * s
	X := C * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - (C / 2)
	var Rnot, Gnot, Bnot float64

	switch {
	case 0 <= h && h < 60:
		Rnot, Gnot, Bnot = C, X, 0
	case 60 <= h && h < 120:
		Rnot, Gnot, Bnot = X, C, 0
	case 120 <= h && h < 180:
		Rnot, Gnot, Bnot = 0, C, X
	case 180 <= h && h < 240:
		Rnot, Gnot, Bnot = 0, X, C
	case 240 <= h && h < 300:
		Rnot, Gnot, Bnot = X, 0, C
	case 300 <= h && h < 360:
		Rnot, Gnot, Bnot = C, 0, X
	}
	r = uint8(math.Round((Rnot + m) * 255))
	g = uint8(math.Round((Gnot + m) * 255))
	b = uint8(math.Round((Bnot + m) * 255))
	return r, g, b, nil
}
