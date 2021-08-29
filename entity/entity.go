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
	Hover      *component.HoverComponent
	Playable   *component.PlayableComponent
	Text       *component.TextComponent
}

func NewEntity() *Entity {
	e := &Entity{
		id: entityId,
	}
	entityId += 1
	return e
}
