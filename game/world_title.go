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
		"click",
	})

	w.Enter = func() error {
		background, err := newImageEntity(g, &imageEntityOptions{path: "sprite/title_bg.png"})
		if err != nil {
			return fmt.Errorf("creating background entity: %w", err)
		}

		cursor, err := newCursorEntity(g)
		if err != nil {
			return fmt.Errorf("creating cursor entity: %w", err)
		}

		buttonWidth := 74.0
		buttonHeight := 30.0
		buttonYSpacing := 1.0
		buttonYStart := 108.0

		startButton, err := newButtonEntity(g, &buttonEntityOptions{
			x:             (float64(g.RenderWidth) / 2) - (buttonWidth / 2),
			y:             buttonYStart,
			z:             50,
			width:         buttonWidth,
			height:        buttonHeight,
			padding:       10,
			text:          g.translator.MustTrans("start", nil),
			textTransform: textComponentTransformUppercase,
			font:          g.fonts["lana_pixel"],
			color:         g.colors["fg-title"],
			image:         "sprite/title_button.png",
			animation:     "sprite/title_button.json",
			isCentered:    true,
			action: func() {
				g.ActivateWorlds("start")
			},
		})
		if err != nil {
			return fmt.Errorf("creating title button entity: %w", err)
		}
		optionsButton, err := newButtonEntity(g, &buttonEntityOptions{
			x:             (float64(g.RenderWidth) / 2) - (buttonWidth / 2),
			y:             buttonYStart + buttonHeight + buttonYSpacing,
			z:             50,
			width:         buttonWidth,
			height:        buttonHeight,
			padding:       10,
			textTransform: textComponentTransformUppercase,
			text:          g.translator.MustTrans("options", nil),
			font:          g.fonts["lana_pixel"],
			color:         g.colors["fg-title"],
			image:         "sprite/title_button.png",
			animation:     "sprite/title_button.json",
			isCentered:    true,
			action: func() {
				fmt.Println("clicked options")
			},
		})
		if err != nil {
			return fmt.Errorf("creating title button entity: %w", err)
		}

		exitButton, err := newButtonEntity(g, &buttonEntityOptions{
			x:             (float64(g.RenderWidth) / 2) - (buttonWidth / 2),
			y:             buttonYStart + (2 * (buttonHeight + buttonYSpacing)),
			z:             50,
			width:         buttonWidth,
			height:        buttonHeight,
			padding:       10,
			textTransform: textComponentTransformUppercase,
			text:          g.translator.MustTrans("exit", nil),
			font:          g.fonts["lana_pixel"],
			color:         g.colors["fg-title"],
			image:         "sprite/title_button.png",
			animation:     "sprite/title_button.json",
			isCentered:    true,
			action: func() {
				g.Quit()
			},
		})
		if err != nil {
			return fmt.Errorf("creating title button entity: %w", err)
		}

		g.AddEntities(w,
			background,
			startButton,
			optionsButton,
			exitButton,
			cursor,
		)

		return nil
	}

	return w
}
