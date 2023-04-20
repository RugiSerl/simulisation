package graphic

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Rect struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

func NewRect(x float32, y float32, width float32, height float32) Rect {
	r := Rect{}
	r.X, r.Y, r.Width, r.Height = x, y, width, height
	return r
}

func NewRectFromVector(position Vector2, size Vector2) Rect {

	return Rect{X: position.X, Y: position.Y, Width: size.X, Height: size.Y}

}

func GetWindowRect() Rect {
	return NewRect(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()))
}

func GetRectAdjustedToWindow(rectRatio float32) Rect {

	winW, winH := float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())
	winRatio := winW / winH

	r := Rect{}

	if rectRatio > winRatio { //change Height
		r.Width = winW
		r.Height = winW / rectRatio
		r.X = 0
		r.Y = winH/2 - r.Height/2

	} else if rectRatio < winRatio {
		r.Height = winH
		r.Width = winH * rectRatio
		r.Y = 0
		r.X = winW/2 - r.Width/2

	} else {
		r.X = 0
		r.Y = 0
		r.Height = winH
		r.Width = winW
	}

	return r

}

// Get rect centered to window center
func GetRectFromWindowCenter(width float32, height float32) Rect {

	r := Rect{}
	r.Width, r.Height = width, height
	r.X = float32(rl.GetScreenWidth())/2 - width/2
	r.Y = float32(rl.GetScreenHeight())/2 - height/2

	return r

}

// Get a rect within another, with padding
func GetInnerRect(sourceRect Rect, padding float32) Rect {
	sourceRect.X += padding
	sourceRect.Y += padding
	sourceRect.Width -= padding * 2
	sourceRect.Height -= padding * 2
	return sourceRect
}

// convert to raylib's Rect object
func (r Rect) ToRaylibRect() rl.Rectangle {
	return rl.NewRectangle(r.X, r.Y, r.Width, r.Height)
}

// draw the rectangle
func (r Rect) Fill(color color.RGBA, roundness float32) {
	rectangle := rl.NewRectangle(r.X, r.Y, r.Width, r.Height)

	rl.DrawRectangleRounded(rectangle, roundness, 5, color)

}

// draw the lines of a rectangle
func (r Rect) DrawLines(color color.RGBA, roundness float32, thickness float32) {
	rectangle := rl.NewRectangle(r.X, r.Y, r.Width, r.Height)

	rl.DrawRectangleRoundedLines(rectangle, roundness, 5, thickness, color)
}

// return true if the 2 rect overlap, else false
func DetectRectCollision(rect1 Rect, rect2 Rect) bool {
	if rect1.X+rect1.Width >= rect2.X && rect1.X <= rect2.X+rect2.Width && rect1.Y+rect1.Height >= rect2.Y && rect1.Y <= rect2.Y+rect2.Height {
		return true
	} else {
		return false
	}
}

// return the rect of the mouse, used for collisions
func GetMouseRect() Rect {
	return NewRect(rl.GetMousePosition().X, rl.GetMousePosition().Y, 1, 1)

}
