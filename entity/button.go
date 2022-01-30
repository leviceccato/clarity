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
	IsCentered                   bool
	Text, Image, Animation       string
	Color                        color.NRGBA
	Font                         font.Face
}

func NewButtonEntity(options *ButtonEntityOptions) (*Entity, error) {
	e := NewEntity()
	e.Hover = &component.HoverComponent{}
	e.Position = &component.PositionComponent{X: options.X, Y: options.Y}
	e.Size = &component.SizeComponent{Width: options.Width, Height: options.Height}
	appearance, err := component.NewAppearanceComponent(options.Image, options.Animation)
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %s", err)
	}
	e.Appearance = appearance

	// No text
	if options.Text == "" {
		return e, nil
	}

	// Button with text
	maxWidth := options.Width - (options.Padding * 2)
	textWidth := getTextWidth(options.Font, options.Text)
	ratio := float64(textWidth) / maxWidth
	// Base max chars on average chars per line
	maxChars := int(math.Ceil(float64(len(options.Text)) / ratio))
	words := strings.Fields(options.Text)
	wordCount := len(words)
	// Fit text within padded rectangle and build a slice
	// of strings with unbroken words
	var (
		lines     []component.TextLine
		line      component.TextLine
		word      string
		lineWidth int
	)
	// Loop over words and then once more to add the final line
	for i := 0; i <= wordCount; i++ {
		if i < wordCount {
			word = words[i]
		}
		// If we are on the last iteration or if we will go on to the next
		// line then add a new line
		if i == wordCount || len(line.Content+" "+word) >= maxChars {
			if options.IsCentered {
				lineWidth = getTextWidth(options.Font, line.Content)
				line.X = (maxWidth - float64(lineWidth)) / 2
			}
			lines = append(lines, line)
			line = component.TextLine{}
		}
		// Put space between every word but not at the start of a line
		if line.Content != "" {
			line.Content += " "
		}
		line.Content += word
	}

	e.Text = &component.TextComponent{
		Lines:      lines,
		Color:      options.Color,
		Font:       options.Font,
		Padding:    options.Padding,
		LineHeight: options.Font.Metrics().Height.Round(),
	}
	return e, nil
}

func getTextWidth(f font.Face, content string) int {
	rect := text.BoundString(f, content)
	return rect.Max.X - rect.Min.X
}
