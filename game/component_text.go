package game

import (
	"image/color"

	"golang.org/x/image/font"
)

type textComponentTransform int

const (
	textComponentTransformNone textComponentTransform = iota
	textComponentTransformUppercase
	textComponentTransformLowercase
	textComponentTransformTitlecase
)

// Text component
type textComponent struct {
	Lines         []textLine
	Color         color.NRGBA
	LineHeight    int
	Font          font.Face
	IsCentered    bool
	TextTransform textComponentTransform
	Padding       float64
}

func (textComponent) Name() string {
	return "text"
}

// For use in text component
type textLine struct {
	Content string
	X       float64
}
