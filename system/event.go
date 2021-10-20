package system

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type eventSystem struct {
	system
}

type quitEvent struct {
	code int
}

type activateWorldsEvent struct {
	names []string
}

func NewEventSystem(state SystemState) *eventSystem {
	s := &eventSystem{}
	s.state = state
	return s
}

func (s *eventSystem) Load() {}

func (s *eventSystem) Update() {
	events := s.state.Events()
	if len(events) == 0 {
		return
	}
	var event interface{}
	for _, event = range events {
		switch e := event.(type) {
		case quitEvent:
			os.Exit(e.code)
		case activateWorldsEvent:
			s.state.ActivateWorlds(e.names)
		}
	}
	// Clear events
	s.state.SetEvents([]interface{}{})
}

func (s *eventSystem) Draw(screen *ebiten.Image) {}

func (s *eventSystem) Enter() {}

func (s *eventSystem) Exit() {}
