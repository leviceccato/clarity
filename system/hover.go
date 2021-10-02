package system

import (
	"github.com/leviceccato/clarity/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type hoverSystem struct {
	system
	state SystemState
}

func NewHoverSystem(state SystemState) *hoverSystem {
	s := &hoverSystem{}
	s.state = state
	s.components = []string{
		"Hover",
		"Position",
		"Size",
	}
	return s
}

func (s *hoverSystem) Load() {}

func (s *hoverSystem) Update() {
	x, y := ebiten.CursorPosition()
	mouseX := float64(x)
	mouseY := float64(y)
	var (
		hasHoverChanged, isHovered, hasHoverSequence bool
		e                                            *entity.Entity
	)
	for _, e = range s.entities {
		// Check if mouse is within x and y ranges
		isHovered = mouseX >= e.Position.X && mouseX <= e.Position.X+e.Size.Width &&
			mouseY >= e.Position.Y && mouseY <= e.Position.Y+e.Size.Height
		hasHoverChanged = isHovered != e.Hover.IsHovered
		if hasHoverChanged {
			e.Hover.IsHovered = isHovered
			// For animations allow toggling a 'hover' sequence
			if e.Appearance == nil {
				continue
			}
			_, hasHoverSequence = e.Appearance.Sequences["hover"]
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
}

func (s *hoverSystem) Draw(screen *ebiten.Image) {}

func (s *hoverSystem) Enter() {}

func (s *hoverSystem) Exit() {}
