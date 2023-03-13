package main

import (
	"errors"
	"flag"
	"os"

	"github.com/leviceccato/clarity/engine"
	"github.com/leviceccato/clarity/game"
	"github.com/leviceccato/clarity/logger"
)

func main() {
	isDebug := flag.Bool("debug", false, "Set debug mode")
	savePath := flag.String("save-path", ".", "Set location for game saves")
	flag.Parse()

	// Create logger for application
	l := logger.New(os.Stdout, os.Stderr)
	l.Info.Printf("starting game")

	// Start game
	err := game.CreateAndRun(&game.Options{
		Title:        "Clarity",
		RenderWidth:  480,
		RenderHeight: 270,
		Logger:       l,
		IsDebug:      *isDebug,
		SavePath:     *savePath,
	})

	// Must return an error in update to close game
	if errors.Is(err, engine.CloseError) {
		l.Info.Printf("window closed")
		return
	}

	// Generic error
	l.Error.Panic(err)
}
