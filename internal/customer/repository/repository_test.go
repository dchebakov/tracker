package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dchebakov/tracker/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestCustomerRepo_GetByID(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	customerRepo := NewCustomerRepository(sqlxDB)
	t.Run("GetByID", func(t *testing.T) {
		id := int64(1)
		rows := sqlmock.NewRows([]string{
			"id",
			"name",
			"active",
		}).AddRow(id, "Big News Media Corp", true)
		mock.ExpectQuery(getCustomerQuery).WithArgs(id).WillReturnRows(rows)

		testCustomer := &models.Customer{
			ID:     id,
			Name:   "Big News Media Corp",
			Active: true,
		}
		customer, err := customerRepo.GetByID(context.Background(), id)
		require.NoError(t, err)
		require.Equal(t, testCustomer, customer)
	})
}
