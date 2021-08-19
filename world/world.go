package world

import (
	"reflect"

	"github.com/leviceccato/clarity/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type WorldSystem interface {
	Components() []string
	EntityCount() int
	AddEntity(*entity.Entity)
	Enter()
	Exit()
	Load()
	Update()
	Draw(*ebiten.Image)
}

type gameState interface {
	RenderWidth() int
	RenderHeight() int
}

type world struct {
	name     string
	systems  []WorldSystem
	entities []*entity.Entity
}

func (w *world) Name() string {
	return w.name
}

func (w *world) Systems() []WorldSystem {
	return w.systems
}

// Add entities to systems based on their components. This is an
// expensive function and should be used sparingly. Ideally after
// multiple system and entity updates.
func (w *world) updateSystems() {
	var (
		entityReflection, field reflect.Value
		system                  WorldSystem
		entity                  *entity.Entity
		components              []string
	)
	// Assume entity to be suitable by default
	hasComponents := true
	for _, system = range w.systems {
		for _, entity = range w.entities {
			entityReflection = reflect.Indirect(reflect.ValueOf(entity))
			components = system.Components()

			// A system w/o components should have no entities
			if len(components) == 0 {
				hasComponents = false
			}

			// Check if entity has required components
			for _, component := range components {
				field = entityReflection.FieldByName(component)
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
