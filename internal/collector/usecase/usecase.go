package usecase

import (
	"context"
	"net"
	"time"

	"github.com/dchebakov/tracker/internal/blacklist"
	"github.com/dchebakov/tracker/internal/collector"
	"github.com/dchebakov/tracker/internal/customer"
	"github.com/dchebakov/tracker/internal/stats"
	"github.com/dchebakov/tracker/pkg/httperrors"
	"github.com/dchebakov/tracker/pkg/utils"
	"go.uber.org/zap"
)

type collectorUC struct {
	logger      *zap.SugaredLogger
	customerUC  customer.UseCase
	blacklistUC blacklist.UseCase
	statsUC     stats.UseCase
}

func NewCollectorUseCase(
	logger *zap.SugaredLogger,
	customerUC customer.UseCase,
	blacklistUC blacklist.UseCase,
	statsUC stats.UseCase,
) collector.UseCase {
	return &collectorUC{
		logger:      logger,
		customerUC:  customerUC,
		blacklistUC: blacklistUC,
		statsUC:     statsUC,
	}
}

func (c *collectorUC) isLogValid(
	ctx context.Context,
	log *collector.CollectorLog,
	ua string,
) (bool, error) {
	customer, err := c.customerUC.GetByID(ctx, log.CustomerID)
	if err != nil {
		return false, err
	}

	if !customer.Active {
		return false, nil
	}

	if log.TagID == nil {
		return false, nil
	}

	if log.UserID == nil {
		return false, nil
	}

	if log.RemoteIP == nil {
		return false, nil
	}

	ip := net.ParseIP(*log.RemoteIP)
	if ip == nil {
		return false, nil
	}

	ipBlacklisted, err := c.blacklistUC.IsIPBlacklisted(ctx, utils.IPp2int(ip))
	if err != nil {
		return false, err
	}
	if ipBlacklisted {
		return false, nil
	}

	uaBlacklisted, err := c.blacklistUC.IsUABlacklisted(ctx, ua)
	if err != nil {
		return false, err
	}
	if uaBlacklisted {
		return false, nil
	}

	return true, nil
}

func (c *collectorUC) isLogValidPartially(
	ctx context.Context,
	log *collector.CollectorLog,
) (bool, error) {
	_, err := c.customerUC.GetByID(ctx, log.CustomerID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *collectorUC) UpdateStats(
	ctx context.Context,
	log *collector.CollectorLog,
	ua string,
) error {
	valid, err := c.isLogValid(ctx, log, ua)
	if err != nil {
		return err
	}

	hour := time.Unix(log.Timestamp, 0).Truncate(time.Hour)
	if valid {
		c.logger.Debugw("Log is valid", "log", log)
		return c.statsUC.UpdateStats(ctx, log.CustomerID, hour, true)
	}

	validPartially, err := c.isLogValidPartially(ctx, log)
	if err != nil {
		return err
	}

	if validPartially {
		c.logger.Debug("Log is valid partially", "log", log)
		err := c.statsUC.UpdateStats(ctx, log.CustomerID, hour, false)
		if err != nil {
			c.logger.Errorw("Failed to save partially valid log", "err", err)
			return err
		}

		c.logger.Debug("Saved call as invalid")
		return httperrors.NewBadRequestError(nil)
	}

	return httperrors.NewBadRequestError("invalid request")
}
