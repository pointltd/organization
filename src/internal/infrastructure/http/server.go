package http

import (
	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pointltd/organization/internal/app"
	"github.com/pointltd/organization/internal/infrastructure/http/route"
	"log/slog"
	"os"
)

type Server struct {
	log *slog.Logger
	app *app.App
}

func NewServer(log *slog.Logger, app *app.App) *Server {
	return &Server{
		log: log,
		app: app,
	}
}

func (s *Server) RunHttpServer() error {
	e := echo.New()

	e.Validator = &Validator{Validator: validator.New()}

	var jwtMiddleware = echojwt.WithConfig(GetJwtConfig())

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	api := e.Group("/v1")
	route.RegisterAuthRoutes(api, s.app.ControllerProvider.AuthController())
	route.RegisterUserRoutes(api, s.app.ControllerProvider.UserController(), jwtMiddleware)

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
