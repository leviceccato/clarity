package game

import (
	"github.com/leviceccato/clarity/engine"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func newInputSystem(g *Game) *engine.System {
	s := engine.NewSystem("input", []string{})

	s.Update = func() error {
		x, y := engine.CursorPosition()

		for cmd, inputs := range g.InputBindings {
			for _, input := range inputs {
				switch i := input.(type) {
				case ebiten.MouseButton:
					if inpututil.IsMouseButtonJustPressed(i) {
						g.setInput(cmd, &inputData{
							isStart: true,
							x:       x,
							y:       y,
						})
						continue
					}

					if inpututil.IsMouseButtonJustReleased(i) {
						g.setInput(cmd, &inputData{
							isStart: false,
							x:       x,
							y:       y,
						})
						continue
					}

					if !ebiten.IsMouseButtonPressed(i) {
						g.setInput(cmd, nil)
					}
				case ebiten.Key:
					if inpututil.IsKeyJustPressed(i) {
						g.setInput(cmd, &inputData{isStart: true})
						continue
					}

					if inpututil.IsKeyJustReleased(i) {
						g.setInput(cmd, &inputData{isStart: false})
						continue
					}

					if !ebiten.IsKeyPressed(i) {
						g.setInput(cmd, nil)
					}
				}
			}
		}

		return nil
	}

	return s
}
