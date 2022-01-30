package world

import (
	"fmt"

	"github.com/leviceccato/clarity/config"
	"github.com/leviceccato/clarity/entity"
	"github.com/leviceccato/clarity/game"
	"github.com/leviceccato/clarity/system"

	"github.com/hajimehoshi/ebiten/v2"
)

type titleWorld struct {
	world
}

func NewTitleWorld(s *game.State) *titleWorld {
	w := &titleWorld{}
	w.name = "title"

	w.systems = []game.GameSystem{
		system.NewDrawSystem(s),
		system.NewAnimationSystem(s),
		system.NewInputSystem(s),
		system.NewEventSystem(s),
		system.NewPlayableSystem(s),
		system.NewCursorSystem(s),
		system.NewHoverSystem(s),
	}

	cursor, err := entity.NewCursorEntity()
	if err != nil {
		panic(fmt.Sprintf("creating cursor entity: %s", err))
	}

	background, err := entity.NewImageEntity("sprite/title_bg.png")
	if err != nil {
		panic(fmt.Sprintf("creating background entity: %s", err))
	}

	buttonWidth := 74.0
	buttonHeight := 30.0
	buttonYSpacing := 5.0
	buttonYStart := 120.0
	startButton, err := entity.NewButtonEntity(&entity.ButtonEntityOptions{
		X:          (config.RenderWidth / 2) - (buttonWidth / 2),
		Y:          buttonYStart,
		Width:      buttonWidth,
		Height:     buttonHeight,
		Padding:    10,
		Text:       "start",
		Font:       s.Fonts["lana_pixel"],
		Color:      s.Colors["fg-title"],
		Image:      "sprite/title_button.png",
		IsCentered: true,
	})
	if err != nil {
		panic(fmt.Sprintf("creating title button entity: %s", err))
	}
	exitButton, err := entity.NewButtonEntity(&entity.ButtonEntityOptions{
		X:          (config.RenderWidth / 2) - (buttonWidth / 2),
		Y:          buttonYStart + buttonHeight + buttonYSpacing,
		Width:      buttonWidth,
		Height:     buttonHeight,
		Padding:    10,
		Text:       "exit",
		Font:       s.Fonts["lana_pixel"],
		Color:      s.Colors["fg-title"],
		Image:      "sprite/title_button.png",
		IsCentered: true,
	})
	if err != nil {
		panic(fmt.Sprintf("creating title button entity: %s", err))
	}

	w.entities = []*entity.Entity{
		background,
		startButton,
		exitButton,
		cursor,
	}
	w.updateSystems()

	return w
}

func (w *titleWorld) Load(s *game.State) {}

func (w *titleWorld) Update(s *game.State) {}

func (w *titleWorld) Draw(s *game.State, screen *ebiten.Image) {}

func (w *titleWorld) Enter(s *game.State) {}

func (w *titleWorld) Exit(s *game.State) {}
