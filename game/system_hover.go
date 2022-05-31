package game

import (
	"github.com/leviceccato/clarity/engine"
)

func newHoverSystem(g *Game) *engine.System {
	s := engine.NewSystem("hover", []string{
		"hoverable",
		"position",
		"size",
	})

	s.Update = func() error {
		for _, entityId := range s.EntityIds {
			e := g.GetEntity(entityId)

			cursorX, cursorY := engine.CursorPosition()

			position, _ := engine.GetComponent(e, &positionComponent{})
			size, _ := engine.GetComponent(e, &sizeComponent{})
			hoverable, _ := engine.GetComponent(e, &hoverableComponent{})
			appearance, hasAppearance := engine.GetComponent(e, &appearanceComponent{})

			// Check if mouse is within x and y ranges
			isHovered := cursorX >= position.X &&
				cursorX <= position.X+size.Width &&
				cursorY >= position.Y &&
				cursorY <= position.Y+size.Height

			// If no change, nothing to do
			hasHoverChanged := isHovered != hoverable.IsHovered
			if !hasHoverChanged {
				continue
			}

			hoverable.IsHovered = isHovered

			// No animation, nothing to do
			if !hasAppearance {
				continue
			}

			// For animations allow toggling a 'hover' sequence if present
			_, hasHoverSequence := appearance.Sequences["hover"]
			if !hasHoverSequence {
				continue
			}

			if isHovered {
				appearance.PreviousSequence = appearance.Sequence
				appearance.Sequence = "hover"
				continue
			}
			appearance.Sequence = appearance.PreviousSequence
			appearance.PreviousSequence = "hover"
		}

		return nil
	}

	return s
}
