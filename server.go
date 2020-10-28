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
	data := struct {
		Vignette string
		Text     string
	}{
		colors[rand.Intn(len(colors))],
		colors[rand.Intn(len(colors))],
	}

	if err := s.tpl.ExecuteTemplate(w, "index.html", data); err != nil {
		zapmw.Extract(r.Context()).Error("execute template", zap.Error(err))
	}
}

func (s *server) checkerboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/svg+xml")
	if err := s.tpl.ExecuteTemplate(w, "checkerboard.svg", makeRandomCheckerboard()); err != nil {
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
				Color:  colors[rand.Intn(len(colors))],
			})
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
