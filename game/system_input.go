package game

import (
	"github.com/leviceccato/clarity/engine"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func newInputSystem(g *Game) *engine.System {
	s := engine.NewSystem("input", []string{})

	s.Update = func() error {
		for input, cmd := range g.mouseInputs {
			x, y := engine.CursorPosition()

			if inpututil.IsMouseButtonJustPressed(input) {
				g.setInput(cmd, &inputData{
					isStart: true,
					x:       x,
					y:       y,
				})
				continue
			}

			if inpututil.IsMouseButtonJustReleased(input) {
				g.setInput(cmd, &inputData{
					isStart: false,
					x:       x,
					y:       y,
				})
				continue
			}

			if !ebiten.IsMouseButtonPressed(input) {
				g.setInput(cmd, nil)
			}
		}

		for input, cmd := range g.keyInputs {
			if inpututil.IsKeyJustPressed(input) {
				g.setInput(cmd, &inputData{isStart: true})
				continue
			}

			if inpututil.IsKeyJustReleased(input) {
				g.setInput(cmd, &inputData{isStart: false})
				continue
			}

			if !ebiten.IsKeyPressed(input) {
				g.setInput(cmd, nil)
			}
		}

		return nil
	}

	return s
}
