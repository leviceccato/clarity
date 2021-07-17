package system

import (
	"github.com/leviceccato/clarity/entity"
)

type system struct {
	entities   []*entity.Entity
	components []string
}

func (s *system) GetComponents() []string {
	return s.components
}

func (s *system) AddEntity(e *entity.Entity) {
	s.entities = append(s.entities, e)
}
