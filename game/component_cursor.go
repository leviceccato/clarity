package game

// Cursor component
type cursorComponent struct {
	isHovering bool
}

func (cursorComponent) Name() string {
	return "cursor"
}
