package app

import (
	"context"
	"net/http"
	"service/api"
	"service/api/middleware"
	"service/internal/app/db"
	"service/internal/app/handlers"
	"service/internal/app/processors"
	cfg "service/internal/config"
	log "service/pkg/logger"
	"time"

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
	server := &Server{
		config: config,
		ctx:    ctx,
	}
	server.scheduler = gocron.NewScheduler(time.UTC)

	return server
}

// Initialize the server's database connection
func (server *Server) initDB() {
	var err error
	server.db, err = pgxpool.Connect(server.ctx, server.config.DbURL)
	if err != nil {
		log.Errorf(err.Error())
	}
}

// Start the server
func (server *Server) start() {
	// Init storage
	storage := db.NewStorage(server.db)

	// Init metrics processor
	metricsProcessor := processors.NewMetricsProcessor(storage)

	// Init metrics handler
	metricsHandler := handlers.NewMetricsHandler(metricsProcessor)

	// Init router
	routes := api.CreateRoutes(metricsHandler)
	routes.Use(middleware.RequestLog)

	// Create HTTP server
	server.srv = &http.Server{
		Addr:    ":" + server.config.Port,
		Handler: routes,
	}

	// Run checkers
	runChecker(context.Background(), server, metricsProcessor)

	log.Infof("[!] Server Started")

	// Run server
	err := server.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Errorf(err.Error())
	}
}

// Serve starts the server and listens for incoming requests
func (server *Server) Serve() {
	log.Infof("[!] Starting Server")

	// Initialize the database connection
	server.initDB()

	// Start the scheduler asynchronously
	server.scheduler.StartAsync()

	// Start the server
	server.start()
}

// Shutdown the server
func (server *Server) Shutdown() {
	log.Infof("[!] Server Stopped")

	// Clears the scheduler
	server.scheduler.Clear()
	server.scheduler.Stop()

	// Close the database connection
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	server.db.Close()
	defer func() {
		cancel()
	}()

	// Shutdown the server
	var err error
	if err = server.srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("[!] Server Shutdown Failed:%s", err)
	}

	log.Infof("[!] Server exited properly")
}
