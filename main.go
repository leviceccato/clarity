package main

import (
	"fmt"
	"image"

	"github.com/leviceccato/clarity/asset"
	"github.com/leviceccato/clarity/config"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Translations
	err := asset.InitTranslations()
	if err != nil {
		panic(fmt.Sprintf("initialising translations: %s", err))
	}

	// Fonts
	fonts, err := asset.LoadFonts(map[string]string{
		"lana_pixel": "font/lana_pixel.ttf",
	})
	if err != nil {
		panic(fmt.Sprintf("loading fonts: %s", err))
	}

	// Add icons
	icon32, err := asset.LoadIcon("icon.iconset/icon_32x32.png")
	if err != nil {
		panic(fmt.Sprintf("loading icon_32: %s", err))
	}
	icon16, err := asset.LoadIcon("icon.iconset/icon_16x16.png")
	if err != nil {
		panic(fmt.Sprintf("loading icon_16: %s", err))
	}
	ebiten.SetWindowIcon([]image.Image{icon32, icon16})

	ebiten.SetWindowSize(config.RenderWidth*2, config.RenderHeight*2)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle(config.Title)

	gameInstance := game.NewGame()

	// Start game
	err = ebiten.RunGame(gameInstance)
	if err != nil {
		panic(fmt.Sprintf("running game: %s", err))
	}
}
