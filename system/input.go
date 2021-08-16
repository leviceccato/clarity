package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type input struct {
	system
	state SystemState
}

type InputData struct {
	X, Y int
}

type Control int

const (
	ControlJump Control = iota + 1
	ControlUp
	ControlLeft
	ControlRight
	ControlDown
	ControlMenu
	ControlDebug
	ControlClick
)

func NewInputSystem(state SystemState) *input {
	s := &input{}
	s.state = state
	return s
}

func (s *input) Load() {}

func (s *input) Update() {
	for mouseInput, control := range s.state.MouseInputs() {
		if inpututil.IsMouseButtonJustPressed(mouseInput) {
			x, y := ebiten.CursorPosition()
			s.state.SetControl(control, &InputData{X: x, Y: y})
		}
		if inpututil.IsMouseButtonJustReleased(mouseInput) {
			s.state.SetControl(control, nil)
		}
	}
	for keyInput, control := range s.state.KeyInputs() {
		if inpututil.IsKeyJustPressed(keyInput) {
			s.state.SetControl(control, &InputData{})
		}
		if inpututil.IsKeyJustReleased(keyInput) {
			s.state.SetControl(control, nil)
		}
	}
}

func (s *input) Draw(screen *ebiten.Image) {}

func (s *input) Enter() {}

func (s *input) Exit() {}
