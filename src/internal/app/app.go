package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pointltd/organization/internal/infrastructure/route"
	"os"
)

type App struct {
	serviceProvider *serviceProvider
}

func NewApp() (*App, error) {
	a := &App{}

	err := a.init()

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) init() error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) RunHttpServer() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	api := e.Group("/v1")
	route.RegisterUserRoutes(api, a.serviceProvider.UserController())

	// Port configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}