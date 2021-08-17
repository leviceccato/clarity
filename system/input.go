package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type inputSystem struct {
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

func NewInputSystem(state SystemState) *inputSystem {
	s := &inputSystem{}
	s.state = state
	return s
}

func (s *inputSystem) Load() {}

func (s *inputSystem) Update() {
	var (
		control    Control
		mouseInput ebiten.MouseButton
		keyInput   ebiten.Key
		x, y       int
	)
	for mouseInput, control = range s.state.MouseInputs() {
		if inpututil.IsMouseButtonJustPressed(mouseInput) {
			x, y = ebiten.CursorPosition()
			s.state.SetControl(control, &InputData{X: x, Y: y})
		}
		if inpututil.IsMouseButtonJustReleased(mouseInput) {
			s.state.SetControl(control, nil)
		}
	}
	for keyInput, control = range s.state.KeyInputs() {
		if inpututil.IsKeyJustPressed(keyInput) {
			s.state.SetControl(control, &InputData{})
		}
		if inpututil.IsKeyJustReleased(keyInput) {
			s.state.SetControl(control, nil)
		}
	}
}

func (s *inputSystem) Draw(screen *ebiten.Image) {}

func (s *inputSystem) Enter() {}

func (s *inputSystem) Exit() {}
