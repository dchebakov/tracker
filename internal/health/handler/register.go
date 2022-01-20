package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *healthHandler) RegisterHTTPEndPoints(version *echo.Group) {
	group := version.Group("/health")

	group.GET("/", h.Health())
	group.GET("/readiness", h.Readiness())
}
