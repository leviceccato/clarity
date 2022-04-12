package game

// Hover component
type hoverableComponent struct {
	IsHovered bool
}

func (_ hoverableComponent) Name() string {
	return "hoverable"
}
