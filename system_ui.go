package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

type uiSystem struct {
	system
}

func newUISystem() runner {
	s := &uiSystem{}
	s.components = []string{"position", "appearance", "ui"}
	return s
}

func (s *uiSystem) run() error {
	r.DrawText("Controls:", 20, 20, 10, r.Black)
	for _, e := range s.entities {
		r.DrawTexture(
			e.appearance.texture,
			int32(e.position.value.X),
			int32(e.position.value.Y),
			r.White,
		)
	}
	return nil
}

func (s *uiSystem) getComponents() []string {
	return s.components
}

func (s *uiSystem) sortEntities() {
}

func (s *uiSystem) getEntities() []*entity {
	return s.entities
}

func (s *uiSystem) addEntity(e *entity) {
	s.entities = append(s.entities, e)
}
