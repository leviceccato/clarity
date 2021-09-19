package system

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type cursorSystem struct {
	system
}

func NewCursorSystem() *cursorSystem {
	s := &cursorSystem{}
	s.components = []string{"Cursor", "Position"}
	return s
}

func (s *cursorSystem) Load() {}

func (s *cursorSystem) Update() {
	for _, e := range s.entities {
		x, y := ebiten.CursorPosition()
		e.Position.X = float64(x)
		e.Position.Y = float64(y)
	}
}

func (s *cursorSystem) Draw(screen *ebiten.Image) {}

func (s *cursorSystem) Enter() {}

func (s *cursorSystem) Exit() {}
