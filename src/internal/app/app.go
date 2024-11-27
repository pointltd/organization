package app

import (
	"context"
	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pointltd/organization/internal/infrastructure/http"
	"github.com/pointltd/organization/internal/infrastructure/http/route"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	controllerProvider *controllerProvider
	serviceProvider    *serviceProvider
	db                 *pgxpool.Pool
	logger             *slog.Logger
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

func (a *App) RunHttpServer() {
	e := echo.New()

	e.Validator = &http.Validator{Validator: validator.New()}

	var jwtMiddleware = echojwt.WithConfig(http.GetJwtConfig())

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	api := e.Group("/v1")
	route.RegisterAuthRoutes(api, a.controllerProvider.AuthController())
	route.RegisterUserRoutes(api, a.controllerProvider.UserController(), jwtMiddleware)

	// Port configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := e.Start(":" + port)
	if err != nil {
		a.logger.Error("failed to start server: %v", err)
	}
}
