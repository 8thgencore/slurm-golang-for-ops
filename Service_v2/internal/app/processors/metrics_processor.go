package processors

import (
	"errors"
	"service/internal/app/db"
	"service/internal/app/models"
	"time"
)

type MetricsProcessor struct {
	storage *db.Storage
}

func NewMetricsProcessor(storage *db.Storage) *MetricsProcessor {
	processor := new(MetricsProcessor)
	processor.storage = storage
	return processor
}

func (processor *MetricsProcessor) Add(metric models.Metric) error {
	if metric.Name == "" {
		return errors.New("Name should not be empty")
	}

	if metric.Value == "" {
		return errors.New("Value should not be empty")
	}

	if metric.Date.IsZero() {
		return errors.New("Date shall be filled")
	}

	return processor.storage.Add(metric)
}

func (processor *MetricsProcessor) List(name string, startDate time.Time, endDate time.Time, offset int, limit int) []models.Metric {
	return processor.storage.List(name, startDate, endDate, offset, limit)
}
