package system

import (
	"github.com/leviceccato/clarity/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type cursorSystem struct {
	system
}

func NewCursorSystem(state SystemState) *cursorSystem {
	s := &cursorSystem{}
	s.state = state
	s.components = []string{"Cursor", "Position", "Appearance"}
	return s
}

func (s *cursorSystem) Load() {}

func (s *cursorSystem) Update() {
	var (
		e    *entity.Entity
		x, y int
	)
	isCursorHovering := s.state.IsCursorHovering()
	for _, e = range s.entities {
		if isCursorHovering && e.Appearance.Sequence == "default" {
			e.Appearance.Sequence = "pointer"
		}
		if !isCursorHovering && e.Appearance.Sequence == "pointer" {
			e.Appearance.Sequence = "default"
		}
		x, y = ebiten.CursorPosition()
		e.Position.X = float64(x)
		e.Position.Y = float64(y)
	}
}

func (s *cursorSystem) Draw(screen *ebiten.Image) {}

func (s *cursorSystem) Enter() {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
}

func (s *cursorSystem) Exit() {
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
}
