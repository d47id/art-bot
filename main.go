package main

import (
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/blendle/zapdriver"
	"github.com/d47id/art-bot/art"
	"github.com/d47id/art-bot/server"
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
	l.Info("server starting")

	// seed random number generator
	rand.Seed(time.Now().Unix())

	// create bot
	bot, err := art.New()
	if err != nil {
		panic(err)
	}

	// create server
	s, err := server.New(l, bot)
	if err != nil {
		panic(err)
	}

	// start http server
	l.Info("server started")
	if err := http.ListenAndServe(":"+port, s); err != nil {
		l.Error("listen and serve", zap.Error(err))
	}

	// TODO listen for signals and gracefully shut down
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
