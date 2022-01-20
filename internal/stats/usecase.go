package stats

import (
	"context"
	"time"
)

type UseCase interface {
	UpdateStats(ctx context.Context, customerID int64, hour time.Time, valid bool) error
}
