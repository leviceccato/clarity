package entity

import (
	"github.com/leviceccato/clarity/component"
)

type Entity struct {
	Position   *component.PositionComponent
	Size       *component.SizeComponent
	Appearance *component.AppearanceComponent
	Hover      *component.HoverComponent
	Playable   *component.PlayableComponent
	Text       *component.TextComponent
	Cursor     *component.CursorComponent
}

func NewEntity() *Entity {
	return &Entity{}
}
