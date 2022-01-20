package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *collectHandler) RegisterHTTPEndPoints(version *echo.Group) {
	group := version.Group("/collect")
	group.POST("/", h.Collect())
}
