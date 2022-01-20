package usecase

import (
	"context"
	"time"

	"github.com/dchebakov/tracker/internal/models"
	"github.com/dchebakov/tracker/internal/stats"
	"github.com/dchebakov/tracker/pkg/httperrors"
	"go.uber.org/zap"
)

type statsUC struct {
	logger    *zap.SugaredLogger
	statsRepo stats.Repository
}

func NewStatsUseCase(logger *zap.SugaredLogger, statsRepo stats.Repository) stats.UseCase {
	return &statsUC{logger: logger, statsRepo: statsRepo}
}

func (u *statsUC) UpdateStats(
	ctx context.Context,
	customerID int64,
	hour time.Time,
	valid bool,
) error {
	err := u.statsRepo.UpdateStats(ctx, customerID, hour, valid)
	if err != nil {
		u.logger.Errorw("Failed to update stats", "cause", err)
		return httperrors.NewInternalServerError(err)
	}

	return nil
}

func (u *statsUC) GetStats(
	ctx context.Context,
	filter *stats.Filter,
) ([]*models.HourlyStats, error) {
	stats, err := u.statsRepo.GetStats(ctx, filter)
	if err != nil {
		u.logger.Errorw("Failed to fetch stats", "err", err, "filter", filter)
		return nil, httperrors.NewInternalServerError(err)
	}

	return stats, nil
}
