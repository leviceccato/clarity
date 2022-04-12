package asset

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	_ "image/png"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Embed assets folder inside the executable
//go:embed **/*
var FS embed.FS

// Simple icon loading util
func LoadIcon(path string) (image.Image, error) {
	iconBytes, err := FS.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading icon file: %w", err)
	}

	icon, _, err := image.Decode(bytes.NewReader(iconBytes))
	if err != nil {
		return nil, fmt.Errorf("decoding icon file: %w", err)
	}

	return icon, nil
}

func LoadFont(path string, size, DPI float64, hinting font.Hinting) (font.Face, error) {
	fontBytes, err := FS.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("loading font file: %w", err)
	}

	parsedFont, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("parsing font file: %w", err)
	}

	fontFace, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
		Size:    size,
		DPI:     DPI,
		Hinting: hinting,
	})
	if err != nil {
		return nil, fmt.Errorf("creating font face: %w", err)
	}

	return fontFace, nil
}
