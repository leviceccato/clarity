package game

// Size component
type sizeComponent struct {
	Width, Height float64
}

func (_ sizeComponent) Name() string {
	return "size"
}
