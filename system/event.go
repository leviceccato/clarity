package system

import (
	"os"

	"github.com/leviceccato/clarity/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type eventSystem struct {
	System
}

type quitEvent struct {
	code int
}

type activateWorldsEvent struct {
	names []string
}

func NewEventSystem(s *game.State) *eventSystem {
	sys := &eventSystem{}
	return sys
}

func (sys *eventSystem) Load(s *game.State) {}

func (sys *eventSystem) Update(s *game.State) {
	// No events, nothing to do
	if len(s.Events) == 0 {
		return
	}

	for _, event := range s.Events {
		switch e := event.(type) {
		case quitEvent:
			os.Exit(e.code)
		case activateWorldsEvent:
			s.ActivateWorlds(e.names...)
		}
	}

	// Clear events
	s.ClearEvents()
}

func (sys *eventSystem) Draw(s *game.State, screen *ebiten.Image) {}

func (sys *eventSystem) Enter(s *game.State) {}

func (sys *eventSystem) Exit(s *game.State) {}
