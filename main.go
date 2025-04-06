package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/blendle/zapdriver"
	"github.com/d47id/art-bot/art"
	"github.com/d47id/art-bot/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           s,
		ReadHeaderTimeout: 5 * time.Second,
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			l.Error("listen and serve error", zap.Error(err))
			return err
		}

		return nil
	})

	g.Go(func() error {
		<-sig

		start := time.Now()
		l.Debug("shutdown signal received")

		if err := server.Shutdown(ctx); err != nil {
			return err
		}

		l.Info("server shut down complete", zap.Duration("duration", time.Since(start)))
		return nil
	})

	// start http server
	l.Info("server started")
	if err := g.Wait(); err != nil {
		panic(err)
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
