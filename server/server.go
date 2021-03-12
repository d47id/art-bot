package server

import (
	"embed"
	"html/template"
	"image/png"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/d47id/art-bot/art"
	"github.com/d47id/art-bot/colors"
	"github.com/d47id/zapmw"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

//go:embed www/*
var www embed.FS

// Server handles http requests
type Server struct {
	l   *zap.Logger
	tpl *template.Template
	bot *art.Bot
	mux *chi.Mux
}

// New builds a Server
func New(l *zap.Logger, bot *art.Bot) (*Server, error) {
	// parse templates
	tpl, err := template.ParseFS(www, "www/*")
	if err != nil {
		return nil, err
	}

	// setup routing
	s := &Server{l: l, tpl: tpl, bot: bot, mux: chi.NewRouter()}
	s.routes()

	return s, nil
}

// ServeHTTP implements http.Handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	backgrounds := []string{"checkerboard.svg", "circles.svg", "pixellated.svg", "bubbles.svg"}
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

func (s *Server) daveBotSVG(w http.ResponseWriter, r *http.Request) {
	size, padding, foreground, _ := parseDaveBotParameters(r)

	w.Header().Add("Content-Type", "image/svg+xml")
	if err := s.bot.WriteDaveBotSVG(size, padding, foreground, w); err != nil {
		zapmw.Extract(r.Context()).Error("execute template", zap.Error(err))
	}
}

func (s *Server) daveBotPNG(w http.ResponseWriter, r *http.Request) {
	img := s.bot.GetDaveBot(parseDaveBotParameters(r))

	w.Header().Add("Content-Type", "image/png")
	if err := png.Encode(w, img); err != nil {
		zapmw.Extract(r.Context()).Error("encode png", zap.Error(err))
	}
}

func parseDaveBotParameters(r *http.Request) (size, padding int, foreground, background string) {
	q := r.URL.Query()
	padding = 2
	if i, err := strconv.Atoi(q.Get("padding")); err == nil {
		padding = i
	}

	size = 600
	if i, err := strconv.Atoi(q.Get("size")); err == nil {
		size = i
	}

	// ensure size is divisible by minimum width
	width := 2*padding + 8
	if size%width != 0 {
		size += width - size%width
	}

	foreground = colors.Random()
	if col := q.Get("foreground"); colors.Color(col) != nil {
		foreground = col
	}
	background = colors.Random()
	if col := q.Get("background"); colors.Color(col) != nil {
		background = col
	}

	return
}

func (s *Server) checkerboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/svg+xml")
	if err := s.bot.WriteCheckerboard(w); err != nil {
		zapmw.Extract(r.Context()).Error("execute template", zap.Error(err))
	}
}

func (s *Server) circles(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/svg+xml")
	if err := s.bot.WriteCircles(w); err != nil {
		zapmw.Extract(r.Context()).Error("execute template", zap.Error(err))
	}
}

func (s *Server) pixellated(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/svg+xml")
	if err := s.bot.WritePixellatedSVG(w); err != nil {
		zapmw.Extract(r.Context()).Error("write pixellated svg", zap.Error(err))
	}
}

func (s *Server) bubbles(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/svg+xml")
	if err := s.bot.WriteBubbleImage(w); err != nil {
		zapmw.Extract(r.Context()).Error("write bubble image", zap.Error(err))
	}
}
