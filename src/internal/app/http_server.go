package app

import (
	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pointltd/organization/internal/infrastructure/http"
	"github.com/pointltd/organization/internal/infrastructure/http/route"
	"log/slog"
	"os"
)

type httpServer struct {
	log *slog.Logger
	app *App
}

func NewHttpServer(log *slog.Logger, app *App) *httpServer {
	return &httpServer{
		log: log,
		app: app,
	}
}

func (s *httpServer) Start() error {
	e := echo.New()

	e.Validator = &http.Validator{Validator: validator.New()}

	var jwtMiddleware = echojwt.WithConfig(http.GetJwtConfig())

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	api := e.Group("/v1")
	route.RegisterAuthRoutes(api, s.app.controllerProvider.AuthController())
	route.RegisterUserRoutes(api, s.app.controllerProvider.UserController(), jwtMiddleware)

	// Port configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := e.Start(":" + port)
	if err != nil {
		s.log.Error("failed to start server: %v", err)
		return err
	}

	return nil
}
