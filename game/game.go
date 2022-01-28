package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type world interface {
	Update()
	Draw()
}

type game struct {
	renderWidth, renderHeight int

	activeWorld string
	worlds      map[string]world
}

func (g *game) Update() error {
	g.worlds[g.activeWorld].Update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.worlds[g.activeWorld].Draw()
}

func (g *game) Layout(w, h int) (int, int) {
	return g.renderWidth, g.renderHeight
}
