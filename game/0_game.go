package game

import (
	"image/color"

	"github.com/leviceccato/clarity/asset"
	"github.com/leviceccato/clarity/engine"
	"github.com/leviceccato/clarity/logger"

	"golang.org/x/image/font"
)

type Game struct {
	*engine.Game

	// Globals
	fonts      map[string]font.Face
	colors     map[string]color.NRGBA
	translator *asset.Translator
	logger     *logger.Logger

	// Settings
	InputBindings map[string]string

	// State
	isCursorHovering bool
	// Controls         map[component.Control]*component.InputData
	// MouseInputs      map[ebiten.MouseButton]component.Control
	// KeyInputs        map[ebiten.Key]component.Control
}

func (g *Game) setIsCursorHovering(to bool) {
	g.isCursorHovering = to
}
