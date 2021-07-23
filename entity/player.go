package entity

import (
	"fmt"
	_ "image/png"

	"github.com/leviceccato/clarity/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var image *ebiten.Image

func NewPlayerEntity() (*Entity, error) {
	e := NewEntity()
	e.Position = &component.Position{X: 10, Y: 10}
	e.Size = &component.Size{Width: 10, Height: 10}
	if image == nil {
		var err error
		image, _, err = ebitenutil.NewImageFromFile("assets/test3.png")
		if err != nil {
			return e, fmt.Errorf("loading player image: %s", err)
		}
	}
	e.Appearance = &component.Appearance{Image: image}
	return e, nil
}
