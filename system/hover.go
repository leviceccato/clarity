package system

import (
	"github.com/leviceccato/clarity/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type hoverSystem struct {
	System
}

func NewHoverSystem(s *game.State) *hoverSystem {
	sys := &hoverSystem{}
	sys.Components = []string{
		"Hover",
		"Position",
		"Size",
	}
	return sys
}

func (sys *hoverSystem) Load(s *game.State) {}

func (sys *hoverSystem) Update(s *game.State) {
	for _, e := range sys.Entities {
		// Check if mouse is within x and y ranges
		isHovered := s.CursorX >= e.Position.X && s.CursorX <= e.Position.X+e.Size.Width &&
			s.CursorY >= e.Position.Y && s.CursorY <= e.Position.Y+e.Size.Height

		// If no change, nothing to do
		hasHoverChanged := isHovered != e.Hover.IsHovered
		if !hasHoverChanged {
			return
		}

		e.Hover.IsHovered = isHovered
		s.SetIsCursorHovering(isHovered)

		// No animation, nothing to do
		if e.Appearance == nil {
			continue
		}

		// For animations allow toggling a 'hover' sequence if present
		_, hasHoverSequence := e.Appearance.Sequences["hover"]
		if !hasHoverSequence {
			continue
		}

		if isHovered {
			e.Appearance.PreviousSequence = e.Appearance.Sequence
			e.Appearance.Sequence = "hover"
			continue
		}
		e.Appearance.PreviousSequence = "hover"
		e.Appearance.Sequence = e.Appearance.PreviousSequence
	}
}

func (sys *hoverSystem) Draw(s *game.State, screen *ebiten.Image) {}

func (sys *hoverSystem) Enter(s *game.State) {}

func (sys *hoverSystem) Exit(s *game.State) {}
