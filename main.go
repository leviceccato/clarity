package main

import (
	"fmt"

	"github.com/leviceccato/clarity/utility"
	"github.com/leviceccato/clarity/world"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	renderWidth  = 480
	renderHeight = 270
)

type game struct {
	state *state
}

func (g *game) Update() error {
	g.state.update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.state.draw(screen)
}

func (g *game) Layout(w, h int) (int, int) {
	return renderWidth, renderHeight
}

func main() {
	// Initilisations
	utility.InitTranslations()
	utility.InitViewport(renderWidth, renderHeight)

	ebiten.SetWindowSize(renderWidth*2, renderHeight*2)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Clarity")

	// Create state and load game data
	mainState := newState()
	mainState.renderWidth = renderWidth
	mainState.renderHeight = renderHeight
	startWorld, err := world.NewStartWorld(mainState)
	if err != nil {
		fmt.Printf("creating start world: %s", err)
	}
	mainState.loadWorld(startWorld)
	mainState.activateWorlds([]string{"start"})

	err = ebiten.RunGame(&game{
		state: mainState,
	})
	if err != nil {
		fmt.Printf("running game: %s", err)
	}
}
