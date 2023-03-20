package processors

import (
	"errors"
	"service/internal/app/db"
	"service/internal/app/models"
	"time"
)

// Объявляем константы для ошибок
var (
	ErrEmptyName  = errors.New("Name should not be empty")
	ErrEmptyValue = errors.New("Value should not be empty")
	ErrEmptyDate  = errors.New("Date shall be filled")
)

type MetricsProcessor struct {
	storage db.StorageInterface
}

func NewMetricsProcessor(storage db.StorageInterface) *MetricsProcessor {
	processor := new(MetricsProcessor)
	processor.storage = storage

	return processor
}

func (processor *MetricsProcessor) Add(metric models.Metric) error {
	if metric.Name == "" {
		return ErrEmptyName
	}

	if metric.Value == "" {
		return ErrEmptyValue
	}

	if metric.Date.IsZero() {
		return ErrEmptyDate
	}

	return processor.storage.Add(metric)
}

func (processor *MetricsProcessor) List(name string, startDate time.Time, endDate time.Time, offset int, limit int) []models.Metric {
	return processor.storage.List(name, startDate, endDate, offset, limit)
}
