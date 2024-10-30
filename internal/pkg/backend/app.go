package backend

import (
	"log/slog"
	"os"

	"backend/internal/backend/database"
	"backend/internal/backend/database/postgres"
	"backend/internal/backend/handler"
	"backend/internal/backend/server"
	"backend/internal/backend/service"
	"backend/internal/config"
	"backend/pkg/logger/sl"
)

type App struct {
	database *database.Database
	handler  *handler.Handler
	service  *service.Service
	server   *server.Server
}

func New(log *slog.Logger, cfg *config.Config) (*App, error) {
	app := &App{}

	db, err := postgres.New(postgres.Config{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Username: cfg.User,
		Password: cfg.Password,
		Database: cfg.Database,
		SSL:      cfg.SSL,
	})

	if err != nil {
		log.Error("Failed to init storage", sl.Error(err))
		os.Exit(1)
	}

	app.database = database.New(db)
	app.service = service.New(app.database, cfg.Salt)
	app.handler = handler.New(log, app.service)
	app.server = server.New(cfg.Env, log, app.handler)

	return app, nil
}

func (app *App) Run(cfg *config.Config) error {
	return app.server.Run(cfg)
}
