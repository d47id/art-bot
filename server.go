package main

import (
	"html/template"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/d47id/art-bot/colors"
	"github.com/d47id/zapmw"

	"go.uber.org/zap"
)

type server struct {
	l   *zap.Logger
	tpl *template.Template
}

func (s *server) index(w http.ResponseWriter, r *http.Request) {
	backgrounds := []string{"checkerboard.svg", "circles.svg"}
	data := struct {
		Vignette   string
		Text       string
		Background string
	}{
		colors.Random(),
		colors.Random(),
		backgrounds[rand.Intn(len(backgrounds))],
	}

	if err := s.tpl.ExecuteTemplate(w, "index.html", data); err != nil {
		zapmw.Extract(r.Context()).Error("execute template", zap.Error(err))
	}
}

func (s *server) daveBotSVG(w http.ResponseWriter, r *http.Request) {
	size, foreground, _ := parseDaveBotParameters(r)

	rects := makeDaveBot(size, foreground)
	data := struct {
		Size       int
		Rectangles []*rectangle
	}{size, rects}

	w.Header().Add("Content-Type", "image/svg+xml")
	if err := s.tpl.ExecuteTemplate(w, "dave-bot.svg", data); err != nil {
		zapmw.Extract(r.Context()).Error("execute template", zap.Error(err))
	}
}

func (s *server) daveBotPNG(w http.ResponseWriter, r *http.Request) {
	size, foreground, background := parseDaveBotParameters(r)

	rects := makeDaveBot(size, foreground)
	img := rasterizeRectangles(size, rects, colors.Color(background))

	w.Header().Add("Content-Type", "image/png")
	if err := png.Encode(w, img); err != nil {
		zapmw.Extract(r.Context()).Error("encode png", zap.Error(err))
	}
}

func parseDaveBotParameters(r *http.Request) (size int, foreground, background string) {
	size = 600
	q := r.URL.Query()
	if i, err := strconv.Atoi(q.Get("size")); err == nil {
		// ensure size is divisible by 12
		if i%12 != 0 {
			i += 12 - (i % 12)
		}
		size = i
	}
	foreground = colors.DeepPink
	if col := q.Get("foreground"); colors.Color(col) != nil {
		foreground = col
	}
	background = colors.Honeydew
	if col := q.Get("background"); colors.Color(col) != nil {
		background = col
	}

	return
}

func makeDaveBot(size int, col string) []*rectangle {
	unit := size / 12
	rects := make([]*rectangle, 0, 4)

	// left eye
	rects = append(rects, &rectangle{
		Height: unit,
		Width:  2 * unit,
		X:      4 * unit,
		Y:      2 * unit,
		Color:  col,
	})

	// right eye
	rects = append(rects, &rectangle{
		Height: unit,
		Width:  2 * unit,
		X:      8 * unit,
		Y:      2 * unit,
		Color:  col,
	})

	// mouth
	rects = append(rects, &rectangle{
		Height: unit,
		Width:  4 * unit,
		X:      6 * unit,
		Y:      4 * unit,
		Color:  col,
	})

	// beard
	rects = append(rects, &rectangle{
		Height: 4 * unit,
		Width:  8 * unit,
		X:      2 * unit,
		Y:      6 * unit,
		Color:  col,
	})

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

func (s *server) checkerboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/svg+xml")
	if err := s.tpl.ExecuteTemplate(w, "checkerboard.svg", makeErodedCheckerboard()); err != nil {
		zapmw.Extract(r.Context()).Error("execute template", zap.Error(err))
	}
}

func makeRandomCheckerboard() []*rectangle {
	rects := make([]*rectangle, 0, 100)

	for x := 0; x < 100; x += 10 {
		for y := 0; y < 100; y += 10 {
			rects = append(rects, &rectangle{
				X:      x,
				Y:      y,
				Width:  10,
				Height: 10,
				Color:  colors.Random(),
			})
		}
	}

	return rects
}

func makeErodedCheckerboard() []*rectangle {
	rects := make([]*rectangle, 0, 100)
	limit := rand.Intn(100)
	gate := func() bool {
		return rand.Intn(100) < limit
	}

	for x := 0; x < 100; x += 10 {
		for y := 0; y < 100; y += 10 {
			if gate() {
				rects = append(rects, &rectangle{
					X:      x,
					Y:      y,
					Width:  10,
					Height: 10,
					Color:  colors.Random(),
				})
			}
		}
	}

	return rects
}

type rectangle struct {
	X      int
	Y      int
	Height int
	Width  int
	Color  string
}

func (s *server) circles(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/svg+xml")
	if err := s.tpl.ExecuteTemplate(w, "circles.svg", makeGaussianCircles()); err != nil {
		zapmw.Extract(r.Context()).Error("execute template", zap.Error(err))
	}
}

func makeGaussianCircles() []*circle {
	count := rand.Intn(100)
	circles := make([]*circle, 0, count)
	xStdDev := rand.Intn(50)
	yStdDev := rand.Intn(50)
	maxSize := rand.Intn(45) + 5
	sample := func(stdDev int) int {
		return int(rand.NormFloat64()*float64(stdDev) + 50)
	}

	for i := 0; i < count; i++ {
		circles = append(circles, &circle{
			X:      sample(xStdDev),
			Y:      sample(yStdDev),
			Radius: rand.Intn(maxSize),
			Color:  colors.Random(),
		})
	}

	return circles
}

type circle struct {
	X      int
	Y      int
	Radius int
	Color  string
}
