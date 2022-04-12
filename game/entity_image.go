package game

import (
	"fmt"

	"github.com/leviceccato/clarity/engine"
)

func newImageEntity(g *Game, path string) (*engine.Entity, error) {
	e := g.NewEntity()

	imageAppearance, err := newAppearanceComponent(path, "")
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %s", err)
	}
	e.AddComponent(imageAppearance)

	return e, nil
}
