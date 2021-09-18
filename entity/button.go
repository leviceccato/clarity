package entity

import (
	"fmt"
	"image/color"
	"math"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/leviceccato/clarity/component"
	"golang.org/x/image/font"
)

type ButtonEntityOptions struct {
	X, Y, Width, Height, Padding float64
	Text, Image, Animation       string
	Color                        color.NRGBA
	Font                         font.Face
}

const space = " "

func NewButtonEntity(options *ButtonEntityOptions) (*Entity, error) {
	e := NewEntity()
	e.Position = &component.PositionComponent{X: options.X, Y: options.Y}
	e.Size = &component.SizeComponent{Width: options.Width, Height: options.Height}
	appearance, err := component.NewAppearanceComponent(options.Image, options.Animation)
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %s", err)
	}
	e.Appearance = appearance

	// Button with text
	if options.Text != "" {
		// Fit text within padded rectangle and build a slice
		// of strings with unbroken words
		maxWidth := options.Width - (options.Padding * 2)
		textRect := text.BoundString(options.Font, options.Text)
		textWidth := textRect.Max.X - textRect.Min.X
		ratio := float64(textWidth) / maxWidth
		// Base max chars on average chars per line
		maxChars := int(math.Ceil(float64(len(options.Text)) / ratio))
		words := strings.Fields(options.Text)
		var (
			lines []string
			line  string
		)
		for i, word := range words {
			// We've reached the end of the line, start a new one
			if len(line+space+word) >= maxChars {
				lines = append(lines, line)
				line = ""
			}
			// Put space between every word
			if i > 0 {
				line += space
			}
			line += word
		}
		lines = append(lines, line)

		e.Text = &component.TextComponent{
			Lines:      lines,
			Color:      options.Color,
			Font:       options.Font,
			LineHeight: options.Font.Metrics().Height.Round(),
		}
	}
	return e, nil
}
