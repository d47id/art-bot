package art

import (
	"context"
	"fmt"
	"image"
	"io"

	// image codecs
	_ "image/jpeg"
	_ "image/png"
)

type pixel struct {
	X     int
	Y     int
	Size  int
	Color string
}

type vector struct {
	Width  int
	Height int
	Pixels []pixel
}

// imageToVector converts an image.Image into a vector
func imageToVector(img image.Image) *vector {
	const (
		width  = 50
		pxSize = 10
	)

	// scale the image down to the svg pixel width
	var (
		origW  = img.Bounds().Max.X
		origH  = img.Bounds().Max.Y
		scale  = float64(origW) / float64(width)
		height = int(float64(origH) / scale)
	)

	// generate svg pixels by sampling source image
	v := &vector{
		Width:  width * pxSize, // scale width and height by pixel size
		Height: height * pxSize,
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := img.At(int(float64(x)*scale), int(float64(y)*scale))
			r, g, b, _ := c.RGBA()
			r, g, b = r/0x101, g/0x101, b/0x101 // convert 16bit rgb values to 8-bit
			v.Pixels = append(v.Pixels, pixel{
				X:     x * pxSize,
				Y:     y * pxSize,
				Size:  pxSize,
				Color: fmt.Sprintf("rgba(%d, %d, %d, 1)", r, g, b),
			})
		}
	}

	return v
}

// WritePixellatedSVG writes a random image rendered as svg+xml "pixels"
// to the given writer
func (b *Bot) WritePixellatedSVG(ctx context.Context, w io.Writer) error {
	img, err := b.imgs.GetImage()
	if err != nil {
		return err
	}

	return b.tpl.ExecuteTemplate(w, "pixels.svg", imageToVector(img))
}
