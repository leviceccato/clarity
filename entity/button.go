package entity

import (
	"fmt"
	"image/color"

	"github.com/leviceccato/clarity/component"
)

type ButtonEntityOptions struct {
	X, Y, Width, Height    float64
	Text, Image, Animation string
	Color                  color.NRGBA
}

func NewButtonEntity(options *ButtonEntityOptions) (*Entity, error) {
	e := NewEntity()
	e.Position = &component.PositionComponent{X: options.X, Y: options.Y}
	e.Size = &component.SizeComponent{Width: options.Width, Height: options.Height}
	e.Text = &component.TextComponent{Content: options.Text, Color: options.Color}
	appearance, err := component.NewAppearanceComponent(options.Image, options.Animation)
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %s", err)
	}
	e.Appearance = appearance
	return e, nil
}
