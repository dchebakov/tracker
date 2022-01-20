package usecase

import (
	"context"
	"time"

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

func (s *statsUC) UpdateStats(
	ctx context.Context,
	customerID int64,
	hour time.Time,
	valid bool,
) error {
	err := s.statsRepo.UpdateStats(ctx, customerID, hour, valid)
	if err != nil {
		s.logger.Errorw("Failed to update stats", "cause", err)
		return httperrors.NewInternalServerError(err)
	}

	return nil
}
