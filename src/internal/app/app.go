package app

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	controllerProvider *controllerProvider
	serviceProvider    *serviceProvider
	db                 *pgxpool.Pool
	logger             *slog.Logger
	httpServer         *httpServer
}

func NewApp(logger *slog.Logger) (*App, error) {
	a := &App{
		logger: logger,
	}

	err := a.init()

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) init() error {
	a.initDatabase()
	a.serviceProvider = newServiceProvider(a.db, a.logger)
	a.controllerProvider = newControllerProvider(a.serviceProvider)
	return nil
}

func (a *App) initDatabase() {
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("Unable to connect to database: %v\n", err)
		return
	}
	slog.Info("Connected to database")

	a.db = dbPool
}

func (a *App) RunHttpServer() error {
	a.httpServer = NewHttpServer(a.logger, a)
	return a.httpServer.Start()
}
