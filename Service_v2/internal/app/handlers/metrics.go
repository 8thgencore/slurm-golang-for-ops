package handlers

import (
	"net/http"
	"service/internal/app/models"
	"service/internal/app/processors"
	log "service/pkg/logger"
	"strconv"
	"time"
)

type MetricsHandler struct {
	processor *processors.MetricsProcessor
}

func NewMetricsHandler(processor *processors.MetricsProcessor) *MetricsHandler { //конструктор
	handler := new(MetricsHandler)
	handler.processor = processor
	return handler
}

func (handler *MetricsHandler) List(w http.ResponseWriter, r *http.Request) {
	log.Infof("Listing metrics...")

	var err error

	// get limit from query parameter or use default value of 10 if not specified or invalid
	limit := 10 // default limit value
	if r.URL.Query().Get("limit") != "" {
		limitStr := r.URL.Query().Get("limit")
		limitInt64, err := strconv.ParseInt(limitStr, 10, 32)
		if err != nil {
			log.Errorf(err.Error())
		}
		limit = int(limitInt64)
	}

	// get offset from query parameter or use default value of 0 if not specified or invalid
	offset := 0 // default offset value
	if r.URL.Query().Get("offset") != "" {
		offsetStr := r.URL.Query().Get("offset")
		offsetInt64, err := strconv.ParseInt(offsetStr, 10, 32)
		if err != nil {
			log.Errorf(err.Error())
		}
		offset = int(offsetInt64)
	}

	// get time_from from query parameter or use zero value if not specified or invalid
	timeFrom := time.Time{} // zero value for time
	if r.URL.Query().Get("time_from") != "" {
		timeFromStr := r.URL.Query().Get("time_from")
		timeFrom, err = time.Parse(time.RFC3339, timeFromStr)
		if err != nil {
			log.Errorf(err.Error())
		}
	}

	// get time_to from query parameter or use current time if not specified or invalid
	timeTo := time.Now() // current time
	if r.URL.Query().Get("time_to") != "" {
		timeToStr := r.URL.Query().Get("time_to")
		timeTo, err = time.Parse(time.RFC3339, timeToStr)
		if err != nil {
			log.Errorf(err.Error())
		}
	}

	// get name from query parameter or use empty string if not specified
	name := "" // empty string for name
	if r.URL.Query().Get("name") != "" {
		name = r.URL.Query().Get("name")
	}

	// create a slice of metrics to store the filtered and sorted results
	result := make([]models.Metric, 0)

	result = handler.processor.List(name, timeFrom, timeTo, offset, limit)

	// wrapper result
	var m = map[string]interface{}{
		"result": "OK",
		"data":   result,
	}

	WrapOK(w, m)
}
