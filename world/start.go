package world

import (
	"fmt"

	"github.com/leviceccato/clarity/config"
	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/system"

	"github.com/hajimehoshi/ebiten/v2"
)

type startWorld struct {
	world
}

func NewStartWorld(state system.SystemState) (*startWorld, error) {
	w := &startWorld{}
	w.name = "start"

	w.systems = []WorldSystem{
		system.NewDrawSystem(state),
		system.NewAnimationSystem(),
		system.NewInputSystem(state),
		system.NewEventSystem(state),
		system.NewPlayableSystem(state),
	}

	// Create and position player
	player, err := entity.NewPlayerEntity()
	if err != nil {
		return nil, fmt.Errorf("creating player entity: %s", err)
	}
	player.Position.X = float64(config.RenderWidth) / 2
	player.Position.Y = float64(config.RenderHeight) / 2

	w.entities = []*entity.Entity{
		player,
	}
	w.updateSystems()

	return w, nil
}

func (w *startWorld) Load() {}

func (w *startWorld) Update() {}

func (w *startWorld) Draw(screen *ebiten.Image) {}

func (w *startWorld) Enter() {}

func (w *startWorld) Exit() {}
