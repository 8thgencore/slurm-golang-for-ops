package app

import (
	"go.uber.org/zap"
	"gometr/internal/handlers"
	"gometr/internal/infrastructure/config"
	"gometr/pkg/graceful"
	"log"
)

type App struct {
	log       *zap.Logger
	cfg       *config.Config
	handler   *handlers.Handler
}

func Start() {
	app := new(App)
	if err := app.Bootstrap(); err != nil {
		log.Fatal(err)
	}

	go func() {
		app.Run()
	}()

	err := graceful.WaitShutdown()
	if err != nil {
		app.log.Fatal("gometr is dead")
	} else {
		app.log.Info("gometr gracefully stopped")
	}
}
