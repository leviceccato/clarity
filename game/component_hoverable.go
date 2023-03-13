package game

// Hover component
type hoverableComponent struct {
	IsHovered bool
}

func (hoverableComponent) Name() string {
	return "hoverable"
}
