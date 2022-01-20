package repository

import "github.com/jmoiron/sqlx"

type healthRepo struct {
	db *sqlx.DB
}

func NewHealthRepository(db *sqlx.DB) *healthRepo {
	return &healthRepo{db: db}
}

func (r *healthRepo) Readiness() error {
	return r.db.Ping()
}
