package system

import (
	"github.com/leviceccato/clarity/component"
	"github.com/leviceccato/clarity/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type inputSystem struct {
	System
}

func NewInputSystem(s *game.State) *inputSystem {
	sys := &inputSystem{}
	return sys
}

func (sys *inputSystem) Load(s *game.State) {}

func (sys *inputSystem) Update(s *game.State) {
	for mouseInput, control := range s.MouseInputs {
		if inpututil.IsMouseButtonJustPressed(mouseInput) {
			x, y := ebiten.CursorPosition()
			s.SetControl(control, &component.InputData{X: x, Y: y})
		}
		if inpututil.IsMouseButtonJustReleased(mouseInput) {
			s.SetControl(control, nil)
		}
	}

	for keyInput, control := range s.KeyInputs {
		if inpututil.IsKeyJustPressed(keyInput) {
			s.SetControl(control, &component.InputData{})
		}
		if inpututil.IsKeyJustReleased(keyInput) {
			s.SetControl(control, nil)
		}
	}
}

func (sys *inputSystem) Draw(s *game.State, screen *ebiten.Image) {}

func (sys *inputSystem) Enter(s *game.State) {}

func (sys *inputSystem) Exit(s *game.State) {}
