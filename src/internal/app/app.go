package app

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pointltd/organization/internal/infrastructure/http/route"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	serviceProvider *serviceProvider
	db              *pgxpool.Pool
	logger          *slog.Logger
}

type Validator struct {
	validator *validator.Validate
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func (cv *Validator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
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
	return nil
}

func (a *App) initDatabase() {
	slog.Info("Connecting to database", "url", os.Getenv("DATABASE_URL"))
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

	// Validator
	e.Validator = &Validator{validator: validator.New()}

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	var jwtMiddleware = echojwt.WithConfig(config)

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	api := e.Group("/v1")
	route.RegisterAuthRoutes(api, a.serviceProvider.AuthController())
	route.RegisterUserRoutes(api, a.serviceProvider.UserController(), jwtMiddleware)

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
