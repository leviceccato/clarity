package game

// Position component
type positionComponent struct {
	X, Y float64
	Z    int
}

func (positionComponent) Name() string {
	return "position"
}
