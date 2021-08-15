package world

import (
	"reflect"

	"github.com/leviceccato/clarity/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type WorldSystem interface {
	GetComponents() []string
	GetEntityCount() int
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

func (w *world) AddEntity(entity *entity.Entity) {
	w.entities = append(w.entities, entity)
}

func (w *world) AddSystem(system WorldSystem) {
	w.systems = append(w.systems, system)
}

func (w *world) GetName() string {
	return w.name
}

func (w *world) GetSystems() []WorldSystem {
	return w.systems
}

// Add entities to systems based on their components. This is an
// expensive function and should be used sparingly. Ideally after
// multiple system and entity updates.
func (w *world) updateSystems() {
	hasComponents := true
	var entityReflection reflect.Value
	var field reflect.Value
	for _, system := range w.systems {
		for _, entity := range w.entities {
			entityReflection = reflect.Indirect(reflect.ValueOf(entity))
			for _, component := range system.GetComponents() {
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
