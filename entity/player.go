package entity

import (
	"github.com/leviceccato/clarity/component"
)

func NewPlayerEntity() *Entity {
	e := NewEntity()
	e.Position = &component.Position{}
	e.Size = &component.Size{}
	e.Appearance = &component.Appearance{}
	e.Animation = &component.Animation{}
	return e
}
