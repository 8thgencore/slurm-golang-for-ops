package main

import (
	"context"
	"os"
	"os/signal"
	"service/internal/app"
	cfg "service/internal/config"
	log "service/pkg/logger"
	"syscall"
)

func main() {
	cfg.Setup("config/settings.yml")
	config := cfg.ApplicationConfig

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	server := app.NewServer(ctx, *config)

	go func() {
		oscall := <-c //если таки что то пришло
		log.Infof("system call:%+v", oscall)
		server.Shutdown() //выключаем сервер
		cancel()
	}()

	server.Serve()
}
