package game

import (
	"fmt"

	"github.com/leviceccato/clarity/engine"
)

func newTitleWorld(g *Game) *engine.World {
	w := engine.NewWorld("title", []string{
		"cursor",
		"animation",
		"input",
		"playable",
		"hover",
		"draw",
	})

	w.Enter = func() error {
		cursor, err := newCursorEntity(g)
		if err != nil {
			return fmt.Errorf("creating cursor entity: %w", err)
		}

		background, err := newImageEntity(g, "sprite/title_bg.png")
		if err != nil {
			return fmt.Errorf("creating background entity: %w", err)
		}

		buttonWidth := 74.0
		buttonHeight := 30.0
		buttonYSpacing := 5.0
		buttonYStart := 120.0
		startButton, err := newButtonEntity(g, &buttonEntityOptions{
			x:          (float64(g.RenderWidth) / 2) - (buttonWidth / 2),
			y:          buttonYStart,
			width:      buttonWidth,
			height:     buttonHeight,
			padding:    10,
			text:       g.translator.MustTrans("start", nil),
			font:       g.fonts["lana_pixel"],
			color:      g.colors["fg-title"],
			image:      "sprite/title_button.png",
			animation:  "sprite/title_button.json",
			isCentered: true,
		})
		if err != nil {
			return fmt.Errorf("creating title button entity: %w", err)
		}
		exitButton, err := newButtonEntity(g, &buttonEntityOptions{
			x:          (float64(g.RenderWidth) / 2) - (buttonWidth / 2),
			y:          buttonYStart + buttonHeight + buttonYSpacing,
			width:      buttonWidth,
			height:     buttonHeight,
			padding:    10,
			text:       g.translator.MustTrans("exit", nil),
			font:       g.fonts["lana_pixel"],
			color:      g.colors["fg-title"],
			image:      "sprite/title_button.png",
			animation:  "sprite/title_button.json",
			isCentered: true,
		})
		if err != nil {
			return fmt.Errorf("creating title button entity: %w", err)
		}

		g.AddEntities(w,
			background,
			startButton,
			exitButton,
			cursor,
		)

		return nil
	}

	return w
}
