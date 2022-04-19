package game

import (
	"github.com/leviceccato/clarity/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

func newCursorSystem(g *Game) *engine.System {
	s := engine.NewSystem("cursor", []string{
		"cursor",
		"position",
		"appearance",
	})

	s.Enter = func() {
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
	}

	s.Update = func() error {
		cursorX, cursorY := engine.CursorPosition()

		for _, entityId := range s.EntityIds {
			e := g.GetEntity(entityId)

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

	s.Exit = func() {
		ebiten.SetCursorMode(ebiten.CursorModeVisible)
	}

	return s
}
