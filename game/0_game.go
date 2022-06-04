package game

import (
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
	shouldQuit       bool
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

func (g *Game) initInputs(bindings map[command][]any) {
	g.InputBindings = bindings
	g.inputs = make(map[command]*inputData, len(bindings))
}

func (g *Game) setInput(cmd command, data *inputData) {
	g.inputs[cmd] = data
}

func (g Game) hasInput(cmd command) bool {
	data := g.inputs[cmd]
	if data == nil {
		return false
	}
	return (data.progress == inputStart) || (data.progress == inputMiddle)
}

func (g Game) isInputStarting(cmd command) bool {
	data := g.inputs[cmd]
	if data == nil {
		return false
	}
	return data.progress == inputStart
}

func (g Game) isInputEnding(cmd command) bool {
	data := g.inputs[cmd]
	if data == nil {
		return false
	}
	return data.progress == inputEnd
}
