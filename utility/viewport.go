package utility

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	lastCursorX, lastCursorY  float64
	renderWidth, renderHeight int
)

func InitViewport(width, height int) {
	renderWidth = width
	renderHeight = height
}

func rawCursorPosition() (float64, float64) {
	x, y := ebiten.CursorPosition()
	w, h, scale := Viewport()
	return (float64(x) - (float64(w-renderWidth)*scale)*0.5) / scale,
		(float64(y) - (float64(h-renderHeight)*scale)*0.5) / scale
}

func Viewport() (int, int, float64) {
	w, h := ebiten.WindowSize()
	return w, h, math.Floor(math.Min(float64(w/renderWidth), float64(h/renderHeight)))
}

func CursorPosition() (float64, float64) {
	x, y := rawCursorPosition()
	if x >= 0 && x <= float64(renderWidth) && y >= 0 && y <= float64(renderHeight) {
		lastCursorX = x
		lastCursorY = y
	}
	return lastCursorX, lastCursorY
}
