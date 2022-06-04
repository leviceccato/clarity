package game

import (
	"fmt"
	"image/color"
	"math"
	"strings"

	"github.com/leviceccato/clarity/engine"

	ebitentext "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type buttonEntityOptions struct {
	x, y, width, height, padding float64
	z                            int
	isCentered                   bool
	text, image, animation       string
	textTransform                textComponentTransform
	color                        color.NRGBA
	font                         font.Face
	action                       func()
}

func newButtonEntity(g *Game, options *buttonEntityOptions) (*engine.Entity, error) {
	e := g.NewEntity()

	e.AddComponent(&hoverableComponent{})

	e.AddComponent(&clickableComponent{action: options.action})

	e.AddComponent(&positionComponent{
		X: options.x,
		Y: options.y,
		Z: options.z,
	})

	e.AddComponent(&sizeComponent{
		Width:  options.width,
		Height: options.height,
	})

	buttonAppearance, err := newAppearanceComponent(options.image, options.animation)
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %w", err)
	}
	e.AddComponent(buttonAppearance)

	// Button w/o text
	if options.text == "" {
		return e, nil
	}

	// Calculate text lines
	maxWidth := options.width - (options.padding * 2)
	textWidth := getTextWidth(options.font, options.text)
	ratio := float64(textWidth) / maxWidth
	// Base max chars on average chars per line
	maxChars := int(math.Ceil(float64(len(options.text)) / ratio))
	words := strings.Fields(options.text)
	wordCount := len(words)

	// Fit text within padded rectangle and build a slice
	// of strings with unbroken words
	var lines []textLine
	var line textLine
	var word string
	var lineWidth int

	// Loop over words and then once more to add the final line
	for i := 0; i <= wordCount; i++ {
		if i < wordCount {
			word = words[i]
		}

		// If we are on the last iteration or if we will go on to the next
		// line then add a new line
		if i == wordCount || len(line.Content+" "+word) >= maxChars {
			if options.isCentered {
				lineWidth = getTextWidth(options.font, line.Content)
				line.X = (maxWidth - float64(lineWidth)) / 2
			}
			lines = append(lines, line)
			line = textLine{}
		}

		// Put space between every word but not at the start of a line
		if line.Content != "" {
			line.Content += " "
		}

		line.Content += word
	}

	// Button w text
	e.AddComponent(&textComponent{
		Lines:         lines,
		Color:         options.color,
		TextTransform: options.textTransform,
		Font:          options.font,
		Padding:       options.padding,
		LineHeight:    options.font.Metrics().CapHeight.Round(),
	})

	return e, nil
}

func getTextWidth(fontFace font.Face, content string) int {
	rect := ebitentext.BoundString(fontFace, content)

	return rect.Max.X - rect.Min.X
}
