package main

import (
	"context"

	"github.com/Wookkie/ToDoRestful/internal"
	"github.com/Wookkie/ToDoRestful/pkg/logger"
	"github.com/rs/zerolog"

	dbstorage "github.com/Wookkie/ToDoRestful/internal/infrastracture/db_storage"
	inmemory "github.com/Wookkie/ToDoRestful/internal/infrastracture/in-memory"
	"github.com/Wookkie/ToDoRestful/internal/server"
)

func main() {
	cfg := internal.ReadConfig()
	log := logger.Get(cfg.Debug)
	log.Info().Msg("service starting")
	dsn := cfg.PostgresDSN

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
	if err := api.Run(); err != nil {
		log.Error().Err(err).Msg("Failed running server")
	}
}

func startWithInMemory(cfg *internal.Config, log zerolog.Logger) {
	repo := inmemory.New()
	log.Info().Msg("Using in-memory storage")
	api := server.New(cfg, repo)
	if err := api.Run(); err != nil {
		log.Error().Err(err).Msg("Failed running server")
	}
}
