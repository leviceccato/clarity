package game

type state struct {
	isCursorHovering bool

	// Controls         map[component.Control]*component.InputData
	// MouseInputs      map[ebiten.MouseButton]component.Control
	// KeyInputs        map[ebiten.Key]component.Control
}

func (s *state) setIsCursorHovering(to bool) {
	s.isCursorHovering = to
}
