package world

import (
	"fmt"

	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/game"
	"github.com/leviceccato/clarity/system"

	"github.com/hajimehoshi/ebiten/v2"
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

func (w *startWorld) Load(s *game.State) {}

func (w *startWorld) Update(s *game.State) {}

func (w *startWorld) Draw(s *game.State, screen *ebiten.Image) {}

func (w *startWorld) Enter(s *game.State) {}

func (w *startWorld) Exit(s *game.State) {}
