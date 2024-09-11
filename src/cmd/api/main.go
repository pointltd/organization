package main

import (
	"github.com/pointltd/organization/internal/app"
	"log"
)

func main() {
	a, err := app.NewApp()

	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	a.RunServer()
}
