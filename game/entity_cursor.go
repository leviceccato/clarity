package game

import (
	"fmt"

	"github.com/leviceccato/clarity/engine"
)

func newCursorEntity(g *Game) (*engine.Entity, error) {
	e := g.NewEntity()

	e.AddComponent(&cursorComponent{})

	x, y := engine.CursorPosition()
	e.AddComponent(&positionComponent{
		X: x,
		Y: y,
		Z: 100,
	})

	e.AddComponent(&sizeComponent{
		Width:  10,
		Height: 10,
	})

	cursorAppearance, err := newAppearanceComponent("sprite/cursor.png", "sprite/cursor.json")
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %s", err)
	}
	e.AddComponent(cursorAppearance)

	return e, nil
}
