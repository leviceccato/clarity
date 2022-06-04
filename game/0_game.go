package game

import (
	"errors"
	"image/color"

	"github.com/leviceccato/clarity/asset"
	"github.com/leviceccato/clarity/engine"
	"github.com/leviceccato/clarity/logger"

	"github.com/hajimehoshi/ebiten/v2"
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
	savePath      string

	// State
	isCursorHovering bool
	inputs           map[command]*inputData
	isDebug          bool
	windowTitle      string
}

var CloseError = errors.New("window closed")

func (g *Game) quit() error {
	return CloseError
}

func (g *Game) setIsCursorHovering(to bool) {
	g.isCursorHovering = to
}

func (g *Game) setTitle(to string) {
	title := to
	if g.isDebug {
		title = title + " - DEBUG"
	}

	ebiten.SetWindowTitle(title)
	g.windowTitle = to
}

func (g *Game) toggleDebug() {
	g.isDebug = !g.isDebug
	g.setTitle(g.windowTitle)
}

type command int

const (
	commandJump command = iota
	commandMoveLeft
	commandMoveRight
	commandMoveDown
	commandMoveUp
	commandToggleMenu
	commandClick
	commandToggleDebug
)

type inputProgress int

const (
	inputStart inputProgress = iota
	inputMiddle
	inputEnd
)

type inputData struct {
	// Track start, middle and end of input
	progress inputProgress

	x, y float64
}

func (g *Game) setInput(cmd command, data *inputData) {
	g.inputs[cmd] = data
}
