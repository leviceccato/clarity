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
	// var (
	// 	e    *entity.Entity
	// 	x, y int
	// )
	// for _, e = range s.entities {
	// 	x, y = ebiten.CursorPosition()
	// 	e.Position.X = float64(x)
	// 	e.Position.Y = float64(y)
	// }
}

func (s *cursorSystem) Draw(screen *ebiten.Image) {}

func (s *cursorSystem) Enter() {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
}

func (s *cursorSystem) Exit() {
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
}
