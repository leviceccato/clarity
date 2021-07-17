package entity

import (
	"github.com/leviceccato/clarity/component"
)

func NewPlayerEntity() *Entity {
	e := newEntity()
	e.position = &component.Position{}
	e.size = &component.Size{}
	e.appearance = &component.Appearance{}
	e.animation = &component.Animation{}
	return e
}
