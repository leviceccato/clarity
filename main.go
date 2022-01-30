package main

import (
	"fmt"

	"github.com/leviceccato/clarity/game"
	"github.com/leviceccato/clarity/world"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := game.NewGame(game.Options{
		Title:        "Clarity",
		RenderWidth:  480,
		RenderHeight: 270,
	})
	if err != nil {
		panic(fmt.Sprintf("creating game: %s", err))
	}

	g.LoadWorld(world.NewStartWorld(g))
	g.LoadWorld(world.NewTitleWorld(g))
	g.ActivateWorlds("start")

	err = ebiten.RunGame(g)
	if err != nil {
		panic(fmt.Sprintf("running game: %s", err))
	}
}
