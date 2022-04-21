package main

import (
	"flag"
	"os"

	"github.com/leviceccato/clarity/game"
	"github.com/leviceccato/clarity/logger"
)

func main() {
	// Allow setting debug mode with flags
	isDebug := flag.Bool("debug", false, "Set debug mode")
	flag.Parse()

	// Create logger for application
	l := logger.New(os.Stdout, os.Stderr)

	// Start game
	err := game.CreateAndRun(&game.Options{
		Title:        "Clarity - DEBUG",
		RenderWidth:  480,
		RenderHeight: 270,
		Logger:       l,
		IsDebug:      *isDebug,
	})
	if err != nil {
		l.Error.Panicf("starting game: %s", err)
	}
}
