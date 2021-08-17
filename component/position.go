package component

type PositionComponent struct {
	X, Y float64
}

func NewPositionComponent(x, y float64) *PositionComponent {
	return &PositionComponent{
		X: x,
		Y: y,
	}
}
