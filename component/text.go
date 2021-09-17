package component

import (
	"image/color"
)

type TextComponent struct {
	Content string
	Color   color.NRGBA
}

func NewTextComponent(text string) *TextComponent {
	return &TextComponent{
		Content: text,
	}
}
