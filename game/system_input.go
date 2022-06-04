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
			func() {
				for _, input := range inputs {
					switch i := input.(type) {
					case ebiten.MouseButton:
						if inpututil.IsMouseButtonJustPressed(i) {
							g.setInput(cmd, &inputData{
								progress: inputStart,
								x:        x,
								y:        y,
							})
							return
						}

						if inpututil.IsMouseButtonJustReleased(i) {
							g.setInput(cmd, &inputData{
								progress: inputEnd,
								x:        x,
								y:        y,
							})
							return
						}
					case ebiten.Key:
						if inpututil.IsKeyJustPressed(i) {
							g.setInput(cmd, &inputData{progress: inputStart})
							return
						}

						if inpututil.IsKeyJustReleased(i) {
							g.setInput(cmd, &inputData{
								progress: inputEnd,
							})
							return
						}
					}
				}

				data := g.inputs[cmd]
				if data == nil {
					return
				}

				switch data.progress {
				case inputStart:
					newData := *data
					newData.progress = inputMiddle

					g.setInput(cmd, &newData)
				case inputEnd:
					g.setInput(cmd, nil)
				}
			}()
		}

		return nil
	}

	return s
}
