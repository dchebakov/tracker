package stats

import (
	"context"
	"time"

	"github.com/dchebakov/tracker/internal/models"
)

type UseCase interface {
	UpdateStats(ctx context.Context, customerID int64, hour time.Time, valid bool) error
	GetStats(ctx context.Context, filter *Filter) ([]*models.HourlyStats, error)
}
