package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pointltd/organization/internal/routes"
	"os"
)

func Run() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	api := e.Group("/v1")
	routes.RegisterArtistRoutes(api)

	// Port configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
