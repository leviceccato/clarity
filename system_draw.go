package main

import (
	"sort"

	r "github.com/gen2brain/raylib-go/raylib"
)

type drawSystem struct {
	system
	camera       *r.Camera2D
	cameraTarget *entity
}

func newDrawSystem() runner {
	s := &drawSystem{}
	s.components = []string{"position", "appearance", "camera"}
	s.camera = &r.Camera2D{
		Offset: r.Vector2{
			X: float32(r.GetScreenWidth() / 2),
			Y: float32(r.GetScreenHeight() / 2),
		},
		Rotation: 0,
		Zoom:     1,
	}
	return s
}

func (s *drawSystem) run() error {
	if s.cameraTarget != nil {
		s.camera.Target = s.cameraTarget.position.value
	}

	r.ClearBackground(r.Blue)
	r.BeginMode2D(*s.camera)
	r.DrawRectangle(-6000, 320, 13000, 8000, r.Green)
	for _, e := range s.entities {
		r.DrawTexture(
			e.appearance.texture,
			int32(e.position.value.X),
			int32(e.position.value.Y),
			r.White,
		)
	}
	r.EndMode2D()
	return nil
}

func (s *drawSystem) getComponents() []string {
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
