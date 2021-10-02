package asset

import (
	"embed"
)

// Embed assets folder inside the executable
//go:embed **/*
var FS embed.FS

// Global config
const (
	ConfTitle       = "Clarity"
	ConfDescription = "A 2D, story-driven, pixel art platformer."
	ConfVersion     = "0.1.0"

	ConfRenderWidth  = 480
	ConfRenderHeight = 270
)
