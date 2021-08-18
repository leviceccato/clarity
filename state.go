package main

import (
	"github.com/leviceccato/clarity/system"
	"github.com/leviceccato/clarity/utility"
	"github.com/leviceccato/clarity/world"

	"github.com/hajimehoshi/ebiten/v2"
)

type stateWorld interface {
	GetName() string
	GetSystems() []world.WorldSystem
	Load()
	Exit()
	Enter()
	Update()
	Draw(*ebiten.Image)
}

type state struct {
	worlds                    map[string]stateWorld
	activeWorlds              []string
	events                    []interface{}
	renderWidth, renderHeight int
	controls                  map[system.Control]*system.InputData
	mouseInputs               map[ebiten.MouseButton]system.Control
	keyInputs                 map[ebiten.Key]system.Control
}

func newState() *state {
	s := &state{}
	s.worlds = map[string]stateWorld{}
	s.mouseInputs = map[ebiten.MouseButton]system.Control{
		ebiten.MouseButtonLeft: system.ControlClick,
	}
	s.keyInputs = map[ebiten.Key]system.Control{
		ebiten.KeySpace:      system.ControlJump,
		ebiten.KeyArrowUp:    system.ControlUp,
		ebiten.KeyW:          system.ControlUp,
		ebiten.KeyArrowLeft:  system.ControlLeft,
		ebiten.KeyA:          system.ControlLeft,
		ebiten.KeyArrowRight: system.ControlRight,
		ebiten.KeyD:          system.ControlRight,
		ebiten.KeyArrowDown:  system.ControlDown,
		ebiten.KeyS:          system.ControlDown,
		ebiten.KeyEscape:     system.ControlMenu,
		ebiten.KeyBackquote:  system.ControlDebug,
	}
	s.UpdateControls()
	return s
}

// Map the controls to the input fields
func (s *state) UpdateControls() {
	// Clear any existing values
	s.controls = map[system.Control]*system.InputData{}
	for _, control := range s.mouseInputs {
		s.controls[control] = nil
	}
	for _, control := range s.keyInputs {
		s.controls[control] = nil
	}
}

func (s *state) loadWorld(w stateWorld) {
	w.Load()
	for _, system := range w.GetSystems() {
		system.Load()
	}
	s.worlds[w.GetName()] = w
}

func (s state) RenderWidth() int {
	return s.renderWidth
}

func (s state) RenderHeight() int {
	return s.renderHeight
}

func (s state) MouseInputs() map[ebiten.MouseButton]system.Control {
	return s.mouseInputs
}

func (s state) KeyInputs() map[ebiten.Key]system.Control {
	return s.keyInputs
}

func (s state) Controls() map[system.Control]*system.InputData {
	return s.controls
}

func (s *state) SetControl(control system.Control, data *system.InputData) {
	s.controls[control] = data
}

func (s state) Events() []interface{} {
	return s.events
}

func (s *state) SetEvents(events []interface{}) {
	s.events = events
}

func (s *state) AddEvent(event interface{}) {
	s.events = append(s.events, event)
}

// Build slices for exiting and entering worlds based on what
// worlds are currently active and those that will be. Then
// exit and enter all of those worlds.
func (s *state) ActivateWorlds(names []string) {
	var w stateWorld
	exitingWorlds := utility.SliceStringDifference(s.activeWorlds, names)
	enteringWorlds := utility.SliceStringDifference(names, s.activeWorlds)
	for _, world := range exitingWorlds {
		w = s.worlds[world]
		w.Exit()
		for _, system := range w.GetSystems() {
			system.Exit()
		}
	}
	s.activeWorlds = names
	for _, world := range enteringWorlds {
		w = s.worlds[world]
		w.Enter()
		for _, system := range w.GetSystems() {
			system.Enter()
		}
	}
}

func (s *state) update() {
	var (
		w      stateWorld
		system world.WorldSystem
	)
	for _, world := range s.activeWorlds {
		w = s.worlds[world]
		w.Update()
		for _, system = range w.GetSystems() {
			system.Update()
		}
	}
}

func (s *state) draw(screen *ebiten.Image) {
	var (
		w      stateWorld
		system world.WorldSystem
	)
	for _, world := range s.activeWorlds {
		w = s.worlds[world]
		w.Draw(screen)
		for _, system = range w.GetSystems() {
			system.Draw(screen)
		}
	}
}
