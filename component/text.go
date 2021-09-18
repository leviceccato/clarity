package component

import (
	"image/color"

	"golang.org/x/image/font"
)

type TextComponent struct {
	Lines      []string
	Color      color.NRGBA
	LineHeight int
	Font       font.Face
	IsCentered bool
}

func NewTextComponent(text string) *TextComponent {
	c := &TextComponent{}
	c.Lines = append(c.Lines, text)
	return c
}
