package game

import (
	"github.com/leviceccato/clarity/engine"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func newInputSystem(g *Game) *engine.System {
	s := engine.NewSystem("input", []string{})

	s.Update = func() error {
		for cmd, inputs := range g.InputBindings {
			func() {
				for _, input := range inputs {
					switch i := input.(type) {
					case ebiten.MouseButton:
						if inpututil.IsMouseButtonJustPressed(i) {
							x, y := engine.CursorPosition()
							g.setInput(cmd, &inputData{
								isStart: true,
								x:       x,
								y:       y,
							})
							return
						}

						if inpututil.IsMouseButtonJustReleased(i) {
							x, y := engine.CursorPosition()
							g.setInput(cmd, &inputData{
								isStart:      false,
								shouldExpire: true,
								x:            x,
								y:            y,
							})
							return
						}
					case ebiten.Key:
						if inpututil.IsKeyJustPressed(i) {
							g.setInput(cmd, &inputData{isStart: true})
							return
						}

						if inpututil.IsKeyJustReleased(i) {
							g.setInput(cmd, &inputData{
								isStart:      false,
								shouldExpire: true,
							})
							return
						}
					}

					data := g.inputs[cmd]
					if data != nil && data.shouldExpire {
						g.setInput(cmd, nil)
						return
					}
				}
			}()
		}

		return nil
	}

	return s
}
