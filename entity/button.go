package entity

import (
	"fmt"

	"github.com/leviceccato/clarity/component"
)

type ButtonEntityOptions struct {
	X, Y, Width, Height    float64
	Text, Image, Animation string
}

func NewButtonEntity(options *ButtonEntityOptions) (*Entity, error) {
	e := NewEntity()
	e.Position = &component.PositionComponent{X: options.X, Y: options.Y}
	e.Size = &component.SizeComponent{Width: options.Width, Height: options.Height}
	e.Text = component.NewTextComponent(options.Text)
	appearance, err := component.NewAppearanceComponent(options.Image, options.Animation)
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %s", err)
	}
	e.Appearance = appearance
	return e, nil
}
