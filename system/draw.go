package system

import (
	"fmt"

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

func (s *draw) Update() {}

func (s *draw) Draw(screen *ebiten.Image) {
	var options *ebiten.DrawImageOptions
	for _, e := range s.entities {
		options = &ebiten.DrawImageOptions{}
		options.GeoM.Translate(e.Position.X, e.Position.Y)
		screen.DrawImage(
			e.Appearance.Image.SubImage(*e.Appearance.Frames[e.Appearance.Frame]).(*ebiten.Image),
			options,
		)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"TPS: %0.2f\nFPS: %0.2f",
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
	))
}

func (s *draw) Enter() {}

func (s *draw) Exit() {}
