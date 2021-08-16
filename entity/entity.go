package entity

import (
	"github.com/leviceccato/clarity/component"
)

var entityId = 0

type Entity struct {
	id         int
	Position   *component.PositionComponent
	Size       *component.SizeComponent
	Appearance *component.AppearanceComponent
}

func NewEntity() *Entity {
	e := &Entity{
		id: entityId,
	}
	entityId += 1
	return e
}
