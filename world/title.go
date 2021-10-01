package world

import (
	"fmt"

	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/system"
	"github.com/leviceccato/clarity/util"

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
		system.NewCursorSystem(systemState),
	}

	titleBg, err := entity.NewTitleBgEntity()
	if err != nil {
		return nil, fmt.Errorf("creating title bg entity: %s", err)
	}

	titleButton, err := entity.NewButtonEntity(&entity.ButtonEntityOptions{
		X:       50,
		Y:       50,
		Width:   100,
		Height:  50,
		Padding: 10,
		Text:    util.Trans("start"),
		Font:    *systemState.Font("lana_pixel"),
		Color:   systemState.Color("fg_title"),
		Image:   "sprite/cursor.png",
	})
	if err != nil {
		return nil, fmt.Errorf("creating title button entity: %s", err)
	}

	cursor, err := entity.NewCursorEntity()
	if err != nil {
		return nil, fmt.Errorf("creating cursor entity: %s", err)
	}

	w.entities = []*entity.Entity{
		titleBg,
		titleButton,
		cursor,
	}
	w.updateSystems()

	return w, nil
}

func (w *titleWorld) Load() {}

func (w *titleWorld) Update() {}

func (w *titleWorld) Draw(screen *ebiten.Image) {}

func (w *titleWorld) Enter() {}

func (w *titleWorld) Exit() {}
