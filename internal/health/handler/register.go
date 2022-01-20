package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *healthHandler) RegisterHTTPEndPoints(group *echo.Group) {
	group.GET("/", h.Health())
	group.GET("/readiness", h.Readiness())
}
