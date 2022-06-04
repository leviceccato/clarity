package game

import (
	"github.com/leviceccato/clarity/engine"
)

func newPlayableSystem(g *Game) *engine.System {
	s := engine.NewSystem("playable", []string{
		"playable",
	})

	s.Update = func() error {
		for _, entityId := range s.EntityIds {
			e := g.GetEntity(entityId)

			position, _ := engine.GetComponent(e, &positionComponent{})

			if g.inputs[commandMoveUp] != nil {
				position.Y -= 3
			}
			if g.inputs[commandMoveLeft] != nil {
				position.X -= 3
			}
			if g.inputs[commandMoveRight] != nil {
				position.X += 3
			}
			if g.inputs[commandMoveDown] != nil {
				position.Y += 3
			}
			if g.inputs[commandToggleMenu] != nil {
				return g.quit()
			}
			toggleDebug := g.inputs[commandToggleDebug]
			if (toggleDebug != nil) && !toggleDebug.isStart {
				g.toggleDebug()
			}
		}

		return nil
	}

	return s
}
