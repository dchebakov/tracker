package customer

import (
	"context"

	"github.com/dchebakov/tracker/internal/models"
)

type Repositry interface {
	GetByID(ctx context.Context, id int64) (*models.Customer, error)
}
