package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

type drawSystem struct {
	system
}

func newDrawSystem() runner {
	s := &drawSystem{}
	s.components = []string{"position", "appearance"}
	return s
}

func (s drawSystem) run() error {
	r.BeginDrawing()
	r.ClearBackground(r.Magenta)
	for _, e := range s.entities {
		r.DrawTexture(
			e.appearance.texture,
			int32(e.position.x),
			int32(e.position.y),
			r.White,
		)
	}
	r.EndDrawing()
	return nil
}

func (s drawSystem) getComponents() []string {
	return s.components
}

func (s *drawSystem) sortEntities() {
	sort.Slice(s.entities, func(a, b int) bool {
		aZIndex := s.entities[a].zIndex
		bZIndex := s.entities[b].zIndex
		if aZIndex == nil || bZIndex == nil {
			return false
		}
		return aZIndex.value > bZIndex.value
	})
}

func (s *drawSystem) getEntities() []*entity {
	return s.entities
}

func (s *drawSystem) addEntity(e *entity) {
	s.entities = append(s.entities, e)
}
