package game

import (
	"image/color"

	"github.com/leviceccato/clarity/asset"
	"github.com/leviceccato/clarity/engine"
	"github.com/leviceccato/clarity/logger"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type command int

const (
	commandJump command = iota
	commandMoveLeft
	commandMoveRight
	commandMoveDown
	commandMoveUp
)

type inputData struct {
	isStart bool
	x, y    float64
}

type Game struct {
	*engine.Game

	// Globals
	fonts      map[string]font.Face
	colors     map[string]color.NRGBA
	translator *asset.Translator
	logger     *logger.Logger

	// Settings
	InputBindings map[command][]any

	// State
	isCursorHovering bool
	mouseInputs      map[ebiten.MouseButton]command
	keyInputs        map[ebiten.Key]command
	inputs           map[command]*inputData
}

func (g *Game) setIsCursorHovering(to bool) {
	g.isCursorHovering = to
}

func (g *Game) setInput(cmd command, data *inputData) {
	g.inputs[cmd] = data
}
