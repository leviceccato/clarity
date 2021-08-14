package entity

import (
	"fmt"

	"github.com/leviceccato/clarity/component"
)

func NewPlayerEntity() (*Entity, error) {
	e := NewEntity()
	e.Position = &component.Position{X: 10, Y: 10}
	e.Size = &component.Size{Width: 10, Height: 10}
	appearance, err := component.NewAppearance("assets/player.png", "assets/player.json")
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %s", err)
	}
	e.Appearance = appearance
	return e, nil
}
