package system

import (
	"math"

	"github.com/leviceccato/clarity/game"

	"github.com/hajimehoshi/ebiten/v2"
)

// 1/60th of a second in milliseconds. Ebiten ensures that the game is always run at 60 FPS.
var delta = (1.0 / 60) * 1000

type animationSystem struct {
	System
}

func NewAnimationSystem(s *game.State) *animationSystem {
	sys := &animationSystem{}
	sys.Components = []string{"Appearance"}
	return sys
}

func (sys *animationSystem) Load(s *game.State) {}

func (sys *animationSystem) Update(s *game.State) {
	for _, e := range sys.Entities {
		// Skip animation if there is only 1 frame
		if len(e.Appearance.Frames) == 1 {
			continue
		}

		e.Appearance.Time += delta
		duration := float64(e.Appearance.Duration)
		if e.Appearance.Time >= duration {
			e.Appearance.Time = math.Min(duration, e.Appearance.Time-duration)
		}
		sequence := e.Appearance.Sequences[e.Appearance.Sequence]
		difference := float64(sequence.To - sequence.From)
		e.Appearance.Frame = int(math.Floor(e.Appearance.Time/duration*difference)) + sequence.From
	}
}

func (sys *animationSystem) Draw(s *game.State, screen *ebiten.Image) {}

func (sys *animationSystem) Enter(s *game.State) {}

func (sys *animationSystem) Exit(s *game.State) {}
