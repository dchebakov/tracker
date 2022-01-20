package models

type Customer struct {
	ID     int64  `json:"id"     db:"id"`
	Name   string `json:"name"   db:"name"`
	Active bool   `json:"active" db:"active"`
}
