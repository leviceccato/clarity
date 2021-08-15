package system

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type input struct {
	system
}

func NewInputSystem() *input {
	return &input{}
}

func (s *input) Load() {}

func (s *input) Update() {

}

func (s *input) Draw(screen *ebiten.Image) {}

func (s *input) Enter() {}

func (s *input) Exit() {}
