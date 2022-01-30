package system

import (
	"github.com/leviceccato/clarity/component"
	"github.com/leviceccato/clarity/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type playableSystem struct {
	System
}

func NewPlayableSystem(s *game.State) *playableSystem {
	sys := &playableSystem{}
	sys.Components = []string{
		"Playable",
	}
	return sys
}

func (sys *playableSystem) Load(s *game.State) {}

func (sys *playableSystem) Update(s *game.State) {
	c := s.Controls
	for _, e := range sys.Entities {
		if c[component.ControlUp] != nil {
			e.Position.Y -= 3
		}
		if c[component.ControlLeft] != nil {
			e.Position.X -= 3
		}
		if c[component.ControlRight] != nil {
			e.Position.X += 3
		}
		if c[component.ControlDown] != nil {
			e.Position.Y += 3
		}
		if c[component.ControlMenu] != nil {
			s.AddEvent(activateWorldsEvent{
				names: []string{"title"},
			})
		}
	}
}

func (sys *playableSystem) Draw(s *game.State, screen *ebiten.Image) {}

func (sys *playableSystem) Enter(s *game.State) {}

func (sys *playableSystem) Exit(s *game.State) {}
