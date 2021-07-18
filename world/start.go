package world

import (
	"fmt"

	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/system"

	"github.com/hajimehoshi/ebiten/v2"
)

type start struct {
	world
}

func NewStartWorld() (*start, error) {
	w := &start{}
	w.name = "start"

	w.AddSystem(system.NewDrawSystem())
	player, err := entity.NewPlayerEntity()
	if err != nil {
		return nil, fmt.Errorf("creating player entity: %s", err)
	}
	w.AddEntity(player)
	w.updateSystems()

	return w, nil
}

func (w *start) Load() {
}

func (w *start) Update() {
}

func (w *start) Draw(screen *ebiten.Image) {
}

func (w *start) Enter() {
}

func (w *start) Exit() {
}
