package main

import (
	"fmt"
	"image"

	"github.com/leviceccato/clarity/asset"
	"github.com/leviceccato/clarity/util"
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
	// Initialisations
	err := util.InitTranslations()
	if err != nil {
		fmt.Printf("initialising translations: %s", err)
		return
	}
	fonts, err := util.LoadFonts(map[string]string{
		"lana_pixel": "font/lana_pixel.ttf",
	})
	if err != nil {
		fmt.Printf("loading fonts: %s", err)
		return
	}
	// Add icons
	icon32, err := util.LoadIcon("icon.iconset/icon_32x32.png")
	if err != nil {
		fmt.Printf("loading icon_32: %s", err)
		return
	}
	icon16, err := util.LoadIcon("icon.iconset/icon_16x16.png")
	if err != nil {
		fmt.Printf("loading icon_16: %s", err)
		return
	}
	ebiten.SetWindowIcon([]image.Image{icon32, icon16})
	ebiten.SetWindowSize(renderWidth*2, renderHeight*2)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle(asset.ConfTitle)

	// Create state and load game data
	mainState := newState()
	mainState.fonts = fonts
	mainState.renderWidth = renderWidth
	mainState.renderHeight = renderHeight
	startWorld, err := world.NewStartWorld(mainState)
	if err != nil {
		fmt.Printf("creating start world: %s", err)
		return
	}
	titleWorld, err := world.NewTitleWorld(mainState)
	if err != nil {
		fmt.Printf("creating title world: %s", err)
		return
	}
	mainState.loadWorld(startWorld)
	mainState.loadWorld(titleWorld)
	mainState.ActivateWorlds([]string{"start"})

	err = ebiten.RunGame(&game{
		state: mainState,
	})
	if err != nil {
		fmt.Printf("running game: %s", err)
		return
	}
}
