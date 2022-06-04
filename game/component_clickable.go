package game

// Clickable component
type clickableComponent struct {
	action func()
}

func (_ clickableComponent) Name() string {
	return "clickable"
}