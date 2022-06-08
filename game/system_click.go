package game

import (
	"github.com/leviceccato/clarity/engine"
	"github.com/leviceccato/clarity/util"
)

func newClickSystem(g *Game) *engine.System {
	s := engine.NewSystem("click", []string{
		"clickable",
		"position",
		"size",
	})

	s.Update = func() error {

		click := g.inputs[commandClick]

		// Don't do anything if not clicking
		if click == nil {
			return nil
		}

		// Only respond on mouse up
		if click.progress != inputEnd {
			return nil
		}

		for _, entityId := range s.EntityIds {
			e := g.GetEntity(entityId)

			position, _ := engine.GetComponent(e, &positionComponent{})
			size, _ := engine.GetComponent(e, &sizeComponent{})

			// Check if mouse is within x and y ranges
			isClickedOn := util.IsWithinRect(
				click.x, click.y,
				position.X, position.Y,
				size.Width, size.Height,
			)

			// Didn't click on this entity
			if !isClickedOn {
				continue
			}

			clickable, _ := engine.GetComponent(e, &clickableComponent{})

			clickable.action()
			break
		}

		return nil
	}

	return s
}
