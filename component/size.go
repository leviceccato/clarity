package component

type Size struct {
	Width, Height float64
}

func NewSize(width, height float64) *Size {
	return &Size{
		Width:  width,
		Height: height,
	}
}
