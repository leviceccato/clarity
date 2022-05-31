package game

import (
	"image/color"

	"golang.org/x/image/font"
)

// Text component
type textComponent struct {
	Lines         []textLine
	Color         color.NRGBA
	LineHeight    int
	Font          font.Face
	IsCentered    bool
	TextTransform string
	Padding       float64
}

func (_ textComponent) Name() string {
	return "text"
}

// For use in text component
type textLine struct {
	Content string
	X       float64
}
