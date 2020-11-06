package main

import (
	"compress/gzip"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/blendle/zapdriver"
	"github.com/d47id/zapmw"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// create logger
	l, err := zapdriver.NewProduction()
	if ok, err := strconv.ParseBool(os.Getenv("DEBUG")); err == nil && ok {
		l, err = getDevelopmentLogger()
	}
	if err != nil {
		panic(err)
	}

	// read port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	l = l.With(zap.String("server_port", port))

	// seed random number generator
	rand.Seed(time.Now().Unix())

	// load template
	tpl, err := template.ParseGlob("www/*")
	if err != nil {
		panic(err)
	}

	// create server
	s := &server{tpl: tpl, l: l}

	// create http router
	mux := chi.NewRouter()
	mux.Use(zapmw.New(l))
	mux.Use(gziphandler.MustNewGzipLevelHandler(gzip.BestCompression))
	mux.Get("/", s.index)
	mux.Get("/checkerboard.svg", s.checkerboard)
	mux.Get("/circles.svg", s.circles)
	mux.Get("/dave-bot.svg", s.daveBotSVG)
	mux.Get("/dave-bot.png", s.daveBotPNG)

	// start http server
	l.Info("server starting")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		l.Error("listen and serve", zap.Error(err))
	}
}

func getDevelopmentLogger() (*zap.Logger, error) {
	enc := zap.NewProductionEncoderConfig()
	enc.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	enc.EncodeDuration = zapcore.StringDurationEncoder
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig = enc
	cfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	return cfg.Build()
}
