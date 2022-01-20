package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dchebakov/tracker/internal/blacklist"
	"github.com/jmoiron/sqlx"
)

type blacklistRepo struct {
	db *sqlx.DB
}

func NewBlacklistRepository(db *sqlx.DB) blacklist.Repository {
	return &blacklistRepo{db: db}
}

func (b *blacklistRepo) HasIP(ctx context.Context, ip uint32) (bool, error) {
	var foundIP int64
	err := b.db.QueryRowxContext(ctx, getIpQuery, ip).Scan(&foundIP)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (b *blacklistRepo) HasUA(ctx context.Context, ua string) (bool, error) {
	var foundUa string
	err := b.db.QueryRowxContext(ctx, getUaQuery, ua).Scan(&foundUa)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
