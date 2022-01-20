package models

import "time"

type HourlyStats struct {
	ID           int64     `json:"id"           db:"id"`
	CustomerID   int64     `json:"customerID"   db:"customer_id"`
	Time         time.Time `json:"time"         db:"time"`
	RequestCount int64     `json:"requestCount" db:"request_count"`
	InvalidCount int64     `json:"invalidCount" db:"invalid_count"`
}
