package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Wookkie/ToDoRestful/internal"
	"github.com/Wookkie/ToDoRestful/pkg/logger"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"

	dbstorage "github.com/Wookkie/ToDoRestful/internal/infrastracture/db_storage"
	inmemory "github.com/Wookkie/ToDoRestful/internal/infrastracture/in-memory"
	"github.com/Wookkie/ToDoRestful/internal/server"
)

func gracefulShutdown(cancel context.CancelFunc) {
	log := logger.Get()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	s := <-c

	log.Info().Msgf("graceful shutdown with signal: %s", s)
	cancel()
}

func main() {
	cfg := internal.ReadConfig()
	log := logger.Get(cfg.Debug)
	log.Info().Msg("service starting")
	dsn := cfg.PostgresDSN

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go gracefulShutdown(cancel)

	if err := dbstorage.ApplyMigrations(dsn); err != nil {
		log.Warn().Err(err).Msg("Failed to apply migrations. Using in-memory storage")
		startWithInMemory(cfg, log)
		return
	}

	repo, err := dbstorage.New(context.Background(), dsn)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to connect to db. Using in-memory storage")
		startWithInMemory(cfg, log)
		return
	}

	log.Info().Msg("Using PostgreSQL storage")

	api := server.New(cfg, repo)

	group, gCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		err := api.Run()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-gCtx.Done()

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCancel()

		return api.Stop(shutdownCtx)
	})

	if err := group.Wait(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Service stopped with error")
		}
	}

}

func startWithInMemory(cfg *internal.Config, log zerolog.Logger) {
	repo := inmemory.New()
	log.Info().Msg("Using in-memory storage")
	api := server.New(cfg, repo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go gracefulShutdown(cancel)

	group, gCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return api.Run()
	})

	group.Go(func() error {
		<-gCtx.Done()
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()
		return api.Stop(shutdownCtx)
	})

	if err := group.Wait(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Service stopped with error")
		}
	}
}
