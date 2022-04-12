package game

import (
	"github.com/leviceccato/clarity/engine"
)

type quitEvent struct {
	code int
}

type activateWorldsEvent struct {
	names []string
}

func newEventSystem(_ *Game) *engine.System {
	s := engine.NewSystem("event", []string{})

	s.Update = func() error {
		// for _, entityId := range s.EntityIds {
		// 	// // No events, nothing to do
		// 	// if len(s.Events) == 0 {
		// 	// 	return
		// 	// }

		// 	// for _, event := range s.Events {
		// 	// 	switch e := event.(type) {
		// 	// 	case quitEvent:
		// 	// 		os.Exit(e.code)
		// 	// 	case activateWorldsEvent:
		// 	// 		s.ActivateWorlds(e.names...)
		// 	// 	}
		// 	// }

		// 	// // Clear events
		// 	// s.ClearEvents()
		// }

		return nil
	}

	return s
}
