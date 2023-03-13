package game

// Clickable component
type clickableComponent struct {
	action func()
}

func (clickableComponent) Name() string {
	return "clickable"
}
