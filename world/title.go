package world

import (
	"fmt"

	"github.com/leviceccato/clarity/config"
	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/system"
	"github.com/leviceccato/clarity/util"

	"github.com/hajimehoshi/ebiten/v2"
)

type titleWorld struct {
	world
}

func NewTitleWorld(state interface{}) (*titleWorld, error) {
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
		system.NewHoverSystem(systemState),
	}

	cursor, err := entity.NewCursorEntity()
	if err != nil {
		return nil, fmt.Errorf("creating cursor entity: %s", err)
	}

	background, err := entity.NewImageEntity("sprite/title_bg.png")
	if err != nil {
		return nil, fmt.Errorf("creating background entity: %s", err)
	}

	buttonWidth := 100.0
	buttonHeight := 50.0
	titleButton, err := entity.NewButtonEntity(&entity.ButtonEntityOptions{
		X:       (config.RenderWidth / 2) - (buttonWidth / 2),
		Y:       120,
		Width:   buttonWidth,
		Height:  buttonHeight,
		Padding: 10,
		Text:    util.Trans("start"),
		Font:    *systemState.Font("lana_pixel"),
		Color:   config.ColFgTitle,
		Image:   "sprite/cursor.png",
	})
	if err != nil {
		return nil, fmt.Errorf("creating title button entity: %s", err)
	}

	w.entities = []*entity.Entity{
		background,
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
