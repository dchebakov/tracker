package usecase

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/dchebakov/tracker/internal/customer"
	"github.com/dchebakov/tracker/internal/models"
	"github.com/dchebakov/tracker/pkg/httperrors"
	"go.uber.org/zap"
)

type customerUC struct {
	customerRepo customer.Repositry
	logger       *zap.SugaredLogger
}

func NewCustomerUseCase(
	logger *zap.SugaredLogger,
	customerRepo customer.Repositry,
) customer.UseCase {
	return &customerUC{customerRepo: customerRepo, logger: logger}
}

func (u *customerUC) GetByID(ctx context.Context, id int64) (*models.Customer, error) {
	customer, err := u.customerRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, httperrors.NewRestError(
				http.StatusNotFound,
				httperrors.ErrNotFound.Error(),
				err,
			)
		}

		return nil, httperrors.NewInternalServerError(err)
	}

	return customer, nil
}
