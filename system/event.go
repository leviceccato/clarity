package system

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type event struct {
	system
	state SystemState
}

type eventQuit struct {
	code int
}

type eventActivateWorlds struct {
	names []string
}

func NewEventSystem(state SystemState) *event {
	s := &event{}
	return s
}

func (s *event) Load() {}

func (s *event) Update() {
	events := s.state.Events()
	for len(events) > 0 {
		switch e := events[0].(type) {
		case eventQuit:
			os.Exit(e.code)
		case eventActivateWorlds:
			s.state.ActivateWorlds(e.names)
		}
	}
}

func (s *event) Draw(screen *ebiten.Image) {}

func (s *event) Enter() {}

func (s *event) Exit() {}
