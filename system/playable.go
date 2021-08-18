package system

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type playableSystem struct {
	system
	state SystemState
}

func NewPlayableSystem(state SystemState) *playableSystem {
	s := &playableSystem{}
	s.state = state
	return s
}

func (s *playableSystem) Load() {}

func (s *playableSystem) Update() {
	c := s.state.Controls()
	for _, e := range s.entities {
		if c[ControlUp] != nil {
			e.Position.Y -= 3
		}
		if c[ControlLeft] != nil {
			e.Position.X -= 3
		}
		if c[ControlRight] != nil {
			e.Position.X += 3
		}
		if c[ControlDown] != nil {
			e.Position.Y += 3
		}
		if c[ControlMenu] != nil {
			s.state.AddEvent(quitEvent{})
		}
	}
}

func (s *playableSystem) Draw(screen *ebiten.Image) {}

func (s *playableSystem) Enter() {}

func (s *playableSystem) Exit() {}
