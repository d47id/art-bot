package server

import (
	"compress/gzip"

	"github.com/NYTimes/gziphandler"
	"github.com/d47id/zapmw"
)

func (s *Server) routes() {
	s.mux.Use(zapmw.New(s.l))
	s.mux.Use(gziphandler.MustNewGzipLevelHandler(gzip.BestCompression))
	s.mux.Use(clientHintsMiddleware)

	s.mux.Get("/", s.index)
	s.mux.Get("/resume", s.resume)
	s.mux.Get("/styles.css", s.styles)
	s.mux.Get("/theme.js", s.theme)
	s.mux.Get("/checkerboard.svg", s.checkerboard)
	s.mux.Get("/circles.svg", s.circles)
	s.mux.Get("/dave-bot.svg", s.daveBotSVG)
	s.mux.Get("/dave-bot.png", s.daveBotPNG)
	s.mux.Get("/pixellated.svg", s.pixellated)
	s.mux.Get("/bubbles.svg", s.bubbles)
}
