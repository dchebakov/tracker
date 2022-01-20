package handler

import (
	"github.com/dchebakov/tracker/internal/collector"
	"github.com/dchebakov/tracker/pkg/httperrors"
	"github.com/dchebakov/tracker/pkg/response"
	"github.com/dchebakov/tracker/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type collectHandler struct {
	logger      *zap.SugaredLogger
	validate    *validator.Validate
	collectorUC collector.UseCase
}

func NewCollectorHandler(
	logger *zap.SugaredLogger,
	validate *validator.Validate,
	collectorUC collector.UseCase,
) collector.Handler {
	return &collectHandler{logger: logger, validate: validate, collectorUC: collectorUC}
}

// @Summary Collector API
// @Description Collect valid/invalid API calls records
// @Accept  json
// @Produce json
// @Router /collect [post]
func (h *collectHandler) Collect() echo.HandlerFunc {
	return func(c echo.Context) error {
		log := &collector.CollectorLog{}
		if err := utils.ReadRequest(c, log, h.validate); err != nil {
			h.logger.Error(err)
			return c.JSON(response.Error(httperrors.NewBadRequestError(err)))
		}

		ctx := c.Request().Context()
		ua := c.Request().Header.Get("User-Agent")
		err := h.collectorUC.UpdateStats(ctx, log, ua)
		if err != nil {
			return c.JSON(response.Error(err))
		}

		return c.JSON(response.Ok(nil))
	}
}
