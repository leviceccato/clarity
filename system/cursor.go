package system

import (
	"github.com/leviceccato/clarity/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type cursorSystem struct {
	System
}

func NewCursorSystem(s *game.State) *cursorSystem {
	sys := &cursorSystem{}
	sys.Components = []string{
		"Cursor",
		"Position",
		"Appearance",
	}
	return sys
}

func (sys *cursorSystem) Load(s *game.State) {}

func (sys *cursorSystem) Update(s *game.State) {
	isCursorHovering := s.IsCursorHovering
	for _, e := range sys.Entities {
		if isCursorHovering && e.Appearance.Sequence == "default" {
			e.Appearance.Sequence = "pointer"
		}
		if !isCursorHovering && e.Appearance.Sequence == "pointer" {
			e.Appearance.Sequence = "default"
		}
		x, y := ebiten.CursorPosition()
		e.Position.X = float64(x)
		e.Position.Y = float64(y)
	}
}

func (sys *cursorSystem) Draw(s *game.State, creen *ebiten.Image) {}

func (sys *cursorSystem) Enter(s *game.State) {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
}

func (sys *cursorSystem) Exit(s *game.State) {
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
}
