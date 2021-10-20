// Simple package containing global project config/variables
package config

import (
	"image/color"
)

const (
	Title       = "Clarity"
	Description = "A 2D, story-driven, pixel art platformer."
	Version     = "0.1.0"

	RenderWidth  = 480
	RenderHeight = 270
)

var (
	ColFgTitle = color.NRGBA{255, 240, 157, 255}
)
