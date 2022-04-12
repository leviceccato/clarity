package game

import (
	"fmt"

	"github.com/leviceccato/clarity/engine"
)

func newStartWorld(g *Game) *engine.World {
	w := engine.NewWorld("start", []string{
		"appearance",
		"draw",
		"animation",
		"input",
		"event",
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

		g.AddEntities(w,
			player,
		)

		return nil
	}

	return w
}
