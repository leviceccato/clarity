package game

import (
	"fmt"

	"github.com/leviceccato/clarity/engine"
)

type imageEntityOptions struct {
	x, y float64
	path string
}

func newImageEntity(g *Game, options *imageEntityOptions) (*engine.Entity, error) {
	e := g.NewEntity()

	e.AddComponent(&positionComponent{
		X: options.x,
		Y: options.y,
	})
	e.AddComponent(&sizeComponent{})

	imageAppearance, err := newAppearanceComponent(options.path, "")
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %s", err)
	}
	e.AddComponent(imageAppearance)

	return e, nil
}
