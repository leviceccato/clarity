package game

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/leviceccato/clarity/engine"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/exp/slices"
)

func newDrawSystem(g *Game) *engine.System {
	s := engine.NewSystem("draw", []string{
		"appearance",
		"position",
		"size",
	})

	s.Draw = func(screen *ebiten.Image) {
		screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})

		// Sort entity ids by Z position
		zSortedEntities := make([]*engine.Entity, 0, len(s.EntityIds))
		for _, id := range s.EntityIds {
			zSortedEntities = append(zSortedEntities, g.GetEntity(id))
		}
		slices.SortFunc(zSortedEntities, func(a, b *engine.Entity) bool {
			aPosition, _ := engine.GetComponent(a, &positionComponent{})
			bPosition, _ := engine.GetComponent(b, &positionComponent{})

			return aPosition.Z < bPosition.Z
		})

		for _, e := range zSortedEntities {
			appearance, _ := engine.GetComponent(e, &appearanceComponent{})
			size, _ := engine.GetComponent(e, &sizeComponent{})
			position, _ := engine.GetComponent(e, &positionComponent{})
			text, hasText := engine.GetComponent(e, &textComponent{})

			// Draw image
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(position.X, position.Y)

			screen.DrawImage(
				appearance.Image.SubImage(*appearance.Frames[appearance.Frame]).(*ebiten.Image),
				options,
			)

			// No text? nothing to do
			if !hasText {
				continue
			}

			// Draw lines of text
			for i, line := range text.Lines {
				x := line.X + text.Padding
				y := (size.Height / 2) - float64((text.LineHeight*len(text.Lines))/2) + float64((i+1)*text.LineHeight)

				// Text transform
				c := line.Content
				switch text.TextTransform {
				case textComponentTransformUppercase:
					c = strings.ToUpper(c)
				case textComponentTransformLowercase:
					c = strings.ToLower(c)
				case textComponentTransformTitlecase:
					c = strings.ToTitle(c)
				}

				ebitentext.Draw(screen, c, text.Font, int(position.X+x), int(position.Y+y), text.Color)
			}
		}

		if g.isDebug {
			ebitenutil.DebugPrint(screen, fmt.Sprintf(
				"TPS: %0.2f\nFPS: %0.2f",
				ebiten.CurrentTPS(),
				ebiten.CurrentFPS(),
			))
		}
	}

	return s
}
