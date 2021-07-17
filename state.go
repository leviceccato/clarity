package main

import (
	"github.com/leviceccato/clarity/utility"
	"github.com/leviceccato/clarity/world"
)

type stateWorld interface {
	GetName() string
	GetSystems() []world.WorldSystem
	Load()
	Exit()
	Enter()
	Update()
	Draw()
}

type state struct {
	worlds       map[string]stateWorld
	activeWorlds []string
}

func newState() *state {
	return &state{
		worlds: map[string]stateWorld{},
	}
}

func (s *state) loadWorld(w stateWorld) {
	w.Load()
	for _, system := range w.GetSystems() {
		system.Load()
	}
	s.worlds[w.GetName()] = w
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

func (s *state) draw() {
	var w stateWorld
	for _, world := range s.activeWorlds {
		w = s.worlds[world]
		w.Draw()
		for _, system := range w.GetSystems() {
			system.Draw()
		}
	}
}
