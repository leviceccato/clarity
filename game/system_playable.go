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

			if g.hasInput(commandMoveUp) {
				position.Y -= 3
			}
			if g.hasInput(commandMoveLeft) {
				position.X -= 3
			}
			if g.hasInput(commandMoveRight) {
				position.X += 3
			}
			if g.hasInput(commandMoveDown) {
				position.Y += 3
			}
			if g.hasInput(commandToggleMenu) {
				g.ActivateWorlds("title")
			}
			if g.isInputEnding(commandToggleDebug) {
				g.toggleDebug()
			}
		}

		return nil
	}

	return s
}
