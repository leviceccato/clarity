package system

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type eventSystem struct {
	system
	state SystemState
}

type quitEvent struct {
	code int
}

type activateWorldsEvent struct {
	names []string
}

func NewEventSystem(state SystemState) *eventSystem {
	s := &eventSystem{}
	return s
}

func (s *eventSystem) Load() {}

func (s *eventSystem) Update() {
	events := s.state.Events()
	for len(events) > 0 {
		switch e := events[0].(type) {
		case quitEvent:
			os.Exit(e.code)
		case activateWorldsEvent:
			s.state.ActivateWorlds(e.names)
		}
	}
}

func (s *eventSystem) Draw(screen *ebiten.Image) {}

func (s *eventSystem) Enter() {}

func (s *eventSystem) Exit() {}
