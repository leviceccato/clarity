package component

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Appearance struct {
	Image *ebiten.Image
}

func NewAppearance() *Appearance {
	return &Appearance{}
}
