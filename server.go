package main

import (
	"html/template"
	"math/rand"
	"net/http"

	"github.com/d47id/zapmw"
	"go.uber.org/zap"
)

type server struct {
	l   *zap.Logger
	tpl *template.Template
}

func (s *server) index(w http.ResponseWriter, r *http.Request) {
	backgrounds := randstr([]string{"checkerboard.svg", "circles.svg"})
	data := struct {
		Vignette   string
		Text       string
		Background string
	}{
		colors.Sample(),
		colors.Sample(),
		backgrounds.Sample(),
	}

	if err := s.tpl.ExecuteTemplate(w, "index.html", data); err != nil {
		zapmw.Extract(r.Context()).Error("execute template", zap.Error(err))
	}
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
				Color:  colors.Sample(),
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
					Color:  colors.Sample(),
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
			Color:  colors.Sample(),
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
