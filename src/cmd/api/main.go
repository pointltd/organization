package main

import (
	"github.com/pointltd/organization/internal/app"
	"github.com/pointltd/organization/internal/infrastructure/http"
	"os"

	"log/slog"
)

func main() {
	logger := initLogger()

	a, err := app.NewApp(logger)
	if err != nil {
		logger.Error("failed to init app: %s", err.Error())
		return
	}

	s := http.NewServer(logger, a)

	err = s.RunHttpServer()
	if err != nil {
		logger.Error("failed to run server: %s", err.Error())
		return
	}
}

func initLogger() *slog.Logger {
	if os.Getenv("ENV") == "local" {
		return slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
