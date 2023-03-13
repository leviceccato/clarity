package game

import (
	"fmt"

	"github.com/leviceccato/clarity/engine"
)

func newStartWorld(g *Game) *engine.World {
	w := engine.NewWorld("start", []string{
		"cursor",
		"draw",
		"animation",
		"input",
		"playable",
	})

	w.Enter = func() error {
		player, err := newPlayerEntity(g)
		if err != nil {
			return fmt.Errorf("creating player entity: %w", err)
		}
		playerPosition, _ := engine.GetComponent(player, &positionComponent{})
		playerPosition.X = float64(g.RenderWidth) / 2
		playerPosition.Y = float64(g.RenderHeight) / 2

		cursor, err := newCursorEntity(g)
		if err != nil {
			return fmt.Errorf("creating cursor entity: %w", err)
		}

		menu, err := newMenuAssemblage(g)
		if err != nil {
			return fmt.Errorf("creating menu assemblage: %w", err)
		}

		g.AddEntities(w, append([]*engine.Entity{
			cursor,
			player,
		}, menu...)...)

		return nil
	}

	return w
}
