package art

import (
	"image"
	"image/color"
	"io"

	"github.com/d47id/art-bot/colors"
)

func makeDaveBot(size, padding int, col string) []*rectangle {
	width := 2*padding + 8
	unit := size / width
	rects := make([]*rectangle, 0, 4)

	rects = append(rects,
		// left eye
		&rectangle{
			Height: unit,
			Width:  2 * unit,
			X:      (padding + 2) * unit,
			Y:      padding * unit,
			Color:  col,
		},
		// right eye
		&rectangle{
			Height: unit,
			Width:  2 * unit,
			X:      (padding + 6) * unit,
			Y:      padding * unit,
			Color:  col,
		},
		// mouth
		&rectangle{
			Height: unit,
			Width:  4 * unit,
			X:      (padding + 4) * unit,
			Y:      (padding + 2) * unit,
			Color:  col,
		},
		// beard
		&rectangle{
			Height: 4 * unit,
			Width:  8 * unit,
			X:      padding * unit,
			Y:      (padding + 4) * unit,
			Color:  col,
		},
	)

	return rects
}

func rasterizeRectangles(size int, rects []*rectangle, background color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			color := contains(x, y, rects)
			if color == nil {
				color = background
			}
			img.Set(x, y, color)
		}
	}

	return img
}

// contains returns the color of the rectangle containing point (x, y) if it exists
// if no rectangle contains the point, contains returns nil
func contains(x, y int, rects []*rectangle) color.Color {
	for _, rect := range rects {
		if x >= rect.X && y >= rect.Y && x < rect.X+rect.Width && y < rect.Y+rect.Height {
			return colors.Color(rect.Color)
		}
	}
	return nil
}

// WriteDaveBotSVG writes an svg+xml encoded dave bot to w
func (b *Bot) WriteDaveBotSVG(size, padding int, foreground string, w io.Writer) error {
	rects := makeDaveBot(size, padding, foreground)
	data := struct {
		Size       int
		Rectangles []*rectangle
	}{size, rects}

	return b.tpl.ExecuteTemplate(w, "dave-bot.svg", data)
}

// GetDaveBot returns a rasterized dave bot
func (b *Bot) GetDaveBot(size, padding int, foreground, background string) image.Image {
	rects := makeDaveBot(size, padding, foreground)
	return rasterizeRectangles(size, rects, colors.Color(background))
}
