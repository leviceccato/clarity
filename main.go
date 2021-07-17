package main

import (
	"fmt"

	"github.com/leviceccato/clarity/world"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth  = 320
	windowHeight = 240
)

type game struct {
	state *state
}

func (g *game) Update() error {
	g.state.update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.state.draw()
}

func (g *game) Layout(_, _ int) (int, int) {
	return windowWidth, windowHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Clarity")
	mainState := newState()
	mainState.loadWorld(world.NewStartWorld())
	err := ebiten.RunGame(&game{
		state: &state{},
	})
	if err != nil {
		fmt.Println(err)
	}
}
