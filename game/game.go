package game

import (
	"fmt"
	"image"
	"image/color"

	"github.com/leviceccato/clarity/asset"
	"github.com/leviceccato/clarity/component"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/text/language"
)

type Options struct {
	RenderWidth, RenderHeight int
	Title                     string
}

func NewGame(options Options) (*State, error) {
	g := &State{}

	// Config
	g.RenderWidth = options.RenderWidth
	g.RenderHeight = options.RenderHeight
	g.Title = options.Title

	// Initialisations
	translator := asset.NewTranslator()
	err := translator.AddLocalizer("translation/en.json", language.English)
	if err != nil {
		return nil, fmt.Errorf("adding localizer: %s", err)
	}

	// Load fonts
	lanaPixel, err := asset.LoadFont("font/lana_pixel.ttf")
	if err != nil {
		return nil, fmt.Errorf("loading font: %s", err)
	}
	g.Fonts = map[string]font.Face{
		"lana_pixel": lanaPixel,
	}

	// Add icons
	icon32, err := asset.LoadIcon("icon.iconset/icon_32x32.png")
	if err != nil {
		return nil, fmt.Errorf("loading icon_32: %s", err)
	}
	icon16, err := asset.LoadIcon("icon.iconset/icon_16x16.png")
	if err != nil {
		return nil, fmt.Errorf("loading icon_16: %s", err)
	}

	ebiten.SetWindowIcon([]image.Image{icon32, icon16})
	ebiten.SetWindowSize(g.RenderWidth*2, g.RenderHeight*2)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle(g.Title)

	g.worlds = map[string]gameWorld{}
	g.MouseInputs = map[ebiten.MouseButton]component.Control{
		ebiten.MouseButtonLeft: component.ControlClick,
	}
	g.KeyInputs = map[ebiten.Key]component.Control{
		ebiten.KeySpace:      component.ControlJump,
		ebiten.KeyArrowUp:    component.ControlUp,
		ebiten.KeyW:          component.ControlUp,
		ebiten.KeyArrowLeft:  component.ControlLeft,
		ebiten.KeyA:          component.ControlLeft,
		ebiten.KeyArrowRight: component.ControlRight,
		ebiten.KeyD:          component.ControlRight,
		ebiten.KeyArrowDown:  component.ControlDown,
		ebiten.KeyS:          component.ControlDown,
		ebiten.KeyEscape:     component.ControlMenu,
		ebiten.KeyBackquote:  component.ControlDebug,
	}
	g.UpdateControls()

	g.Colors = map[string]color.NRGBA{
		"fg-title": {255, 240, 157, 255},
	}

	return g, nil
}
