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

func (r *statsRepo) getHourStats(
	ctx context.Context,
	customerID int64,
	hour time.Time,
) (*models.HourlyStats, error) {
	stats := &models.HourlyStats{}
	err := r.db.QueryRowxContext(ctx, getHourStatsQuery, customerID, hour.Unix()).StructScan(stats)
	if err != nil {
		r.logger.Errorw(
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

func (r *statsRepo) createHourStats(
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

	err := r.db.QueryRowxContext(ctx, createHourStats, customerID, hour.Unix(), validDelta, invalidDelta).
		StructScan(stats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *statsRepo) updateHourStats(ctx context.Context, statsID int64, valid bool) error {
	validDelta := 1
	invalidDelta := 0
	if !valid {
		invalidDelta = 1
	}

	_, err := r.db.ExecContext(ctx, updateStatsQuery, statsID, validDelta, invalidDelta)
	if err != nil {
		return err
	}

	return nil
}

// TODO: make trunsaction there
func (r *statsRepo) UpdateStats(
	ctx context.Context,
	customerID int64,
	hour time.Time,
	valid bool,
) error {
	stats, err := r.getHourStats(ctx, customerID, hour)
	if err == nil {
		err = r.updateHourStats(ctx, stats.ID, valid)
		if err != nil {
			return err
		}

		r.logger.Debugw("Updated stats of existed record", "statsID", stats.ID)
		return nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	_, err = r.createHourStats(ctx, customerID, hour, valid)
	r.logger.Debug("Created new hourly stats record")
	if err != nil {
		return err
	}

	return nil
}

func (r *statsRepo) getStatsRows(ctx context.Context, filter *stats.Filter) (*sqlx.Rows, error) {
	if filter.CustomerID != nil && filter.Day != nil {
		return r.db.QueryxContext(ctx, getCustomerDayStats, filter.CustomerID, filter.Day.Unix())
	}

	if filter.CustomerID != nil {
		return r.db.QueryxContext(ctx, getCustomerStats, filter.CustomerID)
	}

	if filter.Day != nil {
		return r.db.QueryxContext(ctx, getDayStats, filter.Day.Unix())
	}

	return r.db.QueryxContext(ctx, getAllStats)
}

func (r *statsRepo) GetStats(
	ctx context.Context,
	filter *stats.Filter,
) ([]*models.HourlyStats, error) {
	rows, err := r.getStatsRows(ctx, filter)
	if err != nil {
		return nil, err
	}

	stats := make([]*models.HourlyStats, 0)
	for rows.Next() {
		hs := &models.HourlyStats{}
		if err = rows.StructScan(hs); err != nil {
			return nil, err
		}

		stats = append(stats, hs)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}
