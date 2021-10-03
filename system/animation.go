package system

import (
	"math"

	"github.com/leviceccato/clarity/component"
	"github.com/leviceccato/clarity/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

// 1/60th of a second in milliseconds. Ebiten ensures that the game is always run at 60 FPS.
var delta = (1.0 / 60) * 1000

type animationSystem struct {
	system
}

func NewAnimationSystem() *animationSystem {
	s := &animationSystem{}
	s.components = []string{"Appearance"}
	return s
}

func (s *animationSystem) Load() {}

func (s *animationSystem) Update() {
	var (
		duration, difference float64
		sequence             *component.AppearanceSequence
		e                    *entity.Entity
	)
	for _, e = range s.entities {
		// Skip animation if there is only 1 frame
		if len(e.Appearance.Frames) == 1 {
			continue
		}
		e.Appearance.Time += delta
		duration = float64(e.Appearance.Duration)
		if e.Appearance.Time >= duration {
			e.Appearance.Time = math.Min(duration, e.Appearance.Time-duration)
		}
		sequence = e.Appearance.Sequences[e.Appearance.Sequence]
		difference = float64(sequence.To - sequence.From)
		e.Appearance.Frame = int(math.Floor(e.Appearance.Time/duration*difference)) + sequence.From
	}
}

func (s *animationSystem) Draw(screen *ebiten.Image) {}

func (s *animationSystem) Enter() {}

func (s *animationSystem) Exit() {}
