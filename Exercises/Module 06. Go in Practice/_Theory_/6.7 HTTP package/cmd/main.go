package main

import (
	"6_7/example/internals/app"
	"6_7/example/internals/cfg"
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() { //точка входа нашего сервера
	config := cfg.LoadAndStoreConfig() //грузим конфигурацию

	ctx, cancel := context.WithCancel(context.Background()) // создаем контекст для работы контекстнозависимых частей системы

	c := make(chan os.Signal, 1) //создаем канал для сигналов системы
	signal.Notify(c, os.Interrupt) //в случае поступления сигнала завершения - уведомляем наш канал

	server := app.NewServer(config, ctx) // создаем сервер

	go func() { //горутина для ловли сообщений системы
		oscall := <-c //если таки что то пришло
		log.Printf("system call:%+v", oscall)
		server.Shutdown() //выключаем сервер
		cancel() //отменяем контекст
	}()

	server.Serve() //запускаем сервер
}
