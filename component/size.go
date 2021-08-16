package component

type SizeComponent struct {
	Width, Height float64
}

func NewSize(width, height float64) *SizeComponent {
	return &SizeComponent{
		Width:  width,
		Height: height,
	}
}
