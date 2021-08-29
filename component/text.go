package component

import (
	"image/color"
)

type TextComponent struct {
	Text  string
	Color color.NRGBA
}

func NewTextComponent(text string) *TextComponent {
	return &TextComponent{
		Text: text,
	}
}
