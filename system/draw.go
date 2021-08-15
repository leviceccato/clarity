package system

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type draw struct {
	system
}

func NewDrawSystem() *draw {
	s := &draw{}
	s.components = []string{
		"Appearance",
		"Position",
		"Size",
	}
	return s
}

func (s *draw) Load() {}

func (s *draw) Update() {
	for _, e := range s.entities {
		e.Position.X += (rand.Float64() * 5) - 2.5
		e.Position.Y += (rand.Float64() * 5) - 2.5
	}
}

func (s *draw) Draw(screen *ebiten.Image) {
	for _, e := range s.entities {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(e.Position.X, e.Position.Y)
		screen.DrawImage(e.Appearance.Image, options)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"TPS: %0.2f\nFPS: %0.2f",
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
	))
}

func (s *draw) Enter() {}

func (s *draw) Exit() {}
