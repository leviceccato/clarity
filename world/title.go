package world

import (
	"fmt"

	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/system"

	"github.com/hajimehoshi/ebiten/v2"
)

type titleWorld struct {
	world
}

func NewTitleWorld(state gameState) (*titleWorld, error) {
	w := &titleWorld{}
	w.name = "title"

	systemState := state.(system.SystemState)
	w.systems = []WorldSystem{
		system.NewDrawSystem(),
		system.NewAnimationSystem(),
		system.NewInputSystem(systemState),
		system.NewEventSystem(systemState),
		system.NewPlayableSystem(systemState),
	}

	// Create and position player
	player, err := entity.NewPlayerEntity()
	if err != nil {
		return nil, fmt.Errorf("creating player entity: %s", err)
	}
	player.Position.X = 10
	player.Position.Y = 10
	w.AddEntity(player)

	w.updateSystems()

	return w, nil
}

func (w *titleWorld) Load() {}

func (w *titleWorld) Update() {}

func (w *titleWorld) Draw(screen *ebiten.Image) {}

func (w *titleWorld) Enter() {}

func (w *titleWorld) Exit() {}
