package entity

import (
	"github.com/leviceccato/clarity/component"
)

var entityId = 0

type Entity struct {
	id         int
	position   *component.Position
	size       *component.Size
	appearance *component.Appearance
	animation  *component.Animation
}

func newEntity() *Entity {
	e := &Entity{
		id: entityId,
	}
	entityId += 1
	return e
}
