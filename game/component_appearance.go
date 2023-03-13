package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"strings"

	"github.com/leviceccato/clarity/asset"

	"github.com/hajimehoshi/ebiten/v2"
)

// Appearance component
type appearanceComponent struct {
	Image                      *ebiten.Image
	Frame, Duration            int
	Sequence, PreviousSequence string
	Time                       float64
	Frames                     []*image.Rectangle
	Sequences                  map[string]*appearanceSequence
}

func (appearanceComponent) Name() string {
	return "appearance"
}

// For use in appearance component
type appearanceSequence struct {
	From, To   int
	ShouldLoop bool
	Direction  string
}

// Load files, format the data and then create an Appearance component
// Pass an empty string as the second param if an animation isn't required
func newAppearanceComponent(imagePath, animationPath string) (*appearanceComponent, error) {
	c := &appearanceComponent{
		Sequences: map[string]*appearanceSequence{},
	}

	// Load image
	imgBytes, err := asset.FS.ReadFile(imagePath)
	if err != nil {
		return c, fmt.Errorf("loading appearance image: %s", err)
	}
	rawImg, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return c, fmt.Errorf("decoding appearance image: %s", err)
	}
	c.Image = ebiten.NewImageFromImage(rawImg)

	// No animation required
	if animationPath == "" {
		w, h := c.Image.Size()
		c.Frames = []*image.Rectangle{{
			Min: image.Point{0, 0},
			Max: image.Point{w, h},
		}}
		return c, nil
	}

	// Animation required
	anim, err := newAnimationFromFile(animationPath)
	if err != nil {
		return c, fmt.Errorf("loading appearance animation: %s", err)
	}
	for _, f := range anim.Frames {
		c.Duration += f.Duration
		rect := image.Rect(
			f.Frame.X, f.Frame.Y,
			f.Frame.X+f.Frame.W, f.Frame.Y+f.Frame.H,
		)
		c.Frames = append(c.Frames, &rect)
	}
	for _, t := range anim.Meta.FrameTags {
		// Treat tag name ending in '_loop' as meaning it should loop
		// There's no native way to do this in Aseprite
		nameSections := strings.Split(t.Name, "_")
		shouldLoop := nameSections[len(nameSections)-1] == "loop"
		name := nameSections[0]
		c.Sequences[name] = &appearanceSequence{
			From:       t.From,
			To:         t.To,
			ShouldLoop: shouldLoop,
			Direction:  t.Direction,
		}
		// Set first sequence as the default
		if c.Sequence == "" {
			c.Sequence = name
		}
	}

	return c, nil
}

// Struct to match Aseprite JSON output
type animation struct {
	Frames []struct {
		Duration int `json:"duration"`
		Frame    struct {
			X int `json:"x"`
			Y int `json:"y"`
			W int `json:"w"`
			H int `json:"h"`
		} `json:"frame"`
	} `json:"frames"`
	Meta struct {
		Size struct {
			W int `json:"w"`
			H int `json:"h"`
		} `json:"size"`
		FrameTags []struct {
			Name      string `json:"name"`
			From      int    `json:"from"`
			To        int    `json:"to"`
			Direction string `json:"direction"`
		} `json:"frameTags"`
	} `json:"meta"`
}

// Load an Aseprite JSON file as an animation
func newAnimationFromFile(path string) (*animation, error) {
	file, err := asset.FS.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("loading animation json: %s", err)
	}
	var anim animation
	err = json.Unmarshal(file, &anim)
	if err != nil {
		return &anim, fmt.Errorf("unmarshalling animation json: %s", err)
	}
	return &anim, nil
}
