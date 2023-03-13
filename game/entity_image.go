package game

import (
	"fmt"

	"github.com/leviceccato/clarity/engine"
)

type imageEntityOptions struct {
	x, y float64
	z    int
	path string
}

func newImageEntity(g *Game, options *imageEntityOptions) (*engine.Entity, error) {
	e := g.NewEntity()

	e.AddComponent(&positionComponent{
		X: options.x,
		Y: options.y,
		Z: options.z,
	})
	e.AddComponent(&sizeComponent{})

	imageAppearance, err := newAppearanceComponent(options.path, "")
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %s", err)
	}
	e.AddComponent(imageAppearance)

	return e, nil
}
