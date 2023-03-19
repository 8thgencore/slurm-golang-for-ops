package app

import (
	"context"
	"net/http"
	"service/api"
	"service/api/middleware"
	"service/internal/app/db"
	"service/internal/app/handlers"
	"service/internal/app/processors"
	"service/internal/app/schedulers"
	cfg "service/internal/config"
	"time"

	log "service/pkg/logger"

	"github.com/go-co-op/gocron"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Server struct {
	config    cfg.Application
	ctx       context.Context
	srv       *http.Server
	db        *pgxpool.Pool
	scheduler *gocron.Scheduler
}

func NewServer(ctx context.Context, config cfg.Application) *Server {
	server := new(Server)
	server.ctx = ctx
	server.config = config
	// create a new scheduler with UTC time zone
	server.scheduler = gocron.NewScheduler(time.UTC)

	return server
}

func (server *Server) Serve() {
	log.Infof("[!] Starting Server")
	var err error

	// start the scheduler asynchronously
	server.scheduler.StartAsync()

	// init database connection
	server.db, err = pgxpool.Connect(server.ctx, server.config.DbURL)
	if err != nil {
		log.Errorf(err.Error())
	}

	// init storage
	storage := db.NewStorage(server.db)

	// init processors
	metricsProcessor := processors.NewMetricsProcessor(storage)

	// schedule some jobs using fluent syntax
	gometrScheduler := schedulers.NewGometrScheduler(metricsProcessor)
	server.scheduler.Every(30).Seconds().Do(func() { gometrScheduler.ParseGometr() })

	// init handlers
	metricsHandler := handlers.NewMetricsHandler(metricsProcessor)

	// init router
	routes := api.CreateRoutes(metricsHandler)
	routes.Use(middleware.RequestLog)

	server.srv = &http.Server{
		Addr:    ":" + server.config.Port,
		Handler: routes,
	}

	log.Infof("[!] Server Started")

	// run server
	err = server.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Errorf(err.Error())
	}

	return
}

// Shutdown stops the app an
func (server *Server) Shutdown() {
	log.Infof("[!] Server Stopped")

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
		log.Fatalf("[!] Server Shutdown Failed:%s", err)
	}

	log.Infof("[!] Server exited properly")
}
