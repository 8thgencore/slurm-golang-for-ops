package db

import (
	"context"
	"fmt"
	"service/internal/app/models"
	"time"

	log "service/pkg/logger"

	"github.com/georgysavva/scany/pgxscan"
)

func (storage *Storage) Add(metrics ...models.Metric) error {
	ctx := context.Background()
	tx, err := storage.databasePool.Begin(ctx)
	defer tx.Rollback(ctx)

	query := "INSERT INTO metrics (name, value, timestamp) VALUES ($1, $2, $3)"

	// prepare a statement for inserting metrics
	const stmtName = "insert metrics"
	_, err = tx.Prepare(ctx, stmtName, query)
	if err != nil {
		log.Errorf("Prepare: %v", err)
		return err
	}

	// loop over metrics and execute statement for each one
	for _, m := range metrics {
		_, err = tx.Exec(ctx, query, m.Name, m.Value, m.Date)
		if err != nil {
			log.Errorf("Insert row: %v", err)
			return err
		}
	}

	// commit transaction if no errors
	err = tx.Commit(ctx)
	if err != nil {
		log.Errorf("Commit transaction: %v", err)
		return err
	}

	return err
}

func (storage *Storage) List(name string, startDate time.Time, endDate time.Time, offset int, limit int) []models.Metric {
	// prepare a query with placeholders for parameters
	query := "SELECT timestamp, name, value FROM metrics WHERE 1=1"

	placeholderNum := 1
	args := make([]interface{}, 0)
	if name != "" {
		query += fmt.Sprintf(" AND name = $%d", placeholderNum)
		args = append(args, name)
		placeholderNum++
	}
	if !startDate.IsZero() {
		query += fmt.Sprintf(" AND timestamp >= $%d", placeholderNum)
		args = append(args, startDate)
		placeholderNum++
	}
	if !endDate.IsZero() {
		query += fmt.Sprintf(" AND timestamp <= $%d", placeholderNum)
		args = append(args, endDate)
		placeholderNum++
	}
	if offset != 0 {
		query += fmt.Sprintf(" OFFSET $%d", placeholderNum)
		args = append(args, offset)
		placeholderNum++
	}
	if limit != 0 {
		query += fmt.Sprintf(" LIMIT $%d", placeholderNum)
		args = append(args, limit)
		placeholderNum++
	}

	var dbResult []models.Metric

	ctx := context.Background()
	err := pgxscan.Select(ctx, storage.databasePool, &dbResult, query, args...)

	if err != nil {
		log.Errorf(err.Error())
	}

	return dbResult
}
