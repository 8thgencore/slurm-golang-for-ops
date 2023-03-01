package handlers

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


func (h *Handler) newRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", h.GetHealth).Methods("GET")
	router.Handle("/metrics", promhttp.Handler())

	return router
}
