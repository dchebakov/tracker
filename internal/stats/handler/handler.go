package handler

import (
	"github.com/dchebakov/tracker/internal/stats"
	"github.com/dchebakov/tracker/pkg/httperrors"
	"github.com/dchebakov/tracker/pkg/response"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type statsHandler struct {
	logger  *zap.SugaredLogger
	statsUC stats.UseCase
}

func NewStatsHander(logger *zap.SugaredLogger, statsUC stats.UseCase) stats.Handler {
	return &statsHandler{logger: logger, statsUC: statsUC}
}

func (s *statsHandler) GetStats() echo.HandlerFunc {
	return func(c echo.Context) error {
		filter := &stats.Filter{}
		if err := filter.Bind(c); err != nil {
			return c.JSON(response.Error(httperrors.NewBadRequestError(err)))
		}

		ctx := c.Request().Context()
		stats, err := s.statsUC.GetStats(ctx, filter)
		if err != nil {
			return c.JSON(response.Error(err))
		}

		return c.JSON(response.Ok(stats))
	}
}
