package world

import (
	"fmt"

	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/system"
	"github.com/leviceccato/clarity/utility"

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
		system.NewDrawSystem(systemState),
		system.NewAnimationSystem(),
		system.NewInputSystem(systemState),
		system.NewEventSystem(systemState),
		system.NewPlayableSystem(systemState),
	}

	titleBg, err := entity.NewTitleBgEntity()
	if err != nil {
		return nil, fmt.Errorf("creating title bg entity: %s", err)
	}

	titleButton, err := entity.NewButtonEntity(&entity.ButtonEntityOptions{
		X:      50,
		Y:      50,
		Width:  100,
		Height: 50,
		Text:   utility.Trans("start"),
		Image:  "assets/cursor.png",
	})
	if err != nil {
		return nil, fmt.Errorf("creating title button entity: %s", err)
	}

	w.entities = []*entity.Entity{
		titleBg,
		titleButton,
	}
	w.updateSystems()

	return w, nil
}

func (w *titleWorld) Load() {}

func (w *titleWorld) Update() {}

func (w *titleWorld) Draw(screen *ebiten.Image) {}

func (w *titleWorld) Enter() {}

func (w *titleWorld) Exit() {}
