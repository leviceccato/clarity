package game

import (
	"github.com/leviceccato/clarity/engine"
)

func newInputSystem(g *Game) *engine.System {
	s := engine.NewSystem("input", []string{})

	s.Update = func() error {
		// for _, entityId := range s.EntityIds {
		// 	e := g.GetEntity(entityId)
		// 	// for mouseInput, control := range s.MouseInputs {
		// 	// 	if inpututil.IsMouseButtonJustPressed(mouseInput) {
		// 	// 		x, y := ebiten.CursorPosition()
		// 	// 		s.SetControl(control, &component.InputData{X: x, Y: y})
		// 	// 	}
		// 	// 	if inpututil.IsMouseButtonJustReleased(mouseInput) {
		// 	// 		s.SetControl(control, nil)
		// 	// 	}
		// 	// }

		// 	// for keyInput, control := range s.KeyInputs {
		// 	// 	if inpututil.IsKeyJustPressed(keyInput) {
		// 	// 		s.SetControl(control, &component.InputData{})
		// 	// 	}
		// 	// 	if inpututil.IsKeyJustReleased(keyInput) {
		// 	// 		s.SetControl(control, nil)
		// 	// 	}
		// 	// }
		// }

		return nil
	}

	return s
}
