package game

import (
	"fmt"

	"github.com/leviceccato/clarity/engine"
)

func newMenuAssemblage(g *Game) ([]*engine.Entity, error) {
	background, err := newImageEntity(g, &imageEntityOptions{z: 50, path: "sprite/title_bg.png"})
	if err != nil {
		return nil, fmt.Errorf("creating background entity: %w", err)
	}

	buttonWidth := 74.0
	buttonHeight := 30.0
	buttonYStart := 108.0

	exitButton, err := newButtonEntity(g, &buttonEntityOptions{
		x:             (float64(g.RenderWidth) / 2) - (buttonWidth / 2),
		y:             buttonYStart,
		z:             60,
		width:         buttonWidth,
		height:        buttonHeight,
		padding:       10,
		text:          g.translator.MustTrans("exit", nil),
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
		return nil, fmt.Errorf("creating title button entity: %w", err)
	}

	return []*engine.Entity{background, exitButton}, nil
}
