package game

import (
	"github.com/leviceccato/clarity/engine"
)

func newPlayableSystem(g *Game) *engine.System {
	s := engine.NewSystem("playable", []string{
		"playable",
	})

	s.Update = func() error {
		// for _, entityId := range s.EntityIds {
		// 	e := g.GetEntity(entityId)
		// 	// if c[component.ControlUp] != nil {
		// 	// 	e.Position.Y -= 3
		// 	// }
		// 	// if c[component.ControlLeft] != nil {
		// 	// 	e.Position.X -= 3
		// 	// }
		// 	// if c[component.ControlRight] != nil {
		// 	// 	e.Position.X += 3
		// 	// }
		// 	// if c[component.ControlDown] != nil {
		// 	// 	e.Position.Y += 3
		// 	// }
		// 	// if c[component.ControlMenu] != nil {
		// 	// 	s.AddEvent(activateWorldsEvent{
		// 	// 		names: []string{"title"},
		// 	// 	})
		// 	// }
		// }

		return nil
	}

	return s
}
