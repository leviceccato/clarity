package game

// Position component
type positionComponent struct {
	X, Y float64
}

func (_ positionComponent) Name() string {
	return "position"
}
