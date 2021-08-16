package system

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type event struct {
	system
	state SystemState
}

type quitEvent struct {
	code int
}

type activateWorldsEvent struct {
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
		case quitEvent:
			os.Exit(e.code)
		case activateWorldsEvent:
			s.state.ActivateWorlds(e.names)
		}
	}
}

func (s *event) Draw(screen *ebiten.Image) {}

func (s *event) Enter() {}

func (s *event) Exit() {}
