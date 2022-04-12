package game

// Cursor component
type cursorComponent struct {
	isHovering bool
}

func (_ cursorComponent) Name() string {
	return "cursor"
}
