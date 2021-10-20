package main

import (
	"github.com/leviceccato/clarity/system"
	"github.com/leviceccato/clarity/util"
	"github.com/leviceccato/clarity/world"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type stateWorld interface {
	Name() string
	Systems() []world.WorldSystem
	Load()
	Exit()
	Enter()
	Update()
	Draw(*ebiten.Image)
}

type state struct {
	worlds       map[string]stateWorld
	activeWorlds []string
	events       []interface{}

	controls    map[system.Control]*system.InputData
	mouseInputs map[ebiten.MouseButton]system.Control
	keyInputs   map[ebiten.Key]system.Control

	fonts map[string]*font.Face

	isCursorHovering bool
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
	for _, system := range w.Systems() {
		system.Load()
	}
	s.worlds[w.Name()] = w
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

func (s state) Font(name string) *font.Face {
	return s.fonts[name]
}

func (s *state) SetIsCursorHovering(isCursorHovering bool) {
	s.isCursorHovering = isCursorHovering
}

func (s state) IsCursorHovering() bool {
	return s.isCursorHovering
}

// Build slices for exiting and entering worlds based on what
// worlds are currently active and those that will be. Then
// exit and enter all of those worlds.
func (s *state) ActivateWorlds(names []string) {
	exitingWorlds := util.SliceStringDifference(s.activeWorlds, names)
	enteringWorlds := util.SliceStringDifference(names, s.activeWorlds)
	var (
		w         stateWorld
		system    world.WorldSystem
		worldName string
	)
	for _, worldName = range exitingWorlds {
		w = s.worlds[worldName]
		w.Exit()
		for _, system = range w.Systems() {
			system.Exit()
		}
	}
	s.activeWorlds = names
	for _, worldName = range enteringWorlds {
		w = s.worlds[worldName]
		w.Enter()
		for _, system = range w.Systems() {
			system.Enter()
		}
	}
}

func (s *state) update() {
	var (
		w         stateWorld
		system    world.WorldSystem
		worldName string
	)
	for _, worldName = range s.activeWorlds {
		w = s.worlds[worldName]
		w.Update()
		for _, system = range w.Systems() {
			system.Update()
		}
	}
}

func (s *state) draw(screen *ebiten.Image) {
	var (
		w         stateWorld
		system    world.WorldSystem
		worldName string
	)
	for _, worldName = range s.activeWorlds {
		w = s.worlds[worldName]
		w.Draw(screen)
		for _, system = range w.Systems() {
			system.Draw(screen)
		}
	}
}
