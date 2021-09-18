package utility

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func LoadFonts(f map[string]string) (map[string]*font.Face, error) {
	fonts := map[string]*font.Face{}
	for name, path := range f {
		bytes, err := ioutil.ReadFile(path)
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
