package game

// Clickable component
type clickableComponent struct {
	Action clickableAction
}

func (_ clickableComponent) Name() string {
	return "clickable"
}

type clickableAction struct {
	name string
	args []any
}
