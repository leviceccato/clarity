package system

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type animation struct {
	system
}

func NewAnimationSystem() *animation {
	s := &animation{}
	s.components = []string{"Appearance"}
	return s
}

func (s *animation) Load() {}

func (s *animation) Update() {
	for _, e := range s.entities {
		e.Appearance.Time += 16
		duration := float64(e.Appearance.Duration)
		if e.Appearance.Time >= duration {
			e.Appearance.Time = math.Min(duration, e.Appearance.Time-duration)
		}
		sequence := e.Appearance.Sequences[e.Appearance.Sequence]
		length := float64(sequence.To - sequence.From)
		e.Appearance.Frame = int(math.Floor(e.Appearance.Time/duration*length)) + sequence.From
	}
}

func (s *animation) Draw(screen *ebiten.Image) {}

func (s *animation) Enter() {}

func (s *animation) Exit() {}
