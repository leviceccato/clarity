package entity

import (
	"github.com/leviceccato/clarity/component"
)

var entityId = 0

type Entity struct {
	id         int
	Position   *component.Position
	Size       *component.Size
	Appearance *component.Appearance
}

func NewEntity() *Entity {
	e := &Entity{
		id: entityId,
	}
	entityId += 1
	return e
}
