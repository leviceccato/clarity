package system

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type hoverSystem struct {
	system
}

func NewHoverSystem() *hoverSystem {
	s := &hoverSystem{}
	s.components = []string{
		"Hover",
		"Position",
		"Size",
	}
	return s
}

func (s *hoverSystem) Load() {}

func (s *hoverSystem) Update() {
	mouseX, mouseY := ebiten.CursorPosition()
	var hasHoverChanged, isHovered, hasHoverSequence bool
	for _, e := range s.entities {
		// Check if mouse is within x and y ranges
		isHovered = mouseX >= int(e.Position.X) && mouseX <= int(e.Position.X)+int(e.Size.Width) &&
			mouseY >= int(e.Position.Y) && mouseY <= int(e.Position.Y)+int(e.Size.Height)
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
