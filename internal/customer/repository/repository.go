package repository

import (
	"context"

	"github.com/dchebakov/tracker/internal/customer"
	"github.com/dchebakov/tracker/internal/models"
	"github.com/jmoiron/sqlx"
)

type customerRepo struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) customer.Repositry {
	return &customerRepo{db: db}
}

func (r *customerRepo) GetByID(ctx context.Context, id int64) (*models.Customer, error) {
	customer := &models.Customer{}
	err := r.db.QueryRowxContext(ctx, getCustomerQuery, id).StructScan(customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
