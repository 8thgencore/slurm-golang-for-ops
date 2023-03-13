package app

import (
	"context"
	"net/http"
	"service/api"
	"service/api/middleware"
	db3 "service/internal/app/db"
	"service/internal/app/handlers"
	"service/internal/app/processors"
	"service/internal/app/schedulers"
	"service/internal/cfg"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type AppServer struct {
	config    cfg.Cfg
	ctx       context.Context
	srv       *http.Server
	db        *pgxpool.Pool
	scheduler *gocron.Scheduler
}

func NewServer(config cfg.Cfg, ctx context.Context) *AppServer {
	server := new(AppServer)
	server.ctx = ctx
	server.config = config
	// create a new scheduler with UTC time zone
	server.scheduler = gocron.NewScheduler(time.UTC)
	return server
}

func (server *AppServer) Serve() {
	log.Println("Starting server")
	var err error

	// start the scheduler asynchronously
	server.scheduler.StartAsync()

	// init database connection
	log.Println(server.config.GetDBString())
	server.db, err = pgxpool.Connect(server.ctx, server.config.GetDBString())
	if err != nil {
		log.Fatalln(err)
	}

	// init storage
	storage := db3.NewStorage(server.db)

	// init processors
	metricsProcessor := processors.NewMetricsProcessor(storage)

	// schedule some jobs using fluent syntax
	gometrScheduler := schedulers.NewGometrScheduler(metricsProcessor)
	server.scheduler.Every(30).Seconds().Do(gometrScheduler.ParseGometr())

	// init handlers
	metricsHandler := handlers.NewMetricsHandler(metricsProcessor)

	// init router
	routes := api.CreateRoutes(metricsHandler)
	routes.Use(middleware.RequestLog)

	server.srv = &http.Server{
		Addr:    ":" + server.config.Port,
		Handler: routes,
	}

	log.Println("Server started")

	// run server
	err = server.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}

	return
}

// Shutdown stops the app an
func (server *AppServer) Shutdown() {
	log.Printf("server stopped")

	// Clears the scheduler
	server.scheduler.Clear()
	server.scheduler.Stop()

	// Close database connection
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	server.db.Close()
	defer func() {
		cancel()
	}()

	// Shutdown server
	var err error
	if err = server.srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
}
