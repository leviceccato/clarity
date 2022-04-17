package game

import (
	"image/color"
	"os"

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
	InputBindings map[command][]any

	// State
	isCursorHovering bool
	inputs           map[command]*inputData
}

func (g Game) quit(code int) {
	g.logger.Info.Printf("quit with code: %v", code)

	os.Exit(code)
}

func (g *Game) setIsCursorHovering(to bool) {
	g.isCursorHovering = to
}

type command int

const (
	commandJump command = iota
	commandMoveLeft
	commandMoveRight
	commandMoveDown
	commandMoveUp
)

type inputData struct {
	// Track beginning and end of input
	isStart bool

	// After input ended, mark as expired so struct is deleted in next frame
	shouldExpire bool

	x, y float64
}

func (g *Game) setInput(cmd command, data *inputData) {
	g.inputs[cmd] = data
}
