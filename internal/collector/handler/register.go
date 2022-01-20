package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *collectHandler) RegisterHTTPEndPoints(group *echo.Group) {
	group.POST("/", h.Collect())
}
