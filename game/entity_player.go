package game

import (
	"fmt"

	"github.com/leviceccato/clarity/engine"
)

func newPlayerEntity(g *Game) (*engine.Entity, error) {
	e := g.NewEntity()

	e.AddComponent(&playableComponent{})

	e.AddComponent(&positionComponent{
		X: 10,
		Y: 10,
	})

	e.AddComponent(&sizeComponent{
		Width:  10,
		Height: 10,
	})

	playerAppearance, err := newAppearanceComponent("sprite/player.png", "sprite/player.json")
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %s", err)
	}
	e.AddComponent(playerAppearance)

	return e, nil
}
