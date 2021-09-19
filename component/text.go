package component

import (
	"image/color"

	"golang.org/x/image/font"
)

type TextLine struct {
	Content string
	X       float64
}

type TextComponent struct {
	Lines      []TextLine
	Color      color.NRGBA
	LineHeight int
	Font       font.Face
	IsCentered bool
	Padding    float64
}

func NewTextComponent(text string) *TextComponent {
	c := &TextComponent{}
	c.Lines = append(c.Lines, TextLine{Content: text})
	return c
}
