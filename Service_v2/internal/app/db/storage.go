package db

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	databasePool *pgxpool.Pool
}

func NewStorage(pool *pgxpool.Pool) *Storage {
	storage := new(Storage)
	storage.databasePool = pool
	return storage
}
