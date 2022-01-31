package world

import (
	"fmt"

	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/game"
	"github.com/leviceccato/clarity/system"
)

type startWorld struct {
	world
}

func NewStartWorld(s *game.State) *startWorld {
	w := &startWorld{}
	w.name = "start"

	w.systems = []game.GameSystem{
		system.NewDrawSystem(s),
		system.NewAnimationSystem(s),
		system.NewInputSystem(s),
		system.NewEventSystem(s),
		system.NewPlayableSystem(s),
	}

	// Create and position player
	player, err := entity.NewPlayerEntity()
	if err != nil {
		panic(fmt.Errorf("creating player entity: %s", err))
	}
	player.Position.X = float64(s.RenderWidth) / 2
	player.Position.Y = float64(s.RenderHeight) / 2

	w.entities = []*entity.Entity{
		player,
	}
	w.updateSystems()

	return w
}
