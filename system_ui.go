package main

import (
	"fmt"

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
	r.DrawRectangle(10, 10, 100, 50, r.Color{R: 0, G: 0, B: 0, A: 50})
	r.DrawText(fmt.Sprintf("FPS: %d", int32(r.GetFPS())), 20, 20, 10, r.White)
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
