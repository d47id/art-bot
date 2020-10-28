package main

import (
	"compress/gzip"
	"html/template"
	"net/http"
	"os"

	"github.com/d47id/zapmw"

	"github.com/NYTimes/gziphandler"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// create logger
	zapenc := zap.NewProductionEncoderConfig()
	zapenc.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	zapenc.EncodeDuration = zapcore.StringDurationEncoder
	zapcfg := zap.NewProductionConfig()
	zapcfg.EncoderConfig = zapenc
	zapcfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	l, err := zapcfg.Build()
	if err != nil {
		panic(err)
	}

	// read port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	l = l.With(zap.String("server_port", port))

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
	mux.Get("/", s.index)
	mux.Route("/image", func(r chi.Router) {
		r.Use(gziphandler.MustNewGzipLevelHandler(gzip.BestCompression))
		r.Get("/checkerboard.svg", s.checkerboard)
	})

	// start http server
	l.Info("server starting")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		l.Error("listen and serve", zap.Error(err))
	}
}
