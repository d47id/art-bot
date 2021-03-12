package art

import (
	"image"
	"io"
	"math/rand"

	"github.com/d47id/art-bot/colors"
)

type circle struct {
	X      int
	Y      int
	Radius int
	Color  string
}

type circleVals struct {
	Width   int
	Height  int
	Circles []circle
}

func makeGaussianCircles() []circle {
	count := rand.Intn(100)
	circles := make([]circle, 0, count)
	xStdDev := rand.Intn(50)
	yStdDev := rand.Intn(50)
	maxSize := rand.Intn(45) + 5
	sample := func(stdDev int) int {
		return int(rand.NormFloat64()*float64(stdDev) + 50)
	}

	for i := 0; i < count; i++ {
		circles = append(circles, circle{
			X:      sample(xStdDev),
			Y:      sample(yStdDev),
			Radius: rand.Intn(maxSize),
			Color:  colors.Random(),
		})
	}

	return circles
}

// WriteCircles writes gaussian circles svg+xml to w
func (b *Bot) WriteCircles(w io.Writer) error {
	return b.tpl.ExecuteTemplate(w, "circles.svg", circleVals{
		Width:   100,
		Height:  100,
		Circles: makeGaussianCircles(),
	})
}

func makeBubbleImage(img image.Image) (width int, height int, circles []circle) {
	width = 50

	// scale the image down to the svg pixel width
	var (
		origW  = img.Bounds().Max.X
		origH  = img.Bounds().Max.Y
		scale  = float64(origW) / float64(width)
		minRad = 2
		maxRad = 10
		rad    = func() int {
			return rand.Intn(maxRad-minRad) + minRad
		}
	)

	height = int(float64(origH) / scale)
	circles = make([]circle, 0, width*height)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := img.At(int(float64(x)*scale), int(float64(y)*scale))
			circles = append(circles, circle{
				X:      x*10 + 5, // center circle in 10x10 square in output svg
				Y:      y*10 + 5,
				Radius: rad(),
				Color:  cssRGBA(c),
			})
		}
	}

	// multiply width and height by 10 to create 10x10 squares for circles
	return width * 10, height * 10, circles
}

// WriteBubbleImage writes a random bubble image from the cache to w
func (b *Bot) WriteBubbleImage(w io.Writer) error {
	img, err := b.imgs.get()
	if err != nil {
		return err
	}

	width, height, circles := makeBubbleImage(img)
	return b.tpl.ExecuteTemplate(w, "circles.svg", circleVals{
		Width:   width,
		Height:  height,
		Circles: circles,
	})
}
