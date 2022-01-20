package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dchebakov/tracker/internal/models"
	"github.com/dchebakov/tracker/internal/stats"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type statsRepo struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewStatsRepository(db *sqlx.DB, logger *zap.SugaredLogger) stats.Repository {
	return &statsRepo{db: db, logger: logger}
}

func (s *statsRepo) getHourStats(
	ctx context.Context,
	customerID int64,
	hour time.Time,
) (*models.HourlyStats, error) {
	stats := &models.HourlyStats{}
	err := s.db.QueryRowxContext(ctx, getHourStatsQuery, customerID, hour.Unix()).StructScan(stats)
	if err != nil {
		s.logger.Errorw(
			"Failed to find stats for the customer and given hour",
			"customerID",
			customerID,
			"hour",
			hour,
			"error",
			err,
		)
		return nil, err
	}

	return stats, nil
}

func (s *statsRepo) createHourStats(
	ctx context.Context,
	customerID int64,
	hour time.Time,
	valid bool,
) (*models.HourlyStats, error) {
	stats := &models.HourlyStats{}
	validDelta := 1
	invalidDelta := 0
	if !valid {
		invalidDelta = 1
	}

	err := s.db.QueryRowxContext(ctx, createHourStats, customerID, hour.Unix(), validDelta, invalidDelta).
		StructScan(stats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (s *statsRepo) updateHourStats(ctx context.Context, statsID int64, valid bool) error {
	validDelta := 1
	invalidDelta := 0
	if !valid {
		invalidDelta = 1
	}

	_, err := s.db.ExecContext(ctx, updateStatsQuery, statsID, validDelta, invalidDelta)
	if err != nil {
		return err
	}

	return nil
}

func (s *statsRepo) UpdateStats(
	ctx context.Context,
	customerID int64,
	hour time.Time,
	valid bool,
) error {
	stats, err := s.getHourStats(ctx, customerID, hour)
	if err == nil {
		err = s.updateHourStats(ctx, stats.ID, valid)
		if err != nil {
			return err
		}

		s.logger.Debugw("Updated stats of existed record", "statsID", stats.ID)
		return nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	_, err = s.createHourStats(ctx, customerID, hour, valid)
	s.logger.Debug("Created new hourly stats record")
	if err != nil {
		return err
	}

	return nil
}
