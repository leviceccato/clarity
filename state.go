package main

import (
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

type control struct {
}
type state struct {
	worlds        map[string]stateWorld
	activeWorlds  []string
	renderWidth   int
	renderHeight  int
	controls      map[string]*control
	mouseInputMap map[ebiten.MouseButton]string
	keyInputMap   map[ebiten.Key]string
}

func newState() *state {
	s := &state{}
	s.worlds = map[string]stateWorld{}
	s.mouseInputMap = map[ebiten.MouseButton]string{
		ebiten.MouseButtonLeft: "jump",
	}
	s.keyInputMap = map[ebiten.Key]string{
		ebiten.KeyArrowUp:    "up",
		ebiten.KeyW:          "up",
		ebiten.KeyArrowLeft:  "left",
		ebiten.KeyA:          "left",
		ebiten.KeyArrowRight: "right",
		ebiten.KeyD:          "right",
		ebiten.KeyArrowDown:  "down",
		ebiten.KeyS:          "down",
		ebiten.KeyEscape:     "menu",
		ebiten.KeyBackquote:  "debug",
	}
	s.UpdateControls()
	return s
}

// Map the controls to the input fields
func (s *state) UpdateControls() {
	// Clear any existing values
	s.controls = map[string]*control{}
	for _, control := range s.mouseInputMap {
		s.controls[control] = nil
	}
	for _, control := range s.keyInputMap {
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

// Expose as function so it can be used in an interface
func (s state) RenderWidth() int {
	return s.renderWidth
}

// Expose as function so it can be used in an interface
func (s state) RenderHeight() int {
	return s.renderHeight
}

// Build slices for exiting and entering worlds based on what
// worlds are currently active and those that will be. Then
// exit and enter all of those worlds.
func (s *state) activateWorlds(names []string) {
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
	var w stateWorld
	for _, world := range s.activeWorlds {
		w = s.worlds[world]
		w.Update()
		for _, system := range w.GetSystems() {
			system.Update()
		}
	}
}

func (s *state) draw(screen *ebiten.Image) {
	var w stateWorld
	for _, world := range s.activeWorlds {
		w = s.worlds[world]
		w.Draw(screen)
		for _, system := range w.GetSystems() {
			system.Draw(screen)
		}
	}
}
