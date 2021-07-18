package entity

import (
	"fmt"
	_ "image/png"

	"github.com/leviceccato/clarity/component"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func NewPlayerEntity() (*Entity, error) {
	e := NewEntity()
	e.Position = &component.Position{X: 10, Y: 10}
	e.Size = &component.Size{Width: 10, Height: 10}
	image, _, err := ebitenutil.NewImageFromFile("assets/test.png")
	if err != nil {
		return e, fmt.Errorf("loading player image: %s", err)
	}
	e.Appearance = &component.Appearance{Image: image}
	e.Animation = &component.Animation{}
	return e, nil
}
