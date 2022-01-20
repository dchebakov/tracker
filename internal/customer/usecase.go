package customer

import (
	"context"

	"github.com/dchebakov/tracker/internal/models"
)

type UseCase interface {
	GetByID(ctx context.Context, id int64) (*models.Customer, error)
}
