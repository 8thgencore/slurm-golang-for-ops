package db

import (
	"service/internal/app/models"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Определяем интерфейс StorageInterface
type StorageInterface interface {
	Add(metric ...models.Metric) error
	List(name string, startDate time.Time, endDate time.Time, offset int, limit int) []models.Metric
}

type Storage struct {
	databasePool *pgxpool.Pool
}

func NewStorage(pool *pgxpool.Pool) *Storage {
	storage := new(Storage)
	storage.databasePool = pool

	return storage
}
