package handler

import "github.com/labstack/echo/v4"

func (h *statsHandler) RegisterHTTPEndPoints(version *echo.Group) {
	group := version.Group("/stats")
	group.GET("/", h.GetStats())
}
