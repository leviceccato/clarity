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

func NewTitleWorld(state system.SystemState) (*titleWorld, error) {
	w := &titleWorld{}
	w.name = "title"

	w.systems = []WorldSystem{
		system.NewDrawSystem(state),
		system.NewAnimationSystem(),
		system.NewInputSystem(state),
		system.NewEventSystem(state),
		system.NewPlayableSystem(state),
		system.NewCursorSystem(state),
		system.NewHoverSystem(state),
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
	buttonYSpacing := 5.0
	buttonYStart := 120.0
	startButton, err := entity.NewButtonEntity(&entity.ButtonEntityOptions{
		X:          (config.RenderWidth / 2) - (buttonWidth / 2),
		Y:          buttonYStart,
		Width:      buttonWidth,
		Height:     buttonHeight,
		Padding:    10,
		Text:       util.Trans("start"),
		Font:       *state.Font("lana_pixel"),
		Color:      config.ColFgTitle,
		Image:      "sprite/cursor.png",
		IsCentered: true,
	})
	if err != nil {
		return nil, fmt.Errorf("creating title button entity: %s", err)
	}
	exitButton, err := entity.NewButtonEntity(&entity.ButtonEntityOptions{
		X:          (config.RenderWidth / 2) - (buttonWidth / 2),
		Y:          buttonYStart + buttonHeight + buttonYSpacing,
		Width:      buttonWidth,
		Height:     buttonHeight,
		Padding:    10,
		Text:       util.Trans("exit"),
		Font:       *state.Font("lana_pixel"),
		Color:      config.ColFgTitle,
		Image:      "sprite/cursor.png",
		IsCentered: true,
	})
	if err != nil {
		return nil, fmt.Errorf("creating title button entity: %s", err)
	}

	w.entities = []*entity.Entity{
		background,
		startButton,
		exitButton,
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
