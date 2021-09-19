package system

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type drawSystem struct {
	system
	state SystemState
}

func NewDrawSystem(state SystemState) *drawSystem {
	s := &drawSystem{}
	s.state = state
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

		// Draw image
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(e.Position.X, e.Position.Y)
		screen.DrawImage(
			e.Appearance.Image.SubImage(*e.Appearance.Frames[e.Appearance.Frame]).(*ebiten.Image),
			options,
		)

		// Draw lines of text
		if e.Text != nil {
			for i, line := range e.Text.Lines {
				text.Draw(
					screen,
					line.Content,
					e.Text.Font,
					int(e.Position.X)+int(line.X)+int(e.Text.Padding),
					int(e.Position.Y)+(i*e.Text.LineHeight),
					e.Text.Color,
				)
			}
		}
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"TPS: %0.2f\nFPS: %0.2f",
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
	))
}

func (s *drawSystem) Enter() {}

func (s *drawSystem) Exit() {}
