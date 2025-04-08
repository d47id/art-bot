package server

import (
	"embed"
	"fmt"
	"html/template"
	"image/color"
	"image/png"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/d47id/art-bot/art"
	"github.com/d47id/art-bot/colors"

	"github.com/d47id/zapmw"
	"github.com/go-chi/chi/v5"
	colorful "github.com/lucasb-eyer/go-colorful"
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
	if err := s.tpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		zapmw.Extract(r.Context()).Error("execute template", zap.Error(err))
	}
}

func (s *Server) resume(w http.ResponseWriter, r *http.Request) {
	if err := s.tpl.ExecuteTemplate(w, "resume.html", nil); err != nil {
		zapmw.Extract(r.Context()).Error("execute template", zap.Error(err))
	}
}

func (s *Server) styles(w http.ResponseWriter, r *http.Request) {
	// Check for user's color scheme preference
	prefersDark := false

	// First try to use Client Hints header (Chrome, Edge, etc.)
	if r.Header.Get("Sec-CH-Prefers-Color-Scheme") == "dark" {
		prefersDark = true
	} else {
		// Fallback to cookie for browsers that don't support Client Hints (Safari, Firefox)
		for _, cookie := range r.Cookies() {
			if cookie.Name == "prefers-dark-mode" {
				prefersDark = cookie.Value == "true"
				break
			}
		}
	}

	tint, text := makeColors(prefersDark)

	data := struct {
		Tint       template.CSS
		Text       string
		Background string
	}{
		tint,
		text,
		getBackground(),
	}

	w.Header().Add("Content-Type", "text/css; charset=utf-8")

	if err := s.tpl.ExecuteTemplate(w, "styles.css", data); err != nil {
		zapmw.Extract(r.Context()).Error("execute template", zap.Error(err))
	}
}

func getBackground() string {
	backgrounds := []string{
		// "checkerboard.svg",
		// "circles.svg",
		"pixellated.svg",
		"bubbles.svg",
	}

	return backgrounds[rand.Intn(len(backgrounds))]
}

func makeColors(prefersDark bool) (tint template.CSS, text string) {
	if prefersDark {
		// Dark mode: white/light gray text and dark background
		text = "#F5F5F5" // Very light gray, almost white

		// Create a dark, desaturated background
		baseHue := float64(rand.Intn(360))
		tintColor := colorful.Hsl(baseHue, 0.3, 0.15) // Low saturation, very dark
		tint = cssRGBA(tintColor)
	} else {
		// Light mode: dark gray/black text and bright pale background
		text = "#212121" // Very dark gray, almost black

		// Create a bright, slightly desaturated background
		baseHue := float64(rand.Intn(360))
		tintColor := colorful.Hsl(baseHue, 0.3, 0.95) // Low saturation, very bright
		tint = cssRGBA(tintColor)
	}

	return
}

func cssRGBA(c color.Color) template.CSS {
	r, g, b, _ := c.RGBA()
	r, g, b = r/0x101, g/0x101, b/0x101 // convert 16bit rgb values to 8-bit
	return template.CSS(fmt.Sprintf("rgba(%d, %d, %d, 0.75)", r, g, b))
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

func (s *Server) theme(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/javascript")
	if err := s.tpl.ExecuteTemplate(w, "theme.js", nil); err != nil {
		zapmw.Extract(r.Context()).Error("execute template", zap.Error(err))
	}
}
