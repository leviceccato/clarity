package system

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type hoverSystem struct {
	system
}

func NewHoverSystem() *hoverSystem {
	s := &hoverSystem{}
	s.components = []string{
		"Hover",
		"Position",
		"Size",
	}
	return s
}

func (s *hoverSystem) Load() {}

func (s *hoverSystem) Update() {}

func (s *hoverSystem) Draw(screen *ebiten.Image) {}

func (s *hoverSystem) Enter() {}

func (s *hoverSystem) Exit() {}
