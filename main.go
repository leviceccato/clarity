package main

import (
	"errors"
	"flag"
	"os"

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
	if errors.Is(err, game.CloseError) {
		l.Info.Printf("window closed")
		return
	}
	if err != nil {
		l.Error.Panic(err)
	}
}
