package main

import (
	"github.com/pointltd/organization/internal/app"
	"github.com/pointltd/organization/internal/config"
	"os"

	"log/slog"
)

func main() {
	appConfig := config.NewAppConfig()
	logger := initLogger(appConfig)

	a, err := app.NewApp(logger, appConfig)
	if err != nil {
		logger.Error("failed to init app: %s", err.Error())
		return
	}

	err = a.RunHttpServer()
	if err != nil {
		logger.Error("failed to run server: %s", err.Error())
		return
	}
}

func initLogger(appConfig config.AppConfig) *slog.Logger {
	if appConfig.Env() == "local" {
		return slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
