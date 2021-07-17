package world

import (
	"reflect"

	"github.com/leviceccato/clarity/entity"
)

type WorldSystem interface {
	GetComponents() []string
	AddEntity(*entity.Entity)
	Enter()
	Exit()
	Load()
	Update()
	Draw()
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