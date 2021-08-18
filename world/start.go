package world

import (
	"fmt"

	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/system"

	"github.com/hajimehoshi/ebiten/v2"
)

type startWorld struct {
	world
}

func NewStartWorld(state gameState) (*startWorld, error) {
	w := &startWorld{}
	w.name = "start"

	systemState := state.(system.SystemState)
	w.AddSystem(system.NewDrawSystem())
	w.AddSystem(system.NewAnimationSystem())
	w.AddSystem(system.NewInputSystem(systemState))
	w.AddSystem(system.NewEventSystem(systemState))
	w.AddSystem(system.NewPlayableSystem(systemState))

	// Create and position player
	player, err := entity.NewPlayerEntity()
	if err != nil {
		return nil, fmt.Errorf("creating player entity: %s", err)
	}
	player.Position.X = float64(state.RenderWidth()) / 2
	player.Position.Y = float64(state.RenderHeight()) / 2
	w.AddEntity(player)

	w.updateSystems()

	return w, nil
}

func (w *startWorld) Load() {}

func (w *startWorld) Update() {}

func (w *startWorld) Draw(screen *ebiten.Image) {}

func (w *startWorld) Enter() {}

func (w *startWorld) Exit() {}
