package main

import (
	"os"

	"github.com/leviceccato/clarity/game"
	"github.com/leviceccato/clarity/logger"
)

func main() {
	// Create logger for application
	l := logger.New(os.Stdout, os.Stderr)

	// Start game
	err := game.CreateAndRun(&game.Options{
		Title:        "Clarity - DEBUG",
		RenderWidth:  480,
		RenderHeight: 270,
		Logger:       l,
	})
	if err != nil {
		l.Error.Panicf("starting game: %s", err)
	}
}
