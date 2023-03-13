package api

import (
	"github.com/gorilla/mux"
	"service/internal/app/handlers"
)

func CreateRoutes(metricsHandler *handlers.MetricsHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/metrics/list", metricsHandler.List).Methods("Get")

	//оборачиваем 404, для обработки NotFound
	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()
	return r
}
