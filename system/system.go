package system

import (
	"github.com/leviceccato/clarity/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type system struct {
	entities   []*entity.Entity
	components []string
}

type SystemState interface {
	MouseInputs() map[ebiten.MouseButton]Control
	KeyInputs() map[ebiten.Key]Control
	SetControl(Control, *InputData)
	Events() []interface{}
	ActivateWorlds([]string)
}

func (s *system) GetComponents() []string {
	return s.components
}

func (s *system) GetEntityCount() int {
	return len(s.entities)
}

func (s *system) AddEntity(e *entity.Entity) {
	s.entities = append(s.entities, e)
}
