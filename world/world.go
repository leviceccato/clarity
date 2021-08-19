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
	hasComponents := true
	var (
		entityReflection, field reflect.Value
		system                  WorldSystem
		entity                  *entity.Entity
	)
	for _, system = range w.systems {
		for _, entity = range w.entities {
			entityReflection = reflect.Indirect(reflect.ValueOf(entity))
			for _, component := range system.Components() {
				field = entityReflection.FieldByName(component)
				if field.IsNil() {
					hasComponents = false
				}
			}
			if hasComponents {
				system.AddEntity(entity)
			}
			hasComponents = true
		}
	}
}
