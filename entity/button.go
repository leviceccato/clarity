package entity

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/leviceccato/clarity/component"
	"github.com/leviceccato/clarity/utility"
	"golang.org/x/image/font"
)

type ButtonEntityOptions struct {
	X, Y, Width, Height, Padding float64
	Text, Image, Animation       string
	Color                        color.NRGBA
	Font                         font.Face
}

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
		maxWidth := options.Width - (options.Padding * 2)
		textRect := text.BoundString(options.Font, options.Text)
		textWidth := textRect.Max.X - textRect.Min.X
		ratio := float64(textWidth) / maxWidth
		maxChars := int(math.Ceil(float64(len(options.Text)) / ratio))
		lineCount := int(math.Ceil(ratio))
		var lines []string
		for i := 0; i < lineCount; i++ {
			lines = append(lines, utility.Substr(options.Text, i*maxChars, maxChars))
		}
		e.Text = &component.TextComponent{
			Lines:      lines,
			Color:      options.Color,
			Font:       options.Font,
			LineHeight: options.Font.Metrics().Height.Round(),
		}
	}
	return e, nil
}
