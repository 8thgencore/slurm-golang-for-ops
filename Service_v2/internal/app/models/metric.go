package models

import "time"

type Metric struct {
	Id    int64     `json:"id"`
	Name  string    `json:"name"`
	Value string    `json:"value"`
	Date  time.Time `json:"timestamp"`
}
