package game

import (
	"fmt"
	"image"
	"image/color"

	"github.com/leviceccato/clarity/asset"
	"github.com/leviceccato/clarity/engine"
	"github.com/leviceccato/clarity/logger"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/text/language"
)

type Options struct {
	RenderWidth, RenderHeight int

	Title string

	Logger *logger.Logger
}

func CreateAndRun(options *Options) error {
	g := &Game{Game: engine.NewGame()}

	// Init translations
	translator := asset.NewTranslator()
	err := translator.AddLocalizer("translation/en.json", language.English)
	if err != nil {
		return fmt.Errorf("adding localizer: %w", err)
	}
	translator.Lang = language.English
	g.translator = translator

	// Load fonts
	lanaPixel, err := asset.LoadFont(
		"font/lana_pixel.ttf",
		11,
		72,
		font.HintingNone,
	)
	if err != nil {
		return fmt.Errorf("loading font: %w", err)
	}
	g.fonts = map[string]font.Face{
		"lana_pixel": lanaPixel,
	}

	// Add icons
	icon16, err := asset.LoadIcon("icon.iconset/icon_16x16.png")
	if err != nil {
		return fmt.Errorf("loading icon_16: %w", err)
	}
	icon32, err := asset.LoadIcon("icon.iconset/icon_32x32.png")
	if err != nil {
		return fmt.Errorf("loading icon_32: %w", err)
	}
	ebiten.SetWindowIcon([]image.Image{icon32, icon16})

	g.RenderWidth = options.RenderWidth
	g.RenderHeight = options.RenderHeight
	ebiten.SetWindowSize(g.RenderWidth*2, g.RenderHeight*2)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle(options.Title)

	// Set colours
	g.colors = map[string]color.NRGBA{
		"fg-title": {255, 240, 157, 255},
	}

	// Set inputs
	g.inputs = map[command]*inputData{}
	g.InputBindings = map[command][]any{
		commandJump:      {ebiten.KeySpace},
		commandMoveLeft:  {ebiten.KeyA, ebiten.KeyArrowLeft},
		commandMoveRight: {ebiten.KeyD, ebiten.KeyArrowRight},
		commandMoveUp:    {ebiten.KeyW, ebiten.KeyArrowUp},
		commandMoveDown:  {ebiten.KeyS, ebiten.KeyArrowDown},
	}

	// Init systems
	g.AddSystems(
		newAnimationSystem(g),
		newCursorSystem(g),
		newDrawSystem(g),
		newHoverSystem(g),
		newInputSystem(g),
		newPlayableSystem(g),
	)

	// Init worlds
	g.AddWorlds(
		newTitleWorld(g),
		newStartWorld(g),
	)

	// Set title as active world
	g.ActivateWorlds("title")

	// Run the game
	err = g.Run()
	if err != nil {
		return fmt.Errorf("running game: %w", err)
	}

	return nil
}
