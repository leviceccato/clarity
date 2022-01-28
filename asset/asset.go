package asset

import (
	"bytes"
	"embed"
	"fmt"
	"image"

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

// Load fonts into map with names as keys
func LoadFonts(f map[string]string) (map[string]*font.Face, error) {
	fonts := map[string]*font.Face{}

	for name, path := range f {
		bytes, err := FS.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("loading '%s' font file: %s", name, err)
		}

		fontfile, err := opentype.Parse(bytes)
		if err != nil {
			return nil, fmt.Errorf("parsing '%s' font file: %s", name, err)
		}

		face, err := opentype.NewFace(fontfile, &opentype.FaceOptions{
			Size:    11,
			DPI:     72,
			Hinting: font.HintingNone,
		})
		if err != nil {
			return nil, fmt.Errorf("creating '%s' font face: %s", name, err)
		}

		fonts[name] = &face
	}

	return fonts, nil
}
