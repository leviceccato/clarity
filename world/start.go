package world

import (
	"fmt"
	"math/rand"

	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/system"

	"github.com/hajimehoshi/ebiten/v2"
)

type start struct {
	world
}

func NewStartWorld(state gameState) (*start, error) {
	w := &start{}
	w.name = "start"

	w.AddSystem(system.NewDrawSystem())

	var randX float64
	var randY float64
	for i := 0; i < 500; i++ {
		randX = rand.Float64() * float64(state.GetWindowWidth())
		randY = rand.Float64() * float64(state.GetWindowHeight())
		player, err := entity.NewPlayerEntity()
		if err != nil {
			return nil, fmt.Errorf("creating player entity: %s", err)
		}
		player.Position.X = randX
		player.Position.Y = randY
		w.AddEntity(player)
	}

	w.updateSystems()

	return w, nil
}

func (w *start) Load() {}

func (w *start) Update() {}

func (w *start) Draw(screen *ebiten.Image) {}

func (w *start) Enter() {}

func (w *start) Exit() {}
