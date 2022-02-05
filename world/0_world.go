package world

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/game"
)

type world struct {
	name     string
	systems  []game.GameSystem
	entities []*entity.Entity
}

func (w *world) Update(s *game.State) {
	for _, sys := range w.systems {
		sys.Update(s)
	}
}

func (w *world) Draw(s *game.State, screen *ebiten.Image) {
	for _, sys := range w.systems {
		sys.Draw(s, screen)
	}
}

func (w *world) Enter(s *game.State) {
	for _, sys := range w.systems {
		sys.Enter(s)
	}
}

func (w *world) Exit(s *game.State) {
	for _, sys := range w.systems {
		sys.Enter(s)
	}
}

func (w *world) Load(s *game.State) {
	for _, sys := range w.systems {
		sys.Load(s)
	}
}

func (w *world) Name() string {
	return w.name
}

// Add entities to systems based on their components. This is an
// expensive function and should be used sparingly. Ideally after
// multiple system and entity updates.
func (w *world) updateSystems() {
	// Assume entity to be suitable by default
	hasComponents := true

	for _, system := range w.systems {
		for _, entity := range w.entities {
			entityReflection := reflect.Indirect(reflect.ValueOf(entity))
			components := system.GetComponents()

			// A system w/o components should have no entities
			if len(components) == 0 {
				hasComponents = false
			}

			// Check if entity has required components
			for _, component := range components {
				field := entityReflection.FieldByName(component)
				if field.IsNil() {
					hasComponents = false
				}
			}

			// The entity is suitable and we can add it
			if hasComponents {
				system.AddEntity(entity)
			}

			// Reset for next iteration
			hasComponents = true
		}
	}
}
