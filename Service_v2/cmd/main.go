package main

import (
	"context"
	"os"
	"os/signal"
	"service/internal/app"
	"service/internal/cfg"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {
	config := cfg.LoadAndStoreConfig()

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	server := app.NewServer(config, ctx)

	go func() {
		oscall := <-c //если таки что то пришло
		log.Printf("system call:%+v", oscall)
		server.Shutdown() //выключаем сервер
		cancel()
	}()

	server.Serve()
}
