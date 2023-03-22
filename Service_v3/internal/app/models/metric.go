package models

import "time"

type Metric struct {
	ID    int64     `json:"-"`
	Name  string    `json:"name"`
	Value string    `json:"value"`
	Date  time.Time `json:"timestamp" db:"timestamp"`
}
