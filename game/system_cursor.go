package game

import (
	"github.com/leviceccato/clarity/engine"
)

func newCursorSystem(g *Game) *engine.System {
	s := engine.NewSystem("cursor", []string{
		"cursor",
		"position",
		"appearance",
	})

	s.Update = func() error {
		for _, entityId := range s.EntityIds {
			e := g.GetEntity(entityId)
			cursorX, cursorY := engine.CursorPosition()

			appearance, _ := engine.GetComponent(e, &appearanceComponent{})
			position, _ := engine.GetComponent(e, &positionComponent{})

			position.X = cursorX
			position.Y = cursorY

			if g.isCursorHovering && appearance.Sequence == "default" {
				appearance.Sequence = "pointer"
				continue
			}
			if !g.isCursorHovering && appearance.Sequence == "pointer" {
				appearance.Sequence = "default"
			}
		}

		return nil
	}

	return s
}
