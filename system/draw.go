package system

import (
	"github.com/hajimehoshi/ebiten/v2"
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

func (s *draw) Load() {

}

func (s *draw) Update() {
}

func (s *draw) Draw(screen *ebiten.Image) {
	for _, e := range s.entities {

		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(e.Position.X, e.Position.Y)
		screen.DrawImage(e.Appearance.Image, options)
	}
}

func (s *draw) Enter() {

}

func (s *draw) Exit() {

}
