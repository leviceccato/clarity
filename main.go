package main

import (
	"os"

	"github.com/leviceccato/clarity/game"
	"github.com/leviceccato/clarity/logger"
)

func main() {
	// Create logger for application
	l := logger.New(os.Stdout, os.Stderr)

	// Create game
	g, err := game.NewGame(&game.Options{
		Title:        "Clarity",
		RenderWidth:  480,
		RenderHeight: 270,
		Logger:       l,
	})
	if err != nil {
		l.Error.Panicf("creating game: %s", err)
	}

	// Run the game
	err = g.Run()
	if err != nil {
		l.Error.Panicf("running game: %s", err)
	}
}
