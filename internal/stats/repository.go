package stats

import (
	"context"
	"time"
)

type Repository interface {
	UpdateStats(ctx context.Context, customerID int64, timestamp time.Time, valid bool) error
}
