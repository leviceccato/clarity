package util

import (
	"bytes"
	"fmt"
	"image"

	"github.com/leviceccato/clarity/asset"
)

// Simple icon loading util
func LoadIcon(path string) (image.Image, error) {
	iconBytes, err := asset.FS.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading icon file: %s", err)
	}
	icon, _, err := image.Decode(bytes.NewReader(iconBytes))
	if err != nil {
		return nil, fmt.Errorf("decoding icon file: %s", err)
	}
	return icon, nil
}
