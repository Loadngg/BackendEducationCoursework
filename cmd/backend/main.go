package main

import (
	"log/slog"

	"backend/internal/config"
	"backend/internal/pkg/backend"
	"backend/pkg/logger"
	"backend/pkg/logger/sl"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	log.Info("Backend-education work", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	app, err := backend.New(log, cfg)
	if err != nil {
		log.Error("Error creating app:", sl.Error(err))
	}

	log.Info("Server is starting")
	err = app.Run(cfg)
	if err != nil {
		log.Error("Error running app:", sl.Error(err))
	}
}
