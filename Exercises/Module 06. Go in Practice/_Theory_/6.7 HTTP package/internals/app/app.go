package app

import (
	"6_7/example/api"
	"6_7/example/api/middleware"
	db3 "6_7/example/internals/app/db"
	"6_7/example/internals/app/handlers"
	"6_7/example/internals/app/processors"
	"6_7/example/internals/cfg"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"time"
)

type AppServer struct {
	config cfg.Cfg
	ctx    context.Context
	srv    *http.Server
	db     *pgxpool.Pool
}

func NewServer(config cfg.Cfg, ctx context.Context) *AppServer { //задаем поля нашего сервера, для его старта нам нужен контекст и конфигурация
	server := new(AppServer)
	server.ctx = ctx
	server.config = config
	return server
}

func (server *AppServer) Serve() {
	log.Println("Starting server")
	log.Println(server.config.GetDBString())
	var err error
	server.db, err = pgxpool.Connect(server.ctx, server.config.GetDBString()) //создаем пул соединений с БД и сохраним его для закрытия при остановке приложения
	if err != nil {
		log.Fatalln(err)
	}

	carsStorage := db3.NewCarStorage(server.db)    //создаем экземпляр storage для работы с бд и всем что связано с машинами
	usersStorage := db3.NewUsersStorage(server.db) //создаем экземпляр storage для работы с бд и всем что связано с пользователями

	carsProcessor := processors.NewCarsProcessor(carsStorage) //инициализируем процессоры соотвествующими storage
	usersProcessor := processors.NewUsersProcessor(usersStorage)

	userHandler := handlers.NewUsersHandler(usersProcessor) //инициализируем handlerы нашими процессорами
	carsHandler := handlers.NewCarsHandler(carsProcessor)

	routes := api.CreateRoutes(userHandler, carsHandler) //хендлеры напрямую используются в путях
	routes.Use(middleware.RequestLog)                    //middleware используем здесь, хотя можно было бы и в CreateRoutes

	server.srv = &http.Server{ //в отличие от примеров http, здесь мы передаем наш server в поле структуры, для работы в Shutdown
		Addr:    ":" + server.config.Port,
		Handler: routes,
	}

	log.Println("Server started")

	err = server.srv.ListenAndServe() //запускаем сервер

	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}

	return
}

func (server *AppServer) Shutdown() {
	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	server.db.Close() //закрываем соединение с БД
	defer func() {
		cancel()
	}()
	var err error
	if err = server.srv.Shutdown(ctxShutDown); err != nil { //выключаем сервер, с ограниченным по времени контекстом
		log.Fatalf("server Shutdown Failed:%s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}
