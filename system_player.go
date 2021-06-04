package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

type playerSystem struct {
	system
}

func newPlayerSytem() runner {
	s := &playerSystem{}
	s.components = []string{"position", "controls"}
	return s
}

func (s *playerSystem) run() error {
	for _, e := range s.entities {
		if r.IsKeyDown(r.KeyW) {
			e.position.value.Y -= 2
		}
		if r.IsKeyDown(r.KeyS) {
			e.position.value.Y += 2
		}
		if r.IsKeyDown(r.KeyA) {
			e.position.value.X -= 2
		}
		if r.IsKeyDown(r.KeyD) {
			e.position.value.X += 2
		}
	}
	return nil
}

func (s *playerSystem) getComponents() []string {
	return s.components
}

func (s *playerSystem) sortEntities() {}

func (s *playerSystem) getEntities() []*entity {
	return s.entities
}

func (s *playerSystem) addEntity(e *entity) {
	s.entities = append(s.entities, e)
}
