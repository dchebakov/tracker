package handler

import (
	"github.com/dchebakov/tracker/internal/health"
	"github.com/dchebakov/tracker/pkg/httperrors"
	"github.com/dchebakov/tracker/pkg/response"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type healthHandler struct {
	logger  *zap.SugaredLogger
	useCase health.HealthUseCase
}

func NewHelthHandler(logger *zap.SugaredLogger, useCase health.HealthUseCase) health.Handler {
	return &healthHandler{logger: logger, useCase: useCase}
}

// Health checks
// @Summary Checks if API is running
// @Description Call this API to see if API is running in the server
// @Success 200
// @Failure 500
// @route /health [get]
func (h *healthHandler) Health() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(response.Ok(nil))
	}
}

// Readiness checks if database is alive
// @Summary Checks if both API and Database are up
// @Description Call this API to see if both API ands Database are running it the server
// @Success 200
// @Failure 500
// @route /health/readiness [get]
func (h *healthHandler) Readiness() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := h.useCase.Readiness()
		if err != nil {
			h.logger.Errorw("Failed to connect to DB", "err", err)
			return c.JSON(response.Error(httperrors.NewInternalServerError(err)))
		}

		return c.JSON(response.Ok(nil))
	}
}
