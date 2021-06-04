package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

type interfaceSystem struct {
	system
}

func newInterfaceSystem() runner {
	s := &interfaceSystem{}
	s.components = []string{"position", "appearance", "interface"}
	return s
}

func (s *interfaceSystem) run() error {
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

func (s *interfaceSystem) getComponents() []string {
	return s.components
}

func (s *interfaceSystem) sortEntities() {
}

func (s *interfaceSystem) getEntities() []*entity {
	return s.entities
}

func (s *interfaceSystem) addEntity(e *entity) {
	s.entities = append(s.entities, e)
}
