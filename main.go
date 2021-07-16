package main

import (
	"fmt"

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
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {

}

func (g *game) Layout(w, h int) (int, int) {
	return windowWidth, windowHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Clarity")
	mainState := &state{}
	mainState.loadWorld()
	err := ebiten.RunGame(&game{
		state: &state{},
	})
	if err != nil {
		fmt.Println(err)
	}
}
