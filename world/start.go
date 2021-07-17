package world

import (
	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/system"
)

type start struct {
	world
}

func NewStartWorld() *start {
	w := &start{}
	w.name = "start"

	w.AddSystem(system.NewDrawSystem())
	w.AddEntity(entity.NewPlayerEntity())
	w.updateSystems()

	return w
}

func (w *start) Load() {
}

func (w *start) Update() {
}

func (w *start) Draw() {
}

func (w *start) Enter() {
}

func (w *start) Exit() {

}
