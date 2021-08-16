package system

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type drawSystem struct {
	system
}

func NewDrawSystem() *drawSystem {
	s := &drawSystem{}
	s.components = []string{
		"Appearance",
		"Position",
		"Size",
	}
	return s
}

func (s *drawSystem) Load() {}

func (s *drawSystem) Update() {}

func (s *drawSystem) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
	for _, e := range s.entities {
		options := &ebiten.DrawImageOptions{}
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

func (s *drawSystem) Enter() {}

func (s *drawSystem) Exit() {}
