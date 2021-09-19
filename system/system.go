package system

import (
	"image/color"

	"github.com/leviceccato/clarity/entity"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2"
)

type system struct {
	entities   []*entity.Entity
	components []string
}

type SystemState interface {
	MouseInputs() map[ebiten.MouseButton]Control
	KeyInputs() map[ebiten.Key]Control
	Controls() map[Control]*InputData
	SetControl(Control, *InputData)
	Events() []interface{}
	SetEvents([]interface{})
	AddEvent(interface{})
	ActivateWorlds([]string)
	Font(string) *font.Face
	Color(string) color.NRGBA
}

func (s *system) Components() []string {
	return s.components
}

func (s *system) EntityCount() int {
	return len(s.entities)
}

func (s *system) AddEntity(e *entity.Entity) {
	s.entities = append(s.entities, e)
}
