package usecase

import (
	"context"

	"github.com/dchebakov/tracker/internal/blacklist"
	"github.com/dchebakov/tracker/pkg/httperrors"
	"go.uber.org/zap"
)

type blacklistUC struct {
	logger        *zap.SugaredLogger
	blacklistRepo blacklist.Repository
}

func NewBlacklistUseCase(
	logger *zap.SugaredLogger,
	blacklistRepo blacklist.Repository,
) blacklist.UseCase {
	return &blacklistUC{logger: logger, blacklistRepo: blacklistRepo}
}

func (u *blacklistUC) IsIPBlacklisted(ctx context.Context, ip uint32) (bool, error) {
	blacklisted, err := u.blacklistRepo.HasIP(ctx, ip)
	if err != nil {
		return false, httperrors.NewInternalServerError(err)
	}

	return blacklisted, nil
}

func (u *blacklistUC) IsUABlacklisted(ctx context.Context, ua string) (bool, error) {
	blacklisted, err := u.blacklistRepo.HasUA(ctx, ua)
	if err != nil {
		return false, httperrors.NewInternalServerError(err)
	}

	return blacklisted, nil
}
