package app

import (
	"github.com/pointltd/organization/internal/server"
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

func (a *App) RunServer() {
	server.Run()
}
