package game

import (
	"image/color"

	"github.com/leviceccato/clarity/component"
	"github.com/leviceccato/clarity/entity"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type GameSystem interface {
	GetComponents() []string
	AddEntity(*entity.Entity)
	Update(*State)
	Draw(*State, *ebiten.Image)
	Load(*State)
	Exit(*State)
	Enter(*State)
}

type gameWorld interface {
	Name() string
	Systems() []GameSystem
	Load(*State)
	Exit(*State)
	Enter(*State)
	Update(*State)
	Draw(*State, *ebiten.Image)
}

type State struct {
	Options

	worlds       map[string]gameWorld
	activeWorlds []string
	Events       []interface{}

	Controls         map[component.Control]*component.InputData
	MouseInputs      map[ebiten.MouseButton]component.Control
	KeyInputs        map[ebiten.Key]component.Control
	CursorX, CursorY float64

	IsCursorHovering bool

	Fonts map[string]font.Face

	Colors map[string]color.NRGBA
}

func (s *State) Update() error {
	s.UpdateCursor()

	for _, worldName := range s.activeWorlds {
		w := s.worlds[worldName]
		w.Update(s)
	}
	return nil
}

func (s *State) Draw(screen *ebiten.Image) {
	for _, worldName := range s.activeWorlds {
		w := s.worlds[worldName]
		w.Draw(s, screen)
	}
}

func (s *State) Layout(_, _ int) (int, int) {
	return s.RenderWidth, s.RenderHeight
}

// Map the controls to the input fields
func (s *State) UpdateControls() {
	// Clear any existing values
	s.Controls = map[component.Control]*component.InputData{}

	for _, control := range s.MouseInputs {
		s.Controls[control] = nil
	}
	for _, control := range s.KeyInputs {
		s.Controls[control] = nil
	}
}

func (s *State) LoadWorld(w gameWorld) {
	w.Load(s)
	s.worlds[w.Name()] = w
}

// Build slices for exiting and entering worlds based on what
// worlds are currently active and those that will be. Then
// exit and enter all of those worlds.
func (s *State) ActivateWorlds(names ...string) {
	exitingWorlds := sliceStringDifference(s.activeWorlds, names)
	enteringWorlds := sliceStringDifference(names, s.activeWorlds)

	for _, worldName := range exitingWorlds {
		w := s.worlds[worldName]
		w.Exit(s)
	}

	s.activeWorlds = names

	for _, worldName := range enteringWorlds {
		w := s.worlds[worldName]
		w.Enter(s)
	}
}

func (s *State) UpdateCursor() {
	x, y := ebiten.CursorPosition()
	s.CursorX = float64(x)
	s.CursorY = float64(y)
}

func (s *State) SetControl(control component.Control, data *component.InputData) {
	s.Controls[control] = data
}

func (s *State) AddEvent(event interface{}) {
	s.Events = append(s.Events, event)
}

func (s *State) ClearEvents() {
	s.Events = []interface{}{}
}

func (s *State) SetIsCursorHovering(isHovering bool) {
	s.IsCursorHovering = isHovering
}

// Find the unique strings in the first slice between 2 slices
// of strings
func sliceStringDifference(a, b []string) []string {
	uniqueStrings := []string{}
	isStringUnique := true

	for _, stringA := range a {
		for _, stringB := range b {
			if stringA == stringB {
				isStringUnique = false
			}
		}

		if isStringUnique {
			uniqueStrings = append(uniqueStrings, stringA)
		}

		isStringUnique = true
	}

	return uniqueStrings
}
