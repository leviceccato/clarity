package component

type HoverComponent struct {
	IsHovered bool
}

func NewHoverComponent() *HoverComponent {
	return &HoverComponent{}
}
