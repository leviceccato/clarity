package game

import (
	"math"

	"github.com/leviceccato/clarity/engine"
)

func newAnimationSystem(g *Game) *engine.System {
	s := engine.NewSystem("animation", []string{
		"appearance",
	})

	s.Update = func() error {
		for _, entityId := range s.EntityIds {
			e := g.GetEntity(entityId)

			appearance, _ := engine.GetComponent(e, &appearanceComponent{})

			// Skip animation if there is only 1 frame
			if len(appearance.Frames) == 1 {
				continue
			}

			appearance.Time += engine.Delta
			duration := float64(appearance.Duration)

			if appearance.Time >= duration {
				appearance.Time = math.Min(duration, appearance.Time-duration)
			}

			sequence := appearance.Sequences[appearance.Sequence]
			difference := float64(sequence.To - sequence.From)

			appearance.Frame = int(math.Floor(appearance.Time/duration*difference)) + sequence.From
		}

		return nil
	}

	return s
}
